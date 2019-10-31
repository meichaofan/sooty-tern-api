package app

import (
	"errors"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"sooty-tern/internal/app/config"
	igorm "sooty-tern/internal/app/model/impl/gorm"
)

// InitStore 初始化存储
func InitStore(container *dig.Container) (func(), error) {
	var storeCall func()
	cfg := config.GetGlobalConfig()

	switch cfg.Store {
	case "gorm":
		db, err := initGorm()
		if err != nil {
			return nil, err
		}

		storeCall = func() {
			db.Close()
		}

		igorm.SetTablePrefix(cfg.Gorm.TablePrefix)

		if cfg.Gorm.EnableAutoMigrate {
			err = igorm.AutoMigrate(db)
			if err != nil {
				return nil, err
			}
		}

		// 注入DB
		container.Provide(func() *gorm.DB {
			return db
		})
		igorm.Inject(container)

	default:
		return nil, errors.New("unknown store")
	}
	return storeCall, nil
}

func initGorm() (*gorm.DB, error) {
	cfg := config.GetGlobalConfig()
	var dsn string
	switch cfg.Gorm.DBType {
	case "mysql":
		dsn = cfg.MySQL.DSN()
	default:
		return nil, errors.New("unknown db")
	}

	return igorm.NewDB(&igorm.Config{
		Debug:       cfg.Gorm.Debug,
		DBType:      cfg.Gorm.DBType,
		DSN:         dsn,
		MaxIdleCons: cfg.Gorm.MaxIdleConns,
		MaxLifeTime: cfg.Gorm.MaxLifetime,
		MaxOpenCons: cfg.Gorm.MaxOpenConns,
	})
}
