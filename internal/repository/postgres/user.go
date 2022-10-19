package postgres

import (
	"auth/internal/core/user"
	"auth/internal/repository"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) user.Repository {
	return UserRepository{
		db: db,
	}
}

// Create TODO: change
func (u UserRepository) Create(ctx context.Context, user *user.NewUser) error {
	query := `INSERT INTO users (email, username, passhash) VALUES ($1, $2, $3)`
	_, err := u.db.Exec(ctx, query,
		user.Email,
		repository.NewNullString(user.Username),
		repository.NewNullString(user.Password))
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Update(ctx context.Context, user *user.UpdateUser) error {
	//TODO implement me
	panic("implement me")
}

// GetByEmailOrUsername TODO: change
func (u UserRepository) GetByEmailOrUsername(ctx context.Context, emailOrUsername string) (user.User, error) {
	query := `SELECT id, username, email, passhash FROM users WHERE username=$1 OR email=$2`

	var user user.User

	row := u.db.QueryRow(ctx, query, emailOrUsername, emailOrUsername)

	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PassHash); err != nil {
		return user, err
	}

	return user, nil
}
