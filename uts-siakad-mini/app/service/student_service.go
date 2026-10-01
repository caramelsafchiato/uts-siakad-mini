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