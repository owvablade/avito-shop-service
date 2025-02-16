package handler

import (
	"avito-shop-service/internal/handler/purchasehandler"
	"avito-shop-service/internal/handler/transactionhandler"
	"avito-shop-service/internal/handler/userhandler"
	"avito-shop-service/internal/usecase"
)

type Handler struct {
	UserHandler        userhandler.Interface
	PurchaseHandler    purchasehandler.Interface
	TransactionHandler transactionhandler.Interface
}

func New(uc *usecase.UseCase) *Handler {
	return &Handler{
		UserHandler:        userhandler.New(uc.UserUseCase),
		PurchaseHandler:    purchasehandler.New(uc.PurchaseUseCase),
		TransactionHandler: transactionhandler.New(uc.TransactionUseCase),
	}
}
