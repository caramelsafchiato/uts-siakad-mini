package repository

import (
	"context"
	"uts-siakad-mini/app/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseRepository struct {
	DB *pgxpool.Pool
}

func NewCourseRepository(db *pgxpool.Pool) *CourseRepository {
	return &CourseRepository{DB: db}
}

func (r *CourseRepository) GetAll(ctx context.Context) ([]model.Course, error) {
	query := `SELECT id, kode_mk, nama_mk, sks, semester, kuota FROM courses`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.Course
	for rows.Next() {
		var c model.Course
		err := rows.Scan(&c.ID, &c.KodeMK, &c.NamaMK, &c.SKS, &c.Semester, &c.Kuota)
		if err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

func (r *CourseRepository) GetByID(ctx context.Context, id int) (*model.Course, error) {
	query := `SELECT id, kode_mk, nama_mk, sks, semester, kuota FROM courses WHERE id = $1`
	var c model.Course
	err := r.DB.QueryRow(ctx, query, id).Scan(&c.ID, &c.KodeMK, &c.NamaMK, &c.SKS, &c.Semester, &c.Kuota)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CourseRepository) Create(ctx context.Context, c model.Course) error {
	query := `INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(ctx, query, c.KodeMK, c.NamaMK, c.SKS, c.Semester, c.Kuota)
	return err
}

func (r *CourseRepository) Update(ctx context.Context, id int, c model.Course) error {
	query := `UPDATE courses SET kode_mk = $1, nama_mk = $2, sks = $3, semester = $4, kuota = $5 WHERE id = $6`
	_, err := r.DB.Exec(ctx, query, c.KodeMK, c.NamaMK, c.SKS, c.Semester, c.Kuota, id)
	return err
}

func (r *CourseRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM courses WHERE id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}