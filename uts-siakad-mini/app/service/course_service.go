package service

import (
	"context"
	"uts-siakad-mini/app/model"
	"uts-siakad-mini/app/repository"
)

type CourseService struct {
	Repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *CourseService {
	return &CourseService{Repo: repo}
}

func (s *CourseService) GetAllCourses(ctx context.Context) ([]model.Course, error) {
	return s.Repo.GetAll(ctx)
}