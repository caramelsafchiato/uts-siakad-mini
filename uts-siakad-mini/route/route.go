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

// ... (bagian atas biarkan tetap sama)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {
	// Inisialisasi Repo & Service Auth
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// Inisialisasi Repo, Service, & Handler Student
	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	api := app.Group("/api/v1")
	
	// Area Bebas Akses
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)

	// Area TerPROTEKSI (Wajib bawa Token JWT)
	protected := api.Group("/", middleware.Protected())

	protected.Get("/me", func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		role := c.Locals("role")
		return helper.OK(c, fiber.StatusOK, "Berhasil masuk ke area rahasia!", fiber.Map{
			"user_id": userID,
			"role":    role,
		}, nil)
	})

	// Endpoint Mahasiswa
	protected.Get("/students", studentHandler.GetAll)
	protected.Get("/students/:id", studentHandler.GetByID)
}