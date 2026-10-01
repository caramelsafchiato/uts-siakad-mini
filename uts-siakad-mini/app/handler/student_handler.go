package handler

import (
	"uts-siakad-mini/app/service"
	"uts-siakad-mini/helper"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	Service *service.StudentService
}

func NewStudentHandler(service *service.StudentService) *StudentHandler {
	return &StudentHandler{Service: service}
}

func (h *StudentHandler) GetAll(c *fiber.Ctx) error {
	students, err := h.Service.GetAllStudents(c.Context())
	if err != nil {
		return helper.Fail(c, fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa")
	}

	return helper.OK(c, fiber.StatusOK, "Berhasil mengambil daftar mahasiswa", students, nil)
}