package mocks

import (
	"avito-shop-service/internal/database/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type MockMerchRepo struct {
	mock.Mock
}

func (m *MockMerchRepo) GetMerchItemByMerchName(ctx context.Context, merchName string) (*model.MerchItem, error) {
	args := m.Called(ctx, merchName)
	return args.Get(0).(*model.MerchItem), args.Error(1)
}
