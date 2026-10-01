package service

import (
	"context"
	"uts-siakad-mini/app/model"
	"uts-siakad-mini/app/repository"
)

type StudentService struct {
	Repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{Repo: repo}
}

func (s *StudentService) GetAllStudents(ctx context.Context) ([]model.Student, error) {
	return s.Repo.GetAll(ctx)
}

func (s *StudentService) GetStudentByID(ctx context.Context, id int) (*model.Student, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *StudentService) UpdateStudent(ctx context.Context, id int, student model.Student) error {
	// Pastikan dulu data mahasiswanya ada
	_, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.Repo.Update(ctx, id, student)
}

func (s *StudentService) DeleteStudent(ctx context.Context, id int) error {
	// Cek dulu apakah data mahasiswanya ada
	_, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.Repo.Delete(ctx, id)
}