package redis

import (
	"fmt"
	"sooty-tern/pkg/store/redis"
	"time"
)

// NewStore 创建基于redis存储实例
func NewStore(client *redis.Client) *Store {
	return &Store{
		client,
	}
}

type Store struct {
	*redis.Client
}

func (a *Store) wrapperKey(key string) string {
	return fmt.Sprintf("%s%s", a.KeyPrefix, key)
}

// Set ...
func (a *Store) Set(tokenString string, expiration time.Duration) error {
	cmd := a.Cli.Set(a.wrapperKey(tokenString), "1", expiration)
	return cmd.Err()
}

// Check ...
func (a *Store) Check(tokenString string) (bool, error) {
	cmd := a.Cli.Exists(a.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Close ...
func (a *Store) Close() error {
	return a.Cli.Close()
}
