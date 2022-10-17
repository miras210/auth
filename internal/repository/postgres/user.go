package postgres

import (
	"auth/internal/core/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) user.Repository {
	return UserRepository{
		db: db,
	}
}
