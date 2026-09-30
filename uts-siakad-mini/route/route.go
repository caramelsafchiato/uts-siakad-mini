package route

import (
	"uts-siakad-mini/app/handler"
	"uts-siakad-mini/app/repository"
	"uts-siakad-mini/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {
	// 1. Inisialisasi Repository
	userRepo := repository.NewUserRepository(db)

	// 2. Inisialisasi Service
	authService := service.NewAuthService(userRepo)

	// 3. Inisialisasi Handler
	authHandler := handler.NewAuthHandler(authService)

	// 4. Mendaftarkan Grup API
	api := app.Group("/api/v1")
	auth := api.Group("/auth")

	// 5. Endpoint Login (POST)
	auth.Post("/login", authHandler.Login)
}