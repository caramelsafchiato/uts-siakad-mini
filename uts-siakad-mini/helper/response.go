package helper

import (
	"uts-siakad-mini/app/model"

	"github.com/gofiber/fiber/v2"
)

// OK untuk balasan sukses 200/201
func OK(c *fiber.Ctx, status int, message string, data any, meta *model.Meta) error {
	return c.Status(status).JSON(model.Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// Fail untuk balasan error (400, 401, 404, 500, dll)
func Fail(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(model.Response{
		Success: false,
		Message: message,
	})
}

// FailValidation khusus untuk error 422 (validasi form/input)
func FailValidation(c *fiber.Ctx, errors map[string][]string) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(model.Response{
		Success: false,
		Message: "Validasi gagal",
		Errors:  errors,
	})
}