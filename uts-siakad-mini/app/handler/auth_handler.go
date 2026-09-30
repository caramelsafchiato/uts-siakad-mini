package handler

import (
	"uts-siakad-mini/app/model"
	"uts-siakad-mini/app/service"
	"uts-siakad-mini/helper"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest

	// Tangkap body JSON dari Postman
	if err := c.BodyParser(&req); err != nil {
		return helper.Fail(c, fiber.StatusUnprocessableEntity, "Format request tidak valid")
	}

	// Validasi dasar (422)
	if req.Email == "" || len(req.Password) < 8 {
		return helper.Fail(c, fiber.StatusUnprocessableEntity, "Email wajib diisi dan password minimal 8 karakter")
	}

	// Panggil logika bisnis (Service)
	res, err := h.Service.Login(c.Context(), req)
	if err != nil {
		// 401 Unauthorized jika email/password salah
		return helper.Fail(c, fiber.StatusUnauthorized, err.Error())
	}

	// Kembalikan 200 OK beserta data token
	return helper.OK(c, fiber.StatusOK, "Login berhasil", res, nil)
}