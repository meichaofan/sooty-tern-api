package model

import (
	"context"

	"sooty-tern/internal/app/schema"
)

// IUser
type IUser interface {
	Query(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserQueryResult, error)
	Get(ctx context.Context, RecordId string) (*schema.User, error)
	Create(ctx context.Context, item schema.User) error
	Delete(ctx context.Context, RecordId string) error
	UpdateStatus(ctx context.Context, RecordId string, status int) error
}
