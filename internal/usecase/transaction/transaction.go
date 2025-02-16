package transaction

import (
	"avito-shop-service/internal/database/model"
	"avito-shop-service/internal/database/repository/transactionrepo"
	"avito-shop-service/internal/database/repository/userrepo"
	"avito-shop-service/internal/database/txmanager"
	"avito-shop-service/internal/dto"
	"avito-shop-service/internal/errors"
	"context"
)

type Interface interface {
	CreateTransaction(ctx context.Context, fromID int, request *dto.SendCoinRequest) error
}

type UseCase struct {
	txm             txmanager.TxManager
	userRepo        userrepo.Interface
	transactionRepo transactionrepo.Interface
}

func New(
	txm txmanager.TxManager,
	userRepo userrepo.Interface,
	transactionRepo transactionrepo.Interface,
) *UseCase {
	return &UseCase{txm: txm, userRepo: userRepo, transactionRepo: transactionRepo}
}

func (u *UseCase) CreateTransaction(ctx context.Context, fromID int, request *dto.SendCoinRequest) (err error) {
	if request.Amount <= 0 {
		return errors.ErrNegativeOrZeroAmount
	}

	err = u.txm.ReadWrite(ctx, func(ctx context.Context) error {
		var user *model.User
		user, err = u.userRepo.GetUserByUsername(ctx, request.ToUser)
		if err != nil {
			return errors.ErrUserNotFound
		}

		if user.ID == fromID {
			return errors.ErrTransferToYourself
		}

		var coins int
		coins, err = u.userRepo.GetUserCoinsByUserID(ctx, fromID)
		if err != nil {
			return errors.ErrUserNotFound
		}

		if coins < request.Amount {
			return errors.ErrInsufficientCoins
		}

		if err = u.userRepo.SubtractCoinsFromUserByUserID(ctx, fromID, request.Amount); err != nil {
			return errors.ErrInternalServer
		}

		if err = u.userRepo.AddCoinsToUserByUserID(ctx, user.ID, request.Amount); err != nil {
			return errors.ErrInternalServer
		}

		if err = u.transactionRepo.CreateTransaction(ctx, fromID, user.ID, request.Amount); err != nil {
			return errors.ErrInternalServer
		}

		return nil
	})

	return err
}
