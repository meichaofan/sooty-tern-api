package entity

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"sooty-tern/internal/app/config"
	icontext "sooty-tern/internal/app/context"
	"sooty-tern/pkg/util"
	"time"
)

//表名前缀
var tablePrefix string

func SetTablePrefix(prefix string) {
	tablePrefix = prefix
}

func GetTablePrefix() string {
	return tablePrefix
}

// Model base model
type Model struct {
	ID        int        `gorm:"column:id;primary_key;auto_increment;"`
	CreatedAt time.Time  `gorm:"column:created_at;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;"`
}

// table name
func (Model) TableName(name string) string {
	return fmt.Sprintf("%s%s", GetTablePrefix(), name)
}

func toString(v interface{}) string {
	return util.JSONMarshalToString(v)
}

func getDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := icontext.FromTrans(ctx)
	if ok {
		db, ok := trans.(*gorm.DB)
		if ok {
			if icontext.FromTransLock(ctx) {
				if dbType := config.GetGlobalConfig().Gorm.DBType; dbType == "mysql" ||
					dbType == "postgres" {
					db = db.Set("gorm:query_option", "FOR UPDATE")
				}
			}
			return db
		}
	}
	return defDB
}

func getDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return getDB(ctx, defDB).Model(m)
}
