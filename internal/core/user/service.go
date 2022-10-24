package user

import "context"

type Service interface {
	Create(ctx context.Context, creds *NewUser) error
	Update(ctx context.Context, creds *UpdateUser) error
	GetByEmailOrUsername(ctx context.Context, emailOrUsername string) (User, error)
}
