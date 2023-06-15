package main

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(c *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{
		client: c,
		ttl:    ttl,
	}
}

func (c *RedisCache) Get(key int) (string, bool) {
	ctx := context.Background()
	intStr := strconv.Itoa(key)
	val, err := c.client.Get(ctx, intStr).Result()
	if err != nil {
		return "", false
	}

	return val, true
}

func (c *RedisCache) Set(key int, value string) error {
	ctx := context.Background()
	intStr := strconv.Itoa(key)
	if err := c.client.Set(ctx, intStr, value, c.ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (c *RedisCache) Remove(key int) error {
	ctx := context.Background()
	intStr := strconv.Itoa(key)
	if err := c.client.Del(ctx, intStr).Err(); err != nil {
		return err
	}

	return nil
}
