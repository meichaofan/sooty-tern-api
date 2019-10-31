package redis

import (
	"github.com/go-redis/redis"
)

// Config redis配置参数
type Config struct {
	Addr      string
	DB        int
	Password  string
	KeyPrefix string
}

type Client struct {
	Cli       *redis.Client
	KeyPrefix string
}

func NewRedisClient(c *Config) *Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	return &Client{cli, c.KeyPrefix}
}
