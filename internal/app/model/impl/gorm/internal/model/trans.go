package model

import (
	"context"
	"sooty-tern/internal/app/errors"

	"github.com/jinzhu/gorm"
)

// NewTrans 创建事务管理实例
func NewTrans(db *gorm.DB) *Trans {
	return &Trans{db}
}

// Trans 事务管理
type Trans struct {
	db *gorm.DB
}

// Begin
func (a *Trans) Begin(ctx context.Context) (interface{}, error) {
	result := a.db.Begin()
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}

// Commit
func (a *Trans) Commit(ctx context.Context, trans interface{}) error {
	db, ok := trans.(*gorm.DB)
	if !ok {
		return errors.New("unknow trans")
	}

	result := db.Commit()
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Rollback
func (a *Trans) Rollback(ctx context.Context, trans interface{}) error {
	db, ok := trans.(*gorm.DB)
	if !ok {
		return errors.New("unknow trans")
	}

	result := db.Rollback()
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
