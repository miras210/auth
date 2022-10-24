package auth

import (
	"auth/internal/core/user"
	"context"
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	Auth(ctx context.Context, creds *user.SignIn) (*TokenPair, error)
	Refresh(ctx context.Context, refresh string) (*TokenPair, error)
	Logout(ctx context.Context) error
	ValidateAccess(ctx context.Context, accessString string) (jwt.StandardClaims, error)
}
