// Package string 提供Redis字符串操作的封装
// @Author:冯铁城 [17615007230@163.com] 2025-08-07 10:00:00
package string

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// Client Redis字符串操作客户端
type Client struct {
	rdb *redis.Client
}

// New 创建字符串操作客户端
func New(rdb *redis.Client) *Client {
	return &Client{rdb: rdb}
}

// SetWithDefaultExpire 设置key，使用默认过期时间（存在则覆盖）
func (c *Client) SetWithDefaultExpire(ctx context.Context, key string, value interface{}) error {
	return c.Set(ctx, key, value, 15*time.Minute) // 默认15分钟过期
}

// Set 设置key（存在则覆盖）
func (c *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.rdb.Set(ctx, key, value, expiration).Err()
}

// SetNXWithDefaultExpire 设置key，使用默认过期时间（存在不覆盖）
func (c *Client) SetNXWithDefaultExpire(ctx context.Context, key string, value interface{}) (bool, error) {
	return c.SetNX(ctx, key, value, 15*time.Minute)
}

// SetNX 设置key（存在不覆盖）
func (c *Client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.rdb.SetNX(ctx, key, value, expiration).Result()
}

// Get 获取key
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}

// MSet 设置多个key-value（存在则覆盖）
func (c *Client) MSet(ctx context.Context, pairs ...interface{}) error {
	return c.rdb.MSet(ctx, pairs...).Err()
}

// MSetNX 设置多个key-value（存在不覆盖）
func (c *Client) MSetNX(ctx context.Context, pairs ...interface{}) (bool, error) {
	return c.rdb.MSetNX(ctx, pairs...).Result()
}

// Keys 获取所有匹配的key（慎用）
func (c *Client) Keys(ctx context.Context, pattern string) ([]string, error) {
	return c.rdb.Keys(ctx, pattern).Result()
}

// Exists 判断key是否存在（返回匹配个数）
func (c *Client) Exists(ctx context.Context, keys ...string) (int64, error) {
	return c.rdb.Exists(ctx, keys...).Result()
}

// TTL 获取key剩余TTL（秒级）
func (c *Client) TTL(ctx context.Context, key string) (time.Duration, error) {
	return c.rdb.TTL(ctx, key).Result()
}

// Expire 设置过期时间（秒）
func (c *Client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.rdb.Expire(ctx, key, expiration).Err()
}

// PExpire 设置过期时间（毫秒）
func (c *Client) PExpire(ctx context.Context, key string, expiration time.Duration) error {
	return c.rdb.PExpire(ctx, key, expiration).Err()
}

// Incr 增加key的值（整数）
func (c *Client) Incr(ctx context.Context, key string) (int64, error) {
	return c.rdb.Incr(ctx, key).Result()
}

// IncrBy 按指定值增加
func (c *Client) IncrBy(ctx context.Context, key string, increment int64) (int64, error) {
	return c.rdb.IncrBy(ctx, key, increment).Result()
}

// Decr 减少key的值
func (c *Client) Decr(ctx context.Context, key string) (int64, error) {
	return c.rdb.Decr(ctx, key).Result()
}

// DecrBy 按指定值减少
func (c *Client) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	return c.rdb.DecrBy(ctx, key, decrement).Result()
}

// Del 删除key
func (c *Client) Del(ctx context.Context, keys ...string) (int64, error) {
	return c.rdb.Del(ctx, keys...).Result()
}
