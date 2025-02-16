package transactionrepo

import (
	"avito-shop-service/internal/database/model"
	"context"
	"github.com/jmoiron/sqlx"
)

type Interface interface {
	CreateTransaction(ctx context.Context, fromID int, toID int, amount int) error
	GetSentTransactionsFromUserID(ctx context.Context, userID int) ([]model.SentTransaction, error)
	GetReceivedTransactionsToUserID(ctx context.Context, userID int) ([]model.ReceivedTransaction, error)
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateTransaction(ctx context.Context, fromID int, toID int, amount int) error {
	query := "INSERT INTO transactions (from_user_id, to_user_id, amount) VALUES ($1, $2, $3);"
	if _, err := r.db.ExecContext(ctx, query, fromID, toID, amount); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetSentTransactionsFromUserID(ctx context.Context, userID int) ([]model.SentTransaction, error) {
	query := `
SELECT username, amount 
FROM transactions t 
JOIN users u ON t.to_user_id = u.id
WHERE from_user_id = $1`

	transactions := make([]model.SentTransaction, 0)
	if err := r.db.SelectContext(ctx, &transactions, query, userID); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *Repository) GetReceivedTransactionsToUserID(
	ctx context.Context,
	userID int,
) ([]model.ReceivedTransaction, error) {
	query := `
SELECT username, amount 
FROM transactions t 
JOIN users u ON t.from_user_id = u.id
WHERE to_user_id = $1;`

	transactions := make([]model.ReceivedTransaction, 0)
	if err := r.db.SelectContext(ctx, &transactions, query, userID); err != nil {
		return nil, err
	}

	return transactions, nil
}
