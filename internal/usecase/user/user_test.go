package user

import (
	"avito-shop-service/internal/database/model"
	"avito-shop-service/internal/dto"
	"avito-shop-service/internal/errors"
	"avito-shop-service/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"testing"
)

func TestAuthUser(t *testing.T) {
	ctx := context.Background()

	mockJwtService := new(mocks.MockJwtService)
	mockCryptoService := new(mocks.MockCryptoService)
	mockUserRepo := new(mocks.MockUserRepo)

	uc := New(mockJwtService, mockCryptoService, nil, mockUserRepo, nil, nil)

	in := &dto.AuthRequest{Username: "artem", Password: "cherepanov"}
	user := &model.User{ID: 1, Username: in.Username, Password: in.Password, Coins: 1000}
	expResponse := &dto.AuthResponse{Token: "signedToken"}

	mockUserRepo.On("GetUserByUsername", ctx, in.Username).Return(user, nil)
	mockJwtService.On("GenerateSignedTokenFromUserId", user.ID).Return(expResponse.Token, nil)
	mockCryptoService.On("CompareHashAndPassword", mock.Anything, mock.Anything).Return(nil)

	actualResponse, err := uc.CreateOrAuthUser(ctx, in)

	mockUserRepo.AssertExpectations(t)
	mockJwtService.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if actualResponse.Token != expResponse.Token {
		t.Errorf("Expected token %s, but got %s", expResponse.Token, actualResponse.Token)
	}
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	mockJwtService := new(mocks.MockJwtService)
	mockCryptoService := new(mocks.MockCryptoService)
	mockUserRepo := new(mocks.MockUserRepo)

	uc := New(mockJwtService, mockCryptoService, nil, mockUserRepo, nil, nil)

	in := &dto.AuthRequest{Username: "artem", Password: "cherepanov"}
	user := &model.User{ID: 1, Username: in.Username, Password: in.Password, Coins: 1000}
	expResponse := &dto.AuthResponse{Token: "signedToken"}

	mockUserRepo.On("GetUserByUsername", ctx, in.Username).Return(&model.User{}, errors.ErrUserNotFound)
	mockCryptoService.On("HashPassword", in.Password).Return(mock.Anything, nil)
	mockUserRepo.On("CreateUser", ctx, in.Username, mock.Anything).Return(user.ID, nil)
	mockJwtService.On("GenerateSignedTokenFromUserId", user.ID).Return(expResponse.Token, nil)

	actualResponse, err := uc.CreateOrAuthUser(ctx, in)

	mockUserRepo.AssertExpectations(t)
	mockJwtService.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if actualResponse.Token != expResponse.Token {
		t.Errorf("Expected token %s, but got %s", expResponse.Token, actualResponse.Token)
	}
}

func TestGetUserInfo(t *testing.T) {
	ctx := context.Background()

	mockTxManager := new(mocks.MockTxManager)
	mockUserRepo := new(mocks.MockUserRepo)
	mockPurchaseRepo := new(mocks.MockPurchaseRepo)
	mockTransactionRepo := new(mocks.MockTransactionRepo)

	uc := New(nil, nil, mockTxManager, mockUserRepo, mockPurchaseRepo, mockTransactionRepo)

	inUserID := 1
	expectedCoins := 850
	expectedPurchases := []model.Inventory{{
		Type:     "t-shirt",
		Quantity: 100,
	}}
	expectedSentTransactions := []model.SentTransaction{{
		ToUser: "touser",
		Amount: 100,
	}}
	expectedReceivedTransactions := []model.ReceivedTransaction{{
		FromUser: "fromuser",
		Amount:   50,
	}}

	mockTxManager.On("ReadOnly", ctx, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		fn := args.Get(1).(func(ctx context.Context) error)
		err := fn(ctx)
		if err != nil {
			return
		}
	})

	mockUserRepo.On("GetUserCoinsByUserID", ctx, inUserID).Return(expectedCoins, nil)
	mockPurchaseRepo.On("GetPurchasesByUserID", ctx, inUserID).Return(expectedPurchases, nil)
	mockTransactionRepo.On("GetSentTransactionsFromUserID", ctx, inUserID).Return(expectedSentTransactions, nil)
	mockTransactionRepo.On("GetReceivedTransactionsToUserID", ctx, inUserID).Return(expectedReceivedTransactions, nil)

	result, err := uc.GetUserInfo(ctx, inUserID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCoins, result.Coins)
	assert.Equal(t, expectedPurchases, result.Inventory)
	assert.Equal(t, expectedSentTransactions, result.CoinHistory.Sent)
	assert.Equal(t, expectedReceivedTransactions, result.CoinHistory.Received)

	mockTxManager.AssertCalled(t, "ReadOnly", ctx, mock.Anything)
	mockUserRepo.AssertCalled(t, "GetUserCoinsByUserID", ctx, inUserID)
	mockPurchaseRepo.AssertCalled(t, "GetPurchasesByUserID", ctx, inUserID)
	mockTransactionRepo.AssertCalled(t, "GetSentTransactionsFromUserID", ctx, inUserID)
	mockTransactionRepo.AssertCalled(t, "GetReceivedTransactionsToUserID", ctx, inUserID)
}
