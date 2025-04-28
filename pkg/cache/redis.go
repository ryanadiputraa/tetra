package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/ryanadiputraa/tetra/config"
)

func NewRedis(c config.Config) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     c.RedisAddr,
		Password: c.RedisPassword,
		DB:       0,
	})

	_, err = rdb.Ping(context.Background()).Result()
	return
}
