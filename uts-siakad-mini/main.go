package main

import (
	"log"
	"uts-siakad-mini/config"
	"uts-siakad-mini/helper"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Baca konfigurasi dari .env
	config.LoadEnv()

	// 2. Buka koneksi ke Database PostgreSQL
	dbPool := config.ConnectDB()
	defer dbPool.Close()

	// 3. Inisialisasi aplikasi Fiber
	app := fiber.New(fiber.Config{
		AppName: "SIAKAD Mini UTS",
	})

	// 4. Endpoint sekadar untuk ngetes server jalan atau nggak
	app.Get("/", func(c *fiber.Ctx) error {
		return helper.OK(c, fiber.StatusOK, "Server SIAKAD Mini berjalan lancar!", nil, nil)
	})

	// 5. Jalankan server di port yang ada di .env (3000)
	port := config.Get("APP_PORT", "3000")
	log.Fatal(app.Listen(":" + port))
}