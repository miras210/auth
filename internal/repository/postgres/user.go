package postgres

import (
	"auth/internal/core/user"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserRepository struct {
	db     *pgxpool.Pool
	logger *zap.SugaredLogger
}

func NewUserRepository(db *pgxpool.Pool, logger *zap.SugaredLogger) user.Repository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

// Create TODO: change
func (u UserRepository) Create(ctx context.Context, user *user.Model) error {
	query := `INSERT INTO users (login, passhash) VALUES ($1, $2) RETURNING id`
	var userId string

	if err := u.db.QueryRow(ctx, query, user.Login,
		user.PassHash).Scan(&userId); err != nil {
		u.logger.Errorf("error occurred while querying to DB: %s", err.Error())
		return err
	}
	u.logger.Infof("created user with ID: %s", userId)

	return nil
}

func (u UserRepository) Update(ctx context.Context, user *user.Model) error {
	query := `UPDATE users SET login=COALESCE($1, login) WHERE id=$3 RETURNING id`
	var userId string

	if err := u.db.QueryRow(ctx, query, user.Login, user.ID).
		Scan(&userId); err != nil {
		u.logger.Errorf("error occurred while querying to DB: %s", err.Error())
		return err
	}
	u.logger.Infof("updated user with ID: %s", userId)

	return nil
}

// GetByEmailOrUsername TODO: maybe changed to 2 methods for email and username
func (u UserRepository) GetByEmailOrUsername(ctx context.Context, emailOrUsername string) (user.User, error) {
	query := `SELECT id, login, passhash FROM users WHERE login=$1`

	var retrievedUser user.User

	row := u.db.QueryRow(ctx, query, emailOrUsername)

	if err := row.Scan(&retrievedUser.ID, &retrievedUser.Email, &retrievedUser.PassHash); err != nil {
		u.logger.Errorf("error occurred while querying to DB: %s", err.Error())
		return retrievedUser, err
	}
	u.logger.Infof("retrieved user with ID: %s", retrievedUser.ID)

	return retrievedUser, nil
}

func (u UserRepository) IsUnique(ctx context.Context, emailOrUserName string) (bool, error) {
	query := `SELECT count(*) AS all FROM users WHERE login=$1`
	var count int

	row := u.db.QueryRow(ctx, query, emailOrUserName)

	if err := row.Scan(&count); err != nil {
		u.logger.Errorf("error occurred while querying to DB: %s", err.Error())
		return false, err
	}

	if count > 0 {
		return false, nil
	}

	return true, nil
}
