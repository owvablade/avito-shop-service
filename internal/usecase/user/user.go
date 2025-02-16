package user

import (
	"avito-shop-service/internal/database/model"
	"avito-shop-service/internal/database/repository/purchaserepo"
	"avito-shop-service/internal/database/repository/transactionrepo"
	"avito-shop-service/internal/database/repository/userrepo"
	"avito-shop-service/internal/database/txmanager"
	"avito-shop-service/internal/dto"
	"avito-shop-service/internal/errors"
	"avito-shop-service/internal/service/crypto"
	"avito-shop-service/internal/service/jwtservice"
	"context"
)

type Interface interface {
	GetUserInfo(ctx context.Context, userID int) (*model.UserInfo, error)
	CreateOrAuthUser(ctx context.Context, request *dto.AuthRequest) (*dto.AuthResponse, error)
}

type UseCase struct {
	jwtService      jwtservice.Interface
	cryptoService   crypto.Interface
	txm             txmanager.TxManager
	userRepo        userrepo.Interface
	purchaseRepo    purchaserepo.Interface
	transactionRepo transactionrepo.Interface
}

func New(
	jwtService jwtservice.Interface,
	cryptoService crypto.Interface,
	txm txmanager.TxManager,
	userRepo userrepo.Interface,
	purchaseRepo purchaserepo.Interface,
	transactionRepo transactionrepo.Interface,
) *UseCase {
	return &UseCase{
		jwtService:      jwtService,
		cryptoService:   cryptoService,
		txm:             txm,
		userRepo:        userRepo,
		purchaseRepo:    purchaseRepo,
		transactionRepo: transactionRepo,
	}
}

func (u *UseCase) GetUserInfo(ctx context.Context, userID int) (out *model.UserInfo, err error) {
	err = u.txm.ReadOnly(ctx, func(ctx context.Context) error {
		out = &model.UserInfo{}
		out.Coins, err = u.userRepo.GetUserCoinsByUserID(ctx, userID)
		if err != nil {
			return errors.ErrUserNotFound
		}

		out.Inventory, err = u.purchaseRepo.GetPurchasesByUserID(ctx, userID)
		if err != nil {
			return errors.ErrInternalServer
		}

		out.CoinHistory.Sent, err = u.transactionRepo.GetSentTransactionsFromUserID(ctx, userID)
		if err != nil {
			return errors.ErrInternalServer
		}

		out.CoinHistory.Received, err = u.transactionRepo.GetReceivedTransactionsToUserID(ctx, userID)
		if err != nil {
			return errors.ErrInternalServer
		}

		return nil
	})

	return out, err
}

func (u *UseCase) CreateOrAuthUser(ctx context.Context, request *dto.AuthRequest) (*dto.AuthResponse, error) {
	user, err := u.userRepo.GetUserByUsername(ctx, request.Username)
	if err == nil {
		if u.cryptoService.CompareHashAndPassword(user.Password, request.Password) != nil {
			return nil, errors.ErrWrongPassword
		}
		signedToken, err := u.jwtService.GenerateSignedTokenFromUserId(user.ID)
		if err != nil {
			return nil, errors.ErrInternalServer
		}
		return &dto.AuthResponse{Token: signedToken}, nil
	}

	hashedPassword, err := u.cryptoService.HashPassword(request.Password)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	id, err := u.userRepo.CreateUser(ctx, request.Username, hashedPassword)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	signedToken, err := u.jwtService.GenerateSignedTokenFromUserId(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return &dto.AuthResponse{Token: signedToken}, nil
}
