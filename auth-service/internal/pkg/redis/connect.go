package rdb

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	DB *redis.Client
}

func NewRedisClient(c context.Context) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
	if _, err := rdb.Ping(c).Result(); err != nil {
		return nil, err
	}

	return &RedisClient{
		DB: rdb,
	}, nil
}

func (r *RedisClient) Close() error {
	return r.DB.Close()
}

func (r *RedisClient) Set(c context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.DB.Set(c, key, value, expiration)
}

func (r *RedisClient) Get(c context.Context, key string) *redis.StringCmd {
	return r.DB.Get(c, key)
}

func (r *RedisClient) Del(c context.Context, keys ...string) *redis.IntCmd {
	return r.DB.Del(c, keys...)
}
