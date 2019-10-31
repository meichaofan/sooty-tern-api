package app

import (
	"context"
	"fmt"
	"net/http"
	"sooty-tern/internal/app/config"
	"sooty-tern/internal/app/middleware"
	"sooty-tern/internal/app/routers"
	"sooty-tern/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// InitWeb 初始化web引擎
func InitWeb(container *dig.Container) *gin.Engine {
	cfg := config.GetGlobalConfig()
	gin.SetMode(cfg.RunMode)

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	apiPrefixes := []string{"/api/"}

	// 跟踪ID
	app.Use(middleware.TraceMiddleware(middleware.AllowPathPrefixNoSkipper(apiPrefixes...)))

	// 访问日志
	app.Use(middleware.LoggerMiddleware(middleware.AllowPathPrefixNoSkipper(apiPrefixes...)))

	// 注册/api路由
	err := routers.RegisterRouter(app, container)
	handleError(err)

	return app
}

// InitHTTPServer 初始化http服务
func InitHTTPServer(ctx context.Context, container *dig.Container) func() {
	cfg := config.GetGlobalConfig().HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      InitWeb(container),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.Printf(ctx, "HTTP server is running , listen port on：[%s]", addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Errorf(ctx, err.Error())
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf(ctx, err.Error())
		}
	}
}
