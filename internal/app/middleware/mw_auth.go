package middleware

import (
	"github.com/gin-gonic/gin"
	"sooty-tern/internal/app/config"
	"sooty-tern/internal/app/errors"
	"sooty-tern/internal/app/ginplus"
	"sooty-tern/pkg/auth"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auth, skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		if t := ginplus.GetToken(c); t != "" {
			id, err := a.ParseData(t)
			if err != nil {
				if err == auth.ErrInvalidToken {
					ginplus.ResError(c, errors.ErrNoPerm)
					return
				}
				ginplus.ResError(c, errors.WithStack(err))
				return
			}
			userID = id
		}

		if userID != "" {
			c.Set(ginplus.UserIDKey, userID)
		}

		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		if userID == "" {
			if config.GetGlobalConfig().RunMode == "debug" {
				c.Set(ginplus.UserIDKey, config.GetGlobalConfig().Root.UserName)
				c.Next()
				return
			}
			ginplus.ResError(c, errors.ErrNoPerm)
		}
	}
}
