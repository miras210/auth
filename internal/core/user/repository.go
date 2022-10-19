package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *NewUser) error
	Update(ctx context.Context, user *UpdateUser) error
	GetByEmailOrUsername(ctx context.Context, emailOrUsername string) (User, error)
}
