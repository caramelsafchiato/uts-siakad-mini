package handler

import (
	"strconv"
	"uts-siakad-mini/app/model"
	"uts-siakad-mini/app/service"
	"uts-siakad-mini/helper"

	"github.com/gofiber/fiber/v2"
)

type EnrollmentHandler struct {
	Service *service.EnrollmentService
}

func NewEnrollmentHandler(service *service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{Service: service}
}

func (h *EnrollmentHandler) EnrollCourse(c *fiber.Ctx) error {
	var req model.Enrollment
	if err := c.BodyParser(&req); err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "Format JSON tidak valid")
	}

	if req.StudentID == 0 || req.CourseID == 0 || req.TahunAkademik == "" {
		return helper.Fail(c, fiber.StatusBadRequest, "student_id, course_id, dan tahun_akademik wajib diisi")
	}

	if err := h.Service.TakeCourse(c.Context(), req); err != nil {
		if err.Error() == "kuota mata kuliah sudah penuh" {
			return helper.Fail(c, fiber.StatusConflict, err.Error()) // 409 Conflict
		}
		return helper.Fail(c, fiber.StatusInternalServerError, "Gagal mengambil mata kuliah atau data tidak ditemukan")
	}

	return helper.OK(c, fiber.StatusCreated, "Berhasil mengambil mata kuliah (KRS)", nil, nil)
}

func (h *EnrollmentHandler) GetMyEnrollments(c *fiber.Ctx) error {
	studentID, err := strconv.Atoi(c.Params("student_id"))
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID Mahasiswa tidak valid")
	}

	enrollments, err := h.Service.GetStudentEnrollments(c.Context(), studentID)
	if err != nil {
		return helper.Fail(c, fiber.StatusInternalServerError, "Gagal mengambil data KRS")
	}

	return helper.OK(c, fiber.StatusOK, "Berhasil mendapatkan daftar KRS", enrollments, nil)
}