package purchase

import (
	"avito-shop-service/internal/database/repository/merchrepo"
	"avito-shop-service/internal/database/repository/purchaserepo"
	"avito-shop-service/internal/database/repository/userrepo"
	"avito-shop-service/internal/database/txmanager"
	"avito-shop-service/internal/errors"
	"context"
)

type Interface interface {
	CreatePurchase(ctx context.Context, userID int, merchName string) error
}

type UseCase struct {
	txm          txmanager.TxManager
	userRepo     userrepo.Interface
	merchRepo    merchrepo.Interface
	purchaseRepo purchaserepo.Interface
}

func New(
	txm txmanager.TxManager,
	userRepo userrepo.Interface,
	merchRepo merchrepo.Interface,
	purchaseRepo purchaserepo.Interface,
) *UseCase {
	return &UseCase{txm: txm, userRepo: userRepo, merchRepo: merchRepo, purchaseRepo: purchaseRepo}
}

func (u *UseCase) CreatePurchase(ctx context.Context, userID int, merchName string) (err error) {
	err = u.txm.ReadWrite(ctx, func(ctx context.Context) error {
		coins, err := u.userRepo.GetUserCoinsByUserID(ctx, userID)
		if err != nil {
			return errors.ErrUserNotFound
		}

		merchItem, err := u.merchRepo.GetMerchItemByMerchName(ctx, merchName)
		if err != nil {
			return errors.ErrMerchNotFound
		}

		if coins < merchItem.Price {
			return errors.ErrInsufficientCoins
		}

		err = u.userRepo.SubtractCoinsFromUserByUserID(ctx, userID, merchItem.Price)
		if err != nil {
			return errors.ErrInternalServer
		}

		err = u.purchaseRepo.CreatePurchase(ctx, userID, merchItem.ID)
		if err != nil {
			return errors.ErrInternalServer
		}

		return nil
	})

	return err
}
