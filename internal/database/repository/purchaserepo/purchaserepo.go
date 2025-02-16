package purchaserepo

import (
	"avito-shop-service/internal/database/model"
	"context"
	"github.com/jmoiron/sqlx"
)

type Interface interface {
	CreatePurchase(ctx context.Context, userID int, merchID int) error
	GetPurchasesByUserID(ctx context.Context, userID int) ([]model.Inventory, error)
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatePurchase(ctx context.Context, userID int, merchID int) error {
	query := "INSERT INTO purchases (user_id, merch_id) VALUES ($1, $2)"
	if _, err := r.db.ExecContext(ctx, query, userID, merchID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetPurchasesByUserID(ctx context.Context, userID int) ([]model.Inventory, error) {
	query := `
SELECT m.name AS "type", COUNT(m.id) AS "quantity"
FROM purchases p
JOIN merch m ON p.merch_id = m.id
WHERE p.user_id = $1
GROUP BY m.id;`

	inventory := make([]model.Inventory, 0)
	if err := r.db.SelectContext(ctx, &inventory, query, userID); err != nil {
		return nil, err
	}

	return inventory, nil
}
