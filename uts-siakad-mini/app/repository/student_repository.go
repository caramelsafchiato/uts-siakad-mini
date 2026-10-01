package repository

import (
	"context"
	"uts-siakad-mini/app/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepository struct {
	DB *pgxpool.Pool
}

func NewStudentRepository(db *pgxpool.Pool) *StudentRepository {
	return &StudentRepository{DB: db}
}

// GetAll mengambil seluruh data mahasiswa dari database
func (r *StudentRepository) GetAll(ctx context.Context) ([]model.Student, error) {
	query := `SELECT id, user_id, nim, nama, prodi, angkatan, ipk_terakhir FROM students`
	
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var s model.Student
		err := rows.Scan(&s.ID, &s.UserID, &s.Nim, &s.Nama, &s.Prodi, &s.Angkatan, &s.IpkTerakhir)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, nil
}

// GetByID mengambil satu data mahasiswa berdasarkan ID
func (r *StudentRepository) GetByID(ctx context.Context, id int) (*model.Student, error) {
	query := `SELECT id, user_id, nim, nama, prodi, angkatan, ipk_terakhir FROM students WHERE id = $1`
	
	var s model.Student
	err := r.DB.QueryRow(ctx, query, id).Scan(&s.ID, &s.UserID, &s.Nim, &s.Nama, &s.Prodi, &s.Angkatan, &s.IpkTerakhir)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *StudentRepository) Update(ctx context.Context, id int, s model.Student) error {
	query := `UPDATE students SET nim = $1, nama = $2, prodi = $3, angkatan = $4, ipk_terakhir = $5 WHERE id = $6`
	
	_, err := r.DB.Exec(ctx, query, s.Nim, s.Nama, s.Prodi, s.Angkatan, s.IpkTerakhir, id)
	return err
}

// Delete menghapus data mahasiswa berdasarkan ID
func (r *StudentRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM students WHERE id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}