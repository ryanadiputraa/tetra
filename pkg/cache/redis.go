package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/ryanadiputraa/inventra/config"
)

func NewRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.RedisAddr,
		Password: c.RedisPassword,
		DB:       0,
	})
}
