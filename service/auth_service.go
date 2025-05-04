package service

import (
	"auth_service/web"
	"context"
)

type AuthService interface {
	GenerateJwt(ctx context.Context, request *web.Request) (*web.Token, error)
}
