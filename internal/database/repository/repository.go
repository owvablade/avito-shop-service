package repository

import (
	"avito-shop-service/internal/database/repository/merchrepo"
	"avito-shop-service/internal/database/repository/purchaserepo"
	"avito-shop-service/internal/database/repository/transactionrepo"
	"avito-shop-service/internal/database/repository/userrepo"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepo        userrepo.Interface
	MerchRepo       merchrepo.Interface
	PurchaseRepo    purchaserepo.Interface
	TransactionRepo transactionrepo.Interface
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo:        userrepo.New(db),
		MerchRepo:       merchrepo.New(db),
		PurchaseRepo:    purchaserepo.New(db),
		TransactionRepo: transactionrepo.New(db),
	}
}
