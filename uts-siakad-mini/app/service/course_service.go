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

func (s *CourseService) GetCourseByID(ctx context.Context, id int) (*model.Course, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *CourseService) CreateCourse(ctx context.Context, course model.Course) error {
	return s.Repo.Create(ctx, course)
}

func (s *CourseService) UpdateCourse(ctx context.Context, id int, course model.Course) error {
	if _, err := s.Repo.GetByID(ctx, id); err != nil {
		return err
	}
	return s.Repo.Update(ctx, id, course)
}

func (s *CourseService) DeleteCourse(ctx context.Context, id int) error {
	if _, err := s.Repo.GetByID(ctx, id); err != nil {
		return err
	}
	return s.Repo.Delete(ctx, id)
}