package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	// Susun URL koneksi dari variabel env
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		Get("DB_USER", "postgres"),
		Get("DB_PASSWORD", ""),
		Get("DB_HOST", "localhost"),
		Get("DB_PORT", "5432"),
		Get("DB_NAME", "siakad_mini"),
		Get("DB_SSLMODE", "disable"),
	)

	// Persiapan konfigurasi pool
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Gagal membaca konfigurasi database: %v", err)
	}

	// Batas waktu koneksi
	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("Gagal membuat connection pool: %v", err)
	}

	// Tes Ping ke database (pastikan server DB nyala)
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		log.Fatalf("Gagal terhubung ke database PostgreSQL: %v", err)
	}

	fmt.Println("Berhasil terhubung ke database PostgreSQL!")
	return pool
}