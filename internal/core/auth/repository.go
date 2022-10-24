package auth

import (
	"context"
)

type TokenRepository interface {
	Create(ctx context.Context, token *Token) error
	Delete(ctx context.Context, id string) error
	GetByToken(ctx context.Context, tokenString string) (Token, error)
}
