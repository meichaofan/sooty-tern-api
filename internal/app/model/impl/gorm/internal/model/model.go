package model

import (
	"context"
	icontext "sooty-tern/internal/app/context"
	"sooty-tern/internal/app/schema"

	"github.com/jinzhu/gorm"
)

// TransFunc 定义事务执行函数
type TransFunc func(context.Context) error

// ExecTrans 执行事务
func ExecTrans(ctx context.Context, db *gorm.DB, fn TransFunc) error {
	if _, ok := icontext.FromTrans(ctx); ok {
		return fn(ctx)
	}

	transModel := NewTrans(db)
	trans, err := transModel.Begin(ctx)
	if err != nil {
		return err
	}

	ctx = icontext.NewTrans(ctx, trans)
	err = fn(ctx)
	if err != nil {
		_ = transModel.Rollback(ctx, trans)
		return err
	}
	return transModel.Commit(ctx, trans)
}

// ExecTransWithLock 执行事务（加锁）
func ExecTransWithLock(ctx context.Context, db *gorm.DB, fn TransFunc) error {
	if !icontext.FromTransLock(ctx) {
		ctx = icontext.NewTransLock(ctx)
	}
	return ExecTrans(ctx, db, fn)
}

// WrapPageQuery
func WrapPageQuery(ctx context.Context, db *gorm.DB, pi *schema.PageParam, out interface{}) (*schema.PageInfo, error) {
	if pi != nil {
		total, err := FindPage(ctx, db, pi.PageNumber, pi.PageSize, out)
		if err != nil {
			return nil, err
		}
		return &schema.PageInfo{
			PageNumber: pi.PageNumber,
			PageSize:   pi.PageSize,
			Total:      total,
		}, nil
	}
	result := db.Find(out)
	return nil, result.Error
}

// FindPage
func FindPage(ctx context.Context, db *gorm.DB, PageNumber, PageSize int, out interface{}) (int, error) {
	var count int
	result := db.Count(&count)
	if err := result.Error; err != nil {
		return 0, err
	} else if count == 0 {
		return 0, nil
	}

	if PageNumber < 0 || PageSize < 0 {
		return count, nil
	}

	if PageNumber > 0 && PageSize > 0 {
		db = db.Offset((PageNumber - 1) * PageSize)
	}
	if PageSize > 0 {
		db = db.Limit(PageSize)
	}
	result = db.Find(out)
	if err := result.Error; err != nil {
		return 0, err
	}
	return count, nil
}

// FindOne 查询单条数据
func FindOne(ctx context.Context, db *gorm.DB, out interface{}) (bool, error) {
	result := db.First(out)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
