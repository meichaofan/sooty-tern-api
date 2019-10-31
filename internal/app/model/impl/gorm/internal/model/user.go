package model

import (
	"context"
	"github.com/jinzhu/gorm"
	"sooty-tern/internal/app/errors"
	"sooty-tern/internal/app/model/impl/gorm/internal/entity"
	"sooty-tern/internal/app/schema"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db}
}

func (u *User) getQueryOption(opts ...schema.UserQueryOptions) schema.UserQueryOptions {
	var opt schema.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query
func (u *User) Query(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserQueryResult, error) {
	db := entity.GetUserDB(ctx, u.db)
	if v := params.UserName; v != "" {
		db = db.Where("user_name = ?", v)
	}
	if v := params.LikeUserName; v != "" {
		db = db.Where("user_name LIKE ?", "%"+v+"%")
	}
	if v := params.RecordId; v != "" {
		db = db.Where("record_id = ?", "%"+v+"%")
	}
	if v := params.LikeRecordId; v != "" {
		db = db.Where("record_id LIKE ?", "%"+v+"%")
	}
	db = db.Order("id ASC")
	opt := u.getQueryOption(opts...)
	var list entity.Users
	pr, err := WrapPageQuery(ctx, db, opt.PageParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.UserQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaUsers(),
	}
	return qr, nil
}

// Get
func (u *User) Get(ctx context.Context, recordId string) (*schema.User, error) {
	var item entity.User
	ok, err := FindOne(ctx, entity.GetUserDB(ctx, u.db).Where("record_id=?", recordId), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	schemaUser := item.ToSchemaUser()
	return schemaUser, nil
}

// Create
func (u *User) Create(ctx context.Context, item schema.User) error {
	return ExecTrans(ctx, u.db, func(ctx context.Context) error {
		schemaUser := entity.SchemaUser(item)
		result := entity.GetUserDB(ctx, u.db).Create(schemaUser.ToUserEntity())
		if err := result.Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

// Delete
func (u *User) Delete(ctx context.Context, RecordId string) error {
	return ExecTrans(ctx, u.db, func(ctx context.Context) error {
		result := entity.GetUserDB(ctx, u.db).Where("record_id=?", RecordId).Update("is_del", 1)
		if err := result.Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

// UpdateStatus
func (u *User) UpdateStatus(ctx context.Context, RecordId string, status int) error {
	result := entity.GetUserDB(ctx, u.db).Where("record_id=?", RecordId).Update("status", status)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
