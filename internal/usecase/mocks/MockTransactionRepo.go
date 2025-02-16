package mocks

import (
	"avito-shop-service/internal/database/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepo struct {
	mock.Mock
}

func (m *MockTransactionRepo) CreateTransaction(ctx context.Context, fromID int, toID int, amount int) error {
	args := m.Called(ctx, fromID, toID, amount)
	return args.Error(0)
}

func (m *MockTransactionRepo) GetSentTransactionsFromUserID(
	ctx context.Context,
	userID int) ([]model.SentTransaction, error) {

	args := m.Called(ctx, userID)
	return args.Get(0).([]model.SentTransaction), args.Error(1)
}

func (m *MockTransactionRepo) GetReceivedTransactionsToUserID(
	ctx context.Context,
	userID int) ([]model.ReceivedTransaction, error) {

	args := m.Called(ctx, userID)
	return args.Get(0).([]model.ReceivedTransaction), args.Error(1)
}
