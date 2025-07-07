package repository

import (
	"database/sql"
	"go-29/internal/data/entity"

	"go.uber.org/zap"
)

type UserRepository interface {
	Create(user *entity.User) error
}

type userRepositoryImpl struct {
	DB  *sql.DB
	Log *zap.Logger
}

func NewUserRepository(db *sql.DB, log *zap.Logger) UserRepository {
	return &userRepositoryImpl{
		DB:  db,
		Log: log,
	}
}

func (r *userRepositoryImpl) Create(user *entity.User) error {
	query := `
		INSERT INTO users (name, email, password, photo, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`
	return r.DB.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Photo,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}
