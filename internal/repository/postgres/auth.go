package postgres

import (
	"auth/internal/core/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TokenRepository struct {
	db *pgxpool.Pool
}

func NewTokenRepository(db *pgxpool.Pool) auth.TokenRepository {
	return TokenRepository{
		db: db,
	}
}
