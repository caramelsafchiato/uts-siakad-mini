package handler

import (
	"uts-siakad-mini/app/service"
	"uts-siakad-mini/helper"
	"strconv"

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

func (h *StudentHandler) GetByID(c *fiber.Ctx) error {
	// Ambil parameter :id dari URL
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID mahasiswa tidak valid")
	}

	student, err := h.Service.GetStudentByID(c.Context(), id)
	if err != nil {
		return helper.Fail(c, fiber.StatusNotFound, "Data mahasiswa tidak ditemukan")
	}

	return helper.OK(c, fiber.StatusOK, "Berhasil mengambil detail mahasiswa", student, nil)
}