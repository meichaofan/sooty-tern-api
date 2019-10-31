package internal

import (
	"context"
	"sooty-tern/internal/app/errors"
	"sooty-tern/internal/app/model"
	"sooty-tern/internal/app/schema"
	"sooty-tern/pkg/util"
)

// User
type User struct {
	UserModel model.IUser
}

// NewUser
func NewUser(mUser model.IUser) *User {
	return &User{
		UserModel: mUser,
	}
}

// Query
func (u *User) Query(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserQueryResult, error) {
	return u.UserModel.Query(ctx, params, opts...)
}

// Get
func (u *User) Get(ctx context.Context, recordId string) (*schema.User, error) {
	item, err := u.UserModel.Get(ctx, recordId)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	return item, nil
}

// Create
func (u *User) Create(ctx context.Context, item schema.User) (*schema.User, error) {
	item.RecordId = util.NewRecordId()
	err := u.UserModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	user, err := u.Get(ctx, item.RecordId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Delete
func (u *User) Delete(ctx context.Context, recordId string) error {
	oldItem, err := u.UserModel.Get(ctx, recordId)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	err = u.UserModel.Delete(ctx, recordId)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatus
func (u *User) UpdateStatus(ctx context.Context, recordId string, status int) error {
	oldItem, err := u.UserModel.Get(ctx, recordId)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	err = u.UserModel.UpdateStatus(ctx, recordId, status)
	if err != nil {
		return err
	}
	return nil
}
