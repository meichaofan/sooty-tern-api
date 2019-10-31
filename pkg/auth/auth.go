package auth

import "errors"

// 定义错误
var (
	ErrInvalidToken = errors.New("invalid token")
)

// TokenInfo 令牌信息
type TokenInfo interface {
	GetAccessToken() string        // 获取访问令牌
	GetTokenType() string          // 获取令牌类型
	GetExpiresAt() int64           // 获取令牌到期时间戳
	EncodeToJSON() ([]byte, error) // JSON编码
}

// Auth 认证接口
type Auth interface {
	GenerateToken(data string) (TokenInfo, error)   // 生成令牌
	DestroyToken(accessToken string) error          // 销毁令牌
	ParseData(accessToken string) (string, error)   // 解析uuid
	Release() error                                 // 释放资源
}
