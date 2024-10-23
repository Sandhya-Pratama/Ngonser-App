package service

import (
	"context"
	"time"

	"github.com/Sandhya-Pratama/Ngonser-App/common"
	"github.com/Sandhya-Pratama/Ngonser-App/entity"
	"github.com/Sandhya-Pratama/Ngonser-App/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	GenerateAccessToken(ctx context.Context, user *entity.User) (string, error)
}

type TokenService struct {
	cfg *config.Config //secret key nya diambil dari config
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{cfg: cfg}
}

func (s *TokenService) GenerateAccessToken(ctx context.Context, user *entity.User) (string, error) {
	expiredTime := time.Now().Local().Add(10 * time.Minute).Unix()
	claim := common.JwtCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		Roles: user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expiredTime, 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(s.cfg.JWT.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
