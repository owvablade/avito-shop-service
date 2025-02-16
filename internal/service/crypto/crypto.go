package crypto

import "golang.org/x/crypto/bcrypt"

type Interface interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hash string, password string) error
}

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s Service) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s Service) CompareHashAndPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
