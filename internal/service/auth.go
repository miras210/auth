package service

import (
	"auth/internal/core/auth"
	"auth/internal/core/user"
)

type AuthService struct {
	userRepo  user.Repository
	tokenRepo auth.TokenRepository
}

func NewAuthService(userRepo user.Repository, tokenRepo auth.TokenRepository) auth.Service {
	return AuthService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}
