package model

import (
	"context"
	"github.com/jinzhu/gorm"
	"sooty-tern/internal/app/errors"
	"sooty-tern/internal/app/model/impl/gorm/internal/entity"
	"sooty-tern/internal/app/schema"
)

type LoginInfo struct {
	db *gorm.DB
}

func NewLoginInfo(db *gorm.DB) *LoginInfo {
	return &LoginInfo{db}
}

// Get
func (l *LoginInfo) Get(ctx context.Context, params schema.LoginInfoQueryParam) (*schema.LoginInfo, error) {
	var item entity.LoginInfo
	db := entity.GetLoginInfoDB(ctx, l.db)
	if uid := params.Uid; uid != "" {
		db = db.Where("uid=?", uid)
	}
	if uidKey := params.UidKey; uidKey != "" {
		db = db.Where("uid_key=?", uidKey)
	}
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	schemaUser := item.ToSchemaLoginInfo()
	return schemaUser, nil
}

// Create
func (l *LoginInfo) Create(ctx context.Context, item schema.LoginInfo) error {
	return ExecTrans(ctx, l.db, func(ctx context.Context) error {
		schemaLoginInfo := entity.SchemaLoginInfo(item)
		result := entity.GetLoginInfoDB(ctx, l.db).Create(schemaLoginInfo.ToLoginInfoEntity())
		if err := result.Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

// Update
func (l *LoginInfo) Update(ctx context.Context, params schema.LoginInfoQueryParam, item schema.LoginInfo) error {
	return ExecTrans(ctx, l.db, func(ctx context.Context) error {
		schemaLoginInfo := entity.SchemaLoginInfo(item)
		db := entity.GetLoginInfoDB(ctx, l.db)
		if params.Uid != "" {
			db.Where("uid = ?", params.Uid)
		}
		if params.UidKey != "" {
			db.Where("uid_key = ?", params.UidKey)
		}
		result := db.Updates(schemaLoginInfo.ToLoginInfoEntity())
		if err := result.Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}
