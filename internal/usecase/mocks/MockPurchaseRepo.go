package mocks

import (
	"avito-shop-service/internal/database/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type MockPurchaseRepo struct {
	mock.Mock
}

func (m *MockPurchaseRepo) CreatePurchase(ctx context.Context, userID int, merchID int) error {
	args := m.Called(ctx, userID, merchID)
	return args.Error(0)
}

func (m *MockPurchaseRepo) GetPurchasesByUserID(ctx context.Context, userID int) ([]model.Inventory, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]model.Inventory), args.Error(1)
}
