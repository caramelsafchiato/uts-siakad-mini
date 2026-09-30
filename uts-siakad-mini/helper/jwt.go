package helper

import (
	"time"
	"uts-siakad-mini/config"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken membuat JWT token yang berisi user_id dan role
func GenerateToken(userID int, role string) (string, error) {
	secret := config.Get("JWT_SECRET", "rahasia_siakad_uts_2026")
	
	// Isi data token (claims)
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}