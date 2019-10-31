package app

import (
	"context"
	"go.uber.org/dig"
	"os"
	"sooty-tern/internal/app/config"
	"sooty-tern/internal/app/service/impl"
	"sooty-tern/pkg/auth"
	"sooty-tern/pkg/logger"
)

type options struct {
	ConfigFile string
	Version    string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Init 应用初始化
func Init(ctx context.Context, opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	err := config.LoadGlobalConfig(o.ConfigFile)
	handleError(err)

	cfg := config.GetGlobalConfig()

	logger.Printf(ctx, "service started , run mode：%s , version：%s , process id：%d\n", cfg.RunMode, o.Version, os.Getpid())

	loggerCall, err := InitLogger()
	handleError(err)

	// 创建依赖注入容器
	container, containerCall := BuildContainer()

	// init http server
	httpCall := InitHTTPServer(ctx, container)
	return func() {
		if httpCall != nil {
			httpCall()
		}
		if containerCall != nil {
			containerCall()
		}
		if loggerCall != nil {
			loggerCall()
		}
	}
}

func BuildContainer() (*dig.Container, func()) {
	// 创建依赖注入容器
	container := dig.New()
	// inject service
	err := impl.Inject(container)
	handleError(err)
	// inject storage
	storeCall, err := InitStore(container)
	handleError(err)
	// inject auth
	authentication, err := InitAuth()
	handleError(err)
	_ = container.Provide(func() auth.Auth {
		return authentication
	})

	return container, func() {
		if storeCall != nil {
			storeCall()
		}
	}
}
