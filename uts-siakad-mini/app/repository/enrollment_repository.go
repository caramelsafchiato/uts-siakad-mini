package repository

import (
	"context"
	"errors"
	"time"
	"uts-siakad-mini/app/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EnrollmentRepository struct {
	DB *pgxpool.Pool
}

func NewEnrollmentRepository(db *pgxpool.Pool) *EnrollmentRepository {
	return &EnrollmentRepository{DB: db}
}

// Enroll memproses pengambilan KRS dengan Transaction
func (r *EnrollmentRepository) Enroll(ctx context.Context, enrollment model.Enrollment) error {
	// Mulai transaksi database
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// 1. Cek kuota dengan penguncian baris (FOR UPDATE)
	var kuota int
	err = tx.QueryRow(ctx, "SELECT kuota FROM courses WHERE id = $1 FOR UPDATE", enrollment.CourseID).Scan(&kuota)
	if err != nil {
		return err // Mata kuliah mungkin tidak ditemukan
	}

	if kuota <= 0 {
		return errors.New("kuota mata kuliah sudah penuh")
	}

	// 2. Kurangi kuota mata kuliah (-1)
	_, err = tx.Exec(ctx, "UPDATE courses SET kuota = kuota - 1 WHERE id = $1", enrollment.CourseID)
	if err != nil {
		return err
	}

	// 3. Masukkan data KRS
	query := `INSERT INTO enrollments (student_id, course_id, tahun_akademik, created_at) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(ctx, query, enrollment.StudentID, enrollment.CourseID, enrollment.TahunAkademik, time.Now())
	if err != nil {
		return err
	}

	// Selesaikan transaksi
	return tx.Commit(ctx)
}

// GetByStudentID mengambil daftar mata kuliah yang sudah diambil oleh seorang mahasiswa
func (r *EnrollmentRepository) GetByStudentID(ctx context.Context, studentID int) ([]map[string]interface{}, error) {
	query := `
		SELECT e.id, c.kode_mk, c.nama_mk, c.sks, e.tahun_akademik, e.created_at
		FROM enrollments e
		JOIN courses c ON e.course_id = c.id
		WHERE e.student_id = $1
	`
	rows, err := r.DB.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var id, sks int
		var kodeMk, namaMk, tahunAkademik string
		var createdAt time.Time

		if err := rows.Scan(&id, &kodeMk, &namaMk, &sks, &tahunAkademik, &createdAt); err != nil {
			return nil, err
		}

		results = append(results, map[string]interface{}{
			"id":             id,
			"kode_mk":        kodeMk,
			"nama_mk":        namaMk,
			"sks":            sks,
			"tahun_akademik": tahunAkademik,
			"created_at":     createdAt,
		})
	}
	return results, nil
}

// GetAll mengambil seluruh data pendaftaran KRS (biasanya untuk admin)
func (r *EnrollmentRepository) GetAll(ctx context.Context) ([]model.Enrollment, error) {
	query := `SELECT id, student_id, course_id, tahun_akademik, created_at FROM enrollments`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []model.Enrollment
	for rows.Next() {
		var e model.Enrollment
		err := rows.Scan(&e.ID, &e.StudentID, &e.CourseID, &e.TahunAkademik, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, e)
	}
	return enrollments, nil
}

// Delete membatalkan KRS dan mengembalikan kuota (+1)
func (r *EnrollmentRepository) Delete(ctx context.Context, id int) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// 1. Cari course_id dari data pendaftaran ini
	var courseID int
	err = tx.QueryRow(ctx, "SELECT course_id FROM enrollments WHERE id = $1 FOR UPDATE", id).Scan(&courseID)
	if err != nil {
		return err // Data KRS tidak ditemukan
	}

	// 2. Hapus data KRS
	_, err = tx.Exec(ctx, "DELETE FROM enrollments WHERE id = $1", id)
	if err != nil {
		return err
	}

	// 3. Kembalikan kuota (+1) ke tabel courses
	_, err = tx.Exec(ctx, "UPDATE courses SET kuota = kuota + 1 WHERE id = $1", courseID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}