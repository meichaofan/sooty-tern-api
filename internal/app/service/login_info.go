package service

import (
	"context"
	"sooty-tern/internal/app/schema"
)

type ILoginInfo interface {
	Login(ctx context.Context, code string) (*schema.LoginRes, error)
	GenerateToken(openId, sessionKey string) (*schema.LoginTokenInfo, error)
	DestroyToken(tokenString string) error
}
