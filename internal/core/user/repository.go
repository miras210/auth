package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *Model) error
	Update(ctx context.Context, user *Model) error
	GetByEmailOrUsername(ctx context.Context, emailOrUsername string) (User, error)
	IsUnique(ctx context.Context, emailOrUserName string) (bool, error)
}
