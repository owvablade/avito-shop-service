package purchasehandler

import (
	"avito-shop-service/internal/errors"
	"avito-shop-service/internal/usecase/purchase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Interface interface {
	CreatePurchase(ctx *gin.Context)
}

type Handler struct {
	uc purchase.Interface
}

func New(uc purchase.Interface) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) CreatePurchase(ctx *gin.Context) {
	merchName := ctx.Param("id")
	if merchName == "" {
		_ = ctx.Error(errors.ErrMerchNameNotFound)
		return
	}

	userID := ctx.GetInt("user_id")
	if userID == 0 {
		_ = ctx.Error(errors.ErrInternalServer)
		return
	}

	if err := h.uc.CreatePurchase(ctx, userID, merchName); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
