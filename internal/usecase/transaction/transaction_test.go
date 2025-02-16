package transaction

import (
	"avito-shop-service/internal/database/model"
	"avito-shop-service/internal/dto"
	"avito-shop-service/internal/usecase/mocks"
	"context"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	ctx := context.Background()

	mockTxManager := new(mocks.MockTxManager)
	mockUserRepo := new(mocks.MockUserRepo)
	mockTransactionRepo := new(mocks.MockTransactionRepo)

	uc := New(mockTxManager, mockUserRepo, mockTransactionRepo)

	inFromID := 1
	inSendCoinReq := &dto.SendCoinRequest{
		ToUser: "touser",
		Amount: 100,
	}
	expectedCoins := 900
	expectedUser := &model.User{
		ID:       2,
		Username: "touser",
		Password: "pass",
		Coins:    0,
	}

	mockTxManager.On("ReadWrite", ctx, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		fn := args.Get(1).(func(ctx context.Context) error)
		err := fn(ctx)
		if err != nil {
			return
		}
	})

	mockUserRepo.On("GetUserByUsername", ctx, inSendCoinReq.ToUser).Return(expectedUser, nil)
	mockUserRepo.On("GetUserCoinsByUserID", ctx, inFromID).Return(expectedCoins, nil)
	mockUserRepo.On("SubtractCoinsFromUserByUserID", ctx, inFromID, inSendCoinReq.Amount).Return(nil)
	mockUserRepo.On("AddCoinsToUserByUserID", ctx, expectedUser.ID, inSendCoinReq.Amount).Return(nil)
	mockTransactionRepo.On("CreateTransaction", ctx, inFromID, expectedUser.ID, inSendCoinReq.Amount).Return(nil)

	err := uc.CreateTransaction(ctx, inFromID, inSendCoinReq)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	mockUserRepo.AssertCalled(t, "GetUserByUsername", ctx, inSendCoinReq.ToUser)
	mockUserRepo.AssertCalled(t, "GetUserCoinsByUserID", ctx, inFromID)
	mockUserRepo.AssertCalled(t, "SubtractCoinsFromUserByUserID", ctx, inFromID, inSendCoinReq.Amount)
	mockUserRepo.AssertCalled(t, "AddCoinsToUserByUserID", ctx, expectedUser.ID, inSendCoinReq.Amount)
	mockTransactionRepo.AssertCalled(t, "CreateTransaction", ctx, inFromID, expectedUser.ID, inSendCoinReq.Amount)
}
