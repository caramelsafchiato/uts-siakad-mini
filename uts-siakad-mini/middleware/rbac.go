package middleware

import (
	"uts-siakad-mini/helper"

	"github.com/gofiber/fiber/v2"
)

// RequireRole memblokir akses jika role dari token tidak sesuai
func RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil role hasil ekstrak dari JWT (diset oleh middleware Protected)
		userRole := c.Locals("role")
		if userRole == nil {
			return helper.Fail(c, fiber.StatusUnauthorized, "Role tidak ditemukan pada token")
		}

		roleStr := userRole.(string)

		// Cek apakah role user cocok dengan salah satu role yang diizinkan
		for _, role := range allowedRoles {
			if roleStr == role {
				return c.Next() // Lolos! Lanjut ke handler
			}
		}

		// Jika tidak cocok, tolak aksesnya (403 Forbidden)
		return helper.Fail(c, fiber.StatusForbidden, "Akses Ditolak: Fitur ini hanya untuk Admin")
	}
}