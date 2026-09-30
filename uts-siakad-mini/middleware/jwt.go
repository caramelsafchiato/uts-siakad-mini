package middleware

import (
	"strings"
	"uts-siakad-mini/config"
	"uts-siakad-mini/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Protected adalah satpam penjaga endpoint
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Ambil token dari header "Authorization"
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return helper.Fail(c, fiber.StatusUnauthorized, "Akses ditolak: Token tidak ditemukan atau format salah")
		}

		// 2. Potong tulisan "Bearer " untuk mengambil token aslinya
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		secret := config.Get("JWT_SECRET", "rahasia_siakad_uts_2026")

		// 3. Bongkar dan validasi token
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(secret), nil
		})

		// 4. Jika token salah, palsu, atau kadaluarsa
		if err != nil || !token.Valid {
			return helper.Fail(c, fiber.StatusUnauthorized, "Akses ditolak: Token tidak valid atau sudah kadaluarsa")
		}

		// 5. Ambil data user_id dan role dari dalam token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return helper.Fail(c, fiber.StatusUnauthorized, "Gagal memproses data token")
		}

		// 6. Simpan data tersebut agar bisa dibaca oleh Handler selanjutnya
		c.Locals("user_id", claims["user_id"])
		c.Locals("role", claims["role"])

		// Lolos sensor! Silakan lanjut ke tujuan utama
		return c.Next()
	}
}