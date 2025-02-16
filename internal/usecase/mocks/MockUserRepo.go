package mocks

import (
	"avito-shop-service/internal/database/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(ctx context.Context, username string, password string) (int, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockUserRepo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepo) GetUserCoinsByUserID(ctx context.Context, userID int) (int, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockUserRepo) AddCoinsToUserByUserID(ctx context.Context, userID int, amount int) error {
	args := m.Called(ctx, userID, amount)
	return args.Error(0)
}

func (m *MockUserRepo) SubtractCoinsFromUserByUserID(ctx context.Context, userID int, amount int) error {
	args := m.Called(ctx, userID, amount)
	return args.Error(0)
}
