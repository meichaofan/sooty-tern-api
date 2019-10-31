package model

import (
	"context"

	"sooty-tern/internal/app/schema"
)

// IUser
type ILoginInfoModel interface {
	Get(ctx context.Context, params schema.LoginInfoQueryParam) (*schema.LoginInfo, error)
	Create(ctx context.Context, item schema.LoginInfo) error
	Update(ctx context.Context, params schema.LoginInfoQueryParam, item schema.LoginInfo) error
}
