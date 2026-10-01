package handler

import (
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