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
	// --- INISIALISASI ---
	
	// Auth
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// Student
	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	// Course
	courseRepo := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	// --- ROUTING ---
	
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
	protected.Put("/students/:id", studentHandler.Update)
	protected.Delete("/students/:id", studentHandler.Delete)

	// Endpoint Mata Kuliah
	protected.Get("/courses", courseHandler.GetAll)// Endpoint Mata Kuliah
	protected.Get("/courses", courseHandler.GetAll)
	protected.Get("/courses/:id", courseHandler.GetByID)
	protected.Post("/courses", courseHandler.Create)
	protected.Put("/courses/:id", courseHandler.Update)
	protected.Delete("/courses/:id", courseHandler.Delete)
}