package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv membaca file .env
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: file .env tidak ditemukan")
	}
}

// Get mengambil nilai dari .env, kalau kosong pakai nilai fallback
func Get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}