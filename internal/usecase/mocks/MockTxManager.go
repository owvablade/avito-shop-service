package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockTxManager struct {
	mock.Mock
}

func (m *MockTxManager) ReadOnly(ctx context.Context, f func(context.Context) error) error {
	_ = m.Called(ctx, f)
	return f(ctx)
}

func (m *MockTxManager) ReadWrite(ctx context.Context, f func(context.Context) error) error {
	_ = m.Called(ctx, f)
	return f(ctx)
}
