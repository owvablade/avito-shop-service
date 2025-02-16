package jwtservice

import (
	"avito-shop-service/internal/config"
	"avito-shop-service/internal/errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

type Interface interface {
	GenerateSignedTokenFromUserId(userId int) (string, error)
}

type Service struct {
	JwtSecret []byte
}

func New(cfg *config.Config) *Service {
	return &Service{JwtSecret: []byte(cfg.JwtSecret)}
}

func (s *Service) GenerateSignedTokenFromUserId(userId int) (string, error) {
	claims := &CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(s.JwtSecret)
	if err != nil {
		return "", errors.ErrInternalServer
	}

	return signedToken, nil
}
