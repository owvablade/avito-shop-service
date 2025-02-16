package transactionhandler

import (
	"avito-shop-service/internal/dto"
	"avito-shop-service/internal/errors"
	"avito-shop-service/internal/usecase/transaction"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Interface interface {
	CreateTransaction(ctx *gin.Context)
}

type Handler struct {
	uc transaction.Interface
}

func New(uc transaction.Interface) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) CreateTransaction(ctx *gin.Context) {
	var sendCoinReq dto.SendCoinRequest
	if err := ctx.ShouldBindJSON(&sendCoinReq); err != nil {
		_ = ctx.Error(errors.ErrRequiredFieldNotSet)
		return
	}

	fromID := ctx.GetInt("user_id")
	if fromID == 0 {
		_ = ctx.Error(errors.ErrInternalServer)
		return
	}

	if err := h.uc.CreateTransaction(ctx, fromID, &sendCoinReq); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
