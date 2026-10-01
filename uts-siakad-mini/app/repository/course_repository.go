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

// GetAll mengambil seluruh daftar mata kuliah
func (r *CourseRepository) GetAll(ctx context.Context) ([]model.Course, error) {
	query := `SELECT id, kode_mk, nama_mk, sks, semester FROM courses`
	
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.Course
	for rows.Next() {
		var c model.Course
		err := rows.Scan(&c.ID, &c.KodeMK, &c.NamaMK, &c.SKS, &c.Semester)
		if err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}

	return courses, nil
}