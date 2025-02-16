package purchase

import (
	"avito-shop-service/internal/database/model"
	"avito-shop-service/internal/usecase/mocks"
	"context"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreatePurchase(t *testing.T) {
	ctx := context.Background()

	mockTxManager := new(mocks.MockTxManager)
	mockUserRepo := new(mocks.MockUserRepo)
	mockMerchRepo := new(mocks.MockMerchRepo)
	mockPurchaseRepo := new(mocks.MockPurchaseRepo)

	uc := New(mockTxManager, mockUserRepo, mockMerchRepo, mockPurchaseRepo)

	inFromID := 1
	inMerchName := "powerbank"
	expectedPrice := 80
	expectedCoins := 900
	expectedMerch := &model.MerchItem{
		ID:    1,
		Name:  "t-shirt",
		Price: 80,
	}

	mockTxManager.On("ReadWrite", ctx, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		fn := args.Get(1).(func(ctx context.Context) error)
		err := fn(ctx)
		if err != nil {
			return
		}
	})

	mockUserRepo.On("GetUserCoinsByUserID", ctx, inFromID).Return(expectedCoins, nil)
	mockMerchRepo.On("GetMerchItemByMerchName", ctx, inMerchName).Return(expectedMerch, nil)
	mockUserRepo.On("SubtractCoinsFromUserByUserID", ctx, inFromID, expectedPrice).Return(nil)
	mockPurchaseRepo.On("CreatePurchase", ctx, inFromID, expectedMerch.ID).Return(nil)

	err := uc.CreatePurchase(ctx, inFromID, mock.Anything)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	mockUserRepo.AssertCalled(t, "GetUserCoinsByUserID", ctx, inFromID)
	mockMerchRepo.AssertCalled(t, "GetMerchItemByMerchName", ctx, inMerchName)
	mockUserRepo.AssertCalled(t, "SubtractCoinsFromUserByUserID", ctx, inFromID, expectedPrice)
	mockPurchaseRepo.AssertCalled(t, "CreatePurchase", ctx, inFromID, expectedMerch.ID)
}
