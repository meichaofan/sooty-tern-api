package app

import (
	jwt "github.com/dgrijalva/jwt-go"
	"sooty-tern/internal/app/config"
	"sooty-tern/pkg/auth"
	"sooty-tern/pkg/auth/jwtauth"
	authRedis "sooty-tern/pkg/auth/jwtauth/store/redis"
	"sooty-tern/pkg/store/redis"
)

// InitAuth 初始化用户认证
func InitAuth() (auth.Auth, error) {
	cfg := config.GetGlobalConfig().JWTAuth

	var opts []jwtauth.Option
	opts = append(opts, jwtauth.SetExpired(cfg.Expired))
	opts = append(opts, jwtauth.SetSigningKey([]byte(cfg.SigningKey)))
	opts = append(opts, jwtauth.SetKeyFunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(cfg.SigningKey), nil
	}))

	switch cfg.SigningMethod {
	case "HS256":
		opts = append(opts, jwtauth.SetSigningMethod(jwt.SigningMethodHS256))
	case "HS384":
		opts = append(opts, jwtauth.SetSigningMethod(jwt.SigningMethodHS384))
	case "HS512":
		opts = append(opts, jwtauth.SetSigningMethod(jwt.SigningMethodHS512))
	}

	var store jwtauth.Store
	switch cfg.Store {
	case "redis":
		reidsCfg := config.GetGlobalConfig().Redis
		redisClient := redis.NewRedisClient(&redis.Config{
			Addr:      reidsCfg.Addr,
			Password:  reidsCfg.Password,
			DB:        cfg.RedisDB,
			KeyPrefix: cfg.RedisPrefix,
		})
		store = authRedis.NewStore(redisClient)
	}
	return jwtauth.New(store, opts...), nil
}
