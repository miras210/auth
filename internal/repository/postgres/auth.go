package postgres

import (
	"auth/internal/core/auth"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type TokenRepository struct {
	logger *zap.SugaredLogger
	db     *pgxpool.Pool
}

func NewTokenRepository(db *pgxpool.Pool, logger *zap.SugaredLogger) auth.TokenRepository {
	return TokenRepository{
		db:     db,
		logger: logger,
	}
}

func (t TokenRepository) Create(ctx context.Context, token *auth.Token) error {
	query := `INSERT INTO tokens (user_id, refresh_token, expiration_time) VALUES ($1, $2, $3)`

	if _, err := t.db.Exec(ctx, query, &token.UserID, &token.TokenValue, &token.ExpiresAt); err != nil {
		return err
	}

	return nil
}

func (t TokenRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tokens WHERE id = $1`

	if _, err := t.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}

func (t TokenRepository) GetByToken(ctx context.Context, tokenString string) (auth.Token, error) {
	var token auth.Token
	query := `SELECT id, user_id, refresh_token, expiration_time FROM tokens WHERE refresh_token=$1`

	err := t.db.QueryRow(ctx, query, tokenString).Scan(&token.ID, &token.UserID, &token.TokenValue, &token.ExpiresAt)
	if err != nil {
		return auth.Token{}, err
	}

	return token, nil
}
