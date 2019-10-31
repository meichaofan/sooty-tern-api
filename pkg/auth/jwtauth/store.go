package jwtauth

import (
	"time"
)

// Store
type Store interface {
	Set(tokenString string, expiration time.Duration) error // 放入令牌，指定到期时间
	Check(tokenString string) (bool, error)                 // 检查令牌是否存在
	Close() error                                           // 关闭存储
}
