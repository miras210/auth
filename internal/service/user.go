package service

import (
	"auth/internal/core/user"
	"context"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo user.Repository
	logger   *zap.SugaredLogger
}

func NewUserService(userRepo user.Repository, logger *zap.SugaredLogger) user.Service {
	return UserService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Create TODO: do something with errors map
func (u UserService) Create(ctx context.Context, creds *user.NewUser) error {
	if errorsMap := creds.Validate(); errorsMap != nil {
		return NewValidationError("userService.Create() validation error", errorsMap)
	}

	login := GetLogin(creds.Email, creds.Username)
	u.logger.Infof("login: %s", login)
	isUnique, err := u.userRepo.IsUnique(ctx, login)
	if err != nil {
		u.logger.Errorf("error occurred while checking for uniqueness of user: %v", err)
		return err
	}

	if !isUnique {
		return ErrLoginIsNotUnique
	}

	passHash, err := hashAndSalt([]byte(creds.Password))
	if err != nil {
		return err
	}

	model := user.Model{
		Login:    &login,
		PassHash: &passHash,
	}

	return u.userRepo.Create(ctx, &model)
}

// Update TODO: set user id from context
// Update TODO: do something with errors map
func (u UserService) Update(ctx context.Context, creds *user.UpdateUser) error {
	if errorsMap := creds.Validate(); errorsMap != nil {
		return NewValidationError("userService.Update() validation error", errorsMap)
	}

	login := GetLogin(creds.Email, creds.Username)

	model := user.Model{
		Login: &login,
	}

	return u.userRepo.Update(ctx, &model)
}

func (u UserService) GetByEmailOrUsername(ctx context.Context, emailOrUsername string) (user.User, error) {
	return u.userRepo.GetByEmailOrUsername(ctx, emailOrUsername)
}
