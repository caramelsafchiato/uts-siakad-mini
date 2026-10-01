package handler

import (
	"strconv"
	"uts-siakad-mini/app/model"
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

func (h *StudentHandler) GetByID(c *fiber.Ctx) error {
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

func (h *StudentHandler) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID mahasiswa tidak valid")
	}

	var req model.Student
	if err := c.BodyParser(&req); err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "Format JSON tidak valid")
	}

	err = h.Service.UpdateStudent(c.Context(), id, req)
	if err != nil {
		return helper.Fail(c, fiber.StatusNotFound, "Gagal memperbarui: Data mahasiswa tidak ditemukan")
	}

	return helper.OK(c, fiber.StatusOK, "Berhasil memperbarui data mahasiswa", nil, nil)
}

func (h *StudentHandler) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID mahasiswa tidak valid")
	}

	err = h.Service.DeleteStudent(c.Context(), id)
	if err != nil {
		return helper.Fail(c, fiber.StatusNotFound, "Gagal menghapus: Data mahasiswa tidak ditemukan")
	}

	return helper.OK(c, fiber.StatusOK, "Berhasil menghapus data mahasiswa", nil, nil)
}