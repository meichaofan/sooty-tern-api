package service

import (
	"context"
	"sooty-tern/internal/app/schema"
)

type IUser interface {
	Query(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserQueryResult, error)
	Get(ctx context.Context, RecordId string) (*schema.User, error)
	Create(ctx context.Context, user schema.User) (*schema.User, error)
	UpdateStatus(ctx context.Context, RecordId string, status int) error
}
