package service

import (
	"context"
	"uts-siakad-mini/app/model"
	"uts-siakad-mini/app/repository"
)

type EnrollmentService struct {
	Repo *repository.EnrollmentRepository
}

func NewEnrollmentService(repo *repository.EnrollmentRepository) *EnrollmentService {
	return &EnrollmentService{Repo: repo}
}

func (s *EnrollmentService) TakeCourse(ctx context.Context, enrollment model.Enrollment) error {
	return s.Repo.Enroll(ctx, enrollment)
}

func (s *EnrollmentService) GetStudentEnrollments(ctx context.Context, studentID int) ([]map[string]interface{}, error) {
	return s.Repo.GetByStudentID(ctx, studentID)
}