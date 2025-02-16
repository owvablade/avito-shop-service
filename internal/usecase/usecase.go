package usecase

import (
	"avito-shop-service/internal/database/repository"
	"avito-shop-service/internal/database/txmanager"
	"avito-shop-service/internal/service/crypto"
	"avito-shop-service/internal/service/jwtservice"
	"avito-shop-service/internal/usecase/purchase"
	"avito-shop-service/internal/usecase/transaction"
	"avito-shop-service/internal/usecase/user"
)

type UseCase struct {
	UserUseCase        user.Interface
	PurchaseUseCase    purchase.Interface
	TransactionUseCase transaction.Interface
}

func New(
	jwt jwtservice.Interface,
	crypto crypto.Interface,
	txm txmanager.TxManager,
	repos *repository.Repository,
) *UseCase {
	return &UseCase{
		UserUseCase:        user.New(jwt, crypto, txm, repos.UserRepo, repos.PurchaseRepo, repos.TransactionRepo),
		PurchaseUseCase:    purchase.New(txm, repos.UserRepo, repos.MerchRepo, repos.PurchaseRepo),
		TransactionUseCase: transaction.New(txm, repos.UserRepo, repos.TransactionRepo),
	}
}
