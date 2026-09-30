package route

import (
	"uts-siakad-mini/app/handler"
	"uts-siakad-mini/app/repository"
	"uts-siakad-mini/app/service"
	"uts-siakad-mini/helper"
	"uts-siakad-mini/middleware" 

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {
	// ... (Inisialisasi Repo, Service, Handler tetap sama seperti sebelumnya) ...
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	api := app.Group("/api/v1")
	
	// --- AREA BEBAS AKSES ---
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)

	// --- AREA TERPROTEKSI (Harus bawa Token) ---
	protected := api.Group("/", middleware.Protected())

	// Endpoint sementara untuk ngetes middleware
	protected.Get("/me", func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		role := c.Locals("role")

		return helper.OK(c, fiber.StatusOK, "Berhasil masuk ke area rahasia!", fiber.Map{
			"user_id": userID,
			"role":    role,
		}, nil)
	})
}