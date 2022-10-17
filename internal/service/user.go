package service

import "auth/internal/core/user"

type UserService struct {
	userRepo user.Repository
}

func NewUserService(userRepo user.Repository) user.Service {
	return UserService{
		userRepo: userRepo,
	}
}
