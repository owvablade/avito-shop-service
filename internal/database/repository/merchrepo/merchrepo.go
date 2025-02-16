package merchrepo

import (
	"avito-shop-service/internal/database/model"
	"context"
	"github.com/jmoiron/sqlx"
)

type Interface interface {
	GetMerchItemByMerchName(ctx context.Context, name string) (*model.MerchItem, error)
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetMerchItemByMerchName(ctx context.Context, merchName string) (*model.MerchItem, error) {
	query := "SELECT * FROM merch WHERE name = $1;"

	item := &model.MerchItem{}
	if err := r.db.GetContext(ctx, item, query, merchName); err != nil {
		return nil, err
	}

	return item, nil
}
