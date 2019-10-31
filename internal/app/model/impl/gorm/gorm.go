package gorm

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"sooty-tern/internal/app/config"
	"sooty-tern/internal/app/model"
	"sooty-tern/internal/app/model/impl/gorm/internal/entity"
	imodel "sooty-tern/internal/app/model/impl/gorm/internal/model"
	"time"

	// gorm存储注入
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Debug       bool
	DBType      string
	DSN         string
	MaxLifeTime int
	MaxOpenCons int
	MaxIdleCons int
}

// 创建DB实例
func NewDB(c *Config) (*gorm.DB, error) {
	db, err := gorm.Open(c.DBType, c.DSN)
	if err != nil {
		return nil, err
	}
	if c.Debug {
		db = db.Debug()
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	db.DB().SetConnMaxLifetime(time.Duration(c.MaxLifeTime) * time.Second)
	db.DB().SetMaxOpenConns(c.MaxOpenCons)
	db.DB().SetMaxIdleConns(c.MaxIdleCons)
	return db, nil
}

func SetTablePrefix(prefix string) {
	entity.SetTablePrefix(prefix)
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.GetGlobalConfig().Gorm.DBType; dbType == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	return db.AutoMigrate(new(entity.User)).Error
}

func Inject(container *dig.Container) error {
	container.Provide(imodel.NewTrans, dig.As(new(model.ITrans)))
	container.Provide(imodel.NewUser, dig.As(new(model.IUser)))
	container.Provide(imodel.NewLoginInfo, dig.As(new(model.ILoginInfoModel)))
	return nil
}
