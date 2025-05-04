package service

import (
	"auth_service/web"
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthServiceImpl struct{}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (svc *AuthServiceImpl) GenerateJwt(ctx context.Context, request *web.Request) (*web.Token, error) {
	aSecret := []byte(os.Getenv("JWT_ACCESS"))
	rSecret := []byte(os.Getenv("JWT_REFRESH"))
	aClaims := jwt.MapClaims{
		"username": request.Username,
		"exp":      time.Now().Add(5 * time.Minute).Unix(),
	}

	rClaims := jwt.MapClaims{
		"username": request.Username,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	}

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aClaims)
	signedAccessToken, err := aToken.SignedString(aSecret)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	rToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)
	signedRefreshToken, err := rToken.SignedString(rSecret)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	result := web.Token{
		Access:  signedAccessToken,
		Refresh: signedRefreshToken,
	}

	return &result, nil

}
