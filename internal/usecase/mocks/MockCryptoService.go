package mocks

import "github.com/stretchr/testify/mock"

type MockCryptoService struct {
	mock.Mock
}

func (m *MockCryptoService) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockCryptoService) CompareHashAndPassword(hash string, password string) error {
	args := m.Called(hash, password)
	return args.Error(0)
}
