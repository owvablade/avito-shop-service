package userhandler

import (
	"avito-shop-service/internal/dto"
	"avito-shop-service/internal/errors"
	"avito-shop-service/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Interface interface {
	CreateOrAuthUser(ctx *gin.Context)
	GetUserInfo(ctx *gin.Context)
}

type Handler struct {
	uc user.Interface
}

func New(uc user.Interface) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) CreateOrAuthUser(ctx *gin.Context) {
	var in dto.AuthRequest
	if err := ctx.ShouldBindJSON(&in); err != nil {
		_ = ctx.Error(errors.ErrRequiredFieldNotSet)
		return
	}

	response, err := h.uc.CreateOrAuthUser(ctx.Request.Context(), &in)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &response)
}

func (h *Handler) GetUserInfo(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	if userID == 0 {
		_ = ctx.Error(errors.ErrInternalServer)
		return
	}

	userInfo, err := h.uc.GetUserInfo(ctx, userID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}
