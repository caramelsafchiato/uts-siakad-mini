package handler

import (
	"strconv"
	"uts-siakad-mini/app/model"
	"uts-siakad-mini/app/service"
	"uts-siakad-mini/helper"

	"github.com/gofiber/fiber/v2"
)

type CourseHandler struct {
	Service *service.CourseService
}

func NewCourseHandler(service *service.CourseService) *CourseHandler {
	return &CourseHandler{Service: service}
}

func (h *CourseHandler) GetAll(c *fiber.Ctx) error {
	courses, err := h.Service.GetAllCourses(c.Context())
	if err != nil {
		return helper.Fail(c, fiber.StatusInternalServerError, "Gagal mengambil data mata kuliah")
	}
	return helper.OK(c, fiber.StatusOK, "Berhasil mengambil daftar mata kuliah", courses, nil)
}

func (h *CourseHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	course, err := h.Service.GetCourseByID(c.Context(), id)
	if err != nil {
		return helper.Fail(c, fiber.StatusNotFound, "Mata kuliah tidak ditemukan")
	}
	return helper.OK(c, fiber.StatusOK, "Berhasil mengambil detail mata kuliah", course, nil)
}

func (h *CourseHandler) Create(c *fiber.Ctx) error {
	var req model.Course
	if err := c.BodyParser(&req); err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "Format JSON tidak valid")
	}

	if err := h.Service.CreateCourse(c.Context(), req); err != nil {
		return helper.Fail(c, fiber.StatusInternalServerError, "Gagal menambahkan mata kuliah")
	}
	return helper.OK(c, fiber.StatusCreated, "Berhasil menambahkan mata kuliah", nil, nil)
}

func (h *CourseHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	var req model.Course
	if err := c.BodyParser(&req); err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "Format JSON tidak valid")
	}

	if err := h.Service.UpdateCourse(c.Context(), id, req); err != nil {
		return helper.Fail(c, fiber.StatusNotFound, "Mata kuliah tidak ditemukan atau gagal diperbarui")
	}
	return helper.OK(c, fiber.StatusOK, "Berhasil memperbarui mata kuliah", nil, nil)
}

func (h *CourseHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	if err := h.Service.DeleteCourse(c.Context(), id); err != nil {
		return helper.Fail(c, fiber.StatusNotFound, "Mata kuliah tidak ditemukan atau gagal dihapus")
	}
	return helper.OK(c, fiber.StatusOK, "Berhasil menghapus mata kuliah", nil, nil)
}