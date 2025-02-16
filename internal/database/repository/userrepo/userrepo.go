package userrepo

import (
	"avito-shop-service/internal/database/model"
	"context"
	"github.com/jmoiron/sqlx"
)

type Interface interface {
	CreateUser(ctx context.Context, username string, password string) (int, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserCoinsByUserID(ctx context.Context, userID int) (int, error)
	AddCoinsToUserByUserID(ctx context.Context, userID int, amount int) error
	SubtractCoinsFromUserByUserID(ctx context.Context, userID int, amount int) error
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, username string, password string) (int, error) {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id;"

	var id int
	err := r.db.QueryRowContext(ctx, query, username, password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	query := "SELECT * FROM users WHERE username = $1;"

	user := &model.User{}
	if err := r.db.GetContext(ctx, user, query, username); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserCoinsByUserID(ctx context.Context, userID int) (int, error) {
	query := "SELECT coins FROM users WHERE id = $1;"

	var coins int
	if err := r.db.GetContext(ctx, &coins, query, userID); err != nil {
		return 0, err
	}

	return coins, nil
}

func (r *Repository) AddCoinsToUserByUserID(ctx context.Context, userID int, amount int) error {
	query := "UPDATE users SET coins = coins + $1 WHERE id = $2;"
	if _, err := r.db.ExecContext(ctx, query, amount, userID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) SubtractCoinsFromUserByUserID(ctx context.Context, userID int, amount int) error {
	query := "UPDATE users SET coins = coins - $1 WHERE id = $2;"
	if _, err := r.db.ExecContext(ctx, query, amount, userID); err != nil {
		return err
	}

	return nil
}
