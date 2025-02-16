package mocks

import "github.com/stretchr/testify/mock"

type MockJwtService struct {
	mock.Mock
}

func (m *MockJwtService) GenerateSignedTokenFromUserId(userId int) (string, error) {
	args := m.Called(userId)
	return args.Get(0).(string), args.Error(1)
}
