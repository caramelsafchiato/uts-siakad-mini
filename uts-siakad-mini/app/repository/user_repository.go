package repository

import (
	"context"
	"errors"
	"uts-siakad-mini/app/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

// FindByEmail mencari user berdasarkan email untuk keperluan login
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := `SELECT id, email, password, role FROM users WHERE email = $1`
	
	err := r.DB.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}
	return &user, nil
}