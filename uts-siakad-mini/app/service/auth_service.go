package service

import (
	"context"
	"errors"
	"uts-siakad-mini/app/model"
	"uts-siakad-mini/app/repository"
	"uts-siakad-mini/helper"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	// 1. Cari user di database berdasarkan email
	user, err := s.Repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	// 2. Bandingkan password input dengan hash di database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	// 3. Buat Token JWT
	token, err := helper.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, errors.New("sistem gagal membuat token")
	}

	// 4. Susun balasan sesuai format dari dosen
	res := &model.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   86400, // Berlaku 24 jam (dalam detik)
		User:        *user, // Password tidak akan tampil karena di entity.go sudah diberi tag json:"-"
	}

	return res, nil
}