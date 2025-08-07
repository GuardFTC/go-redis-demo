// Package hash 提供Redis哈希操作的封装
// @Author:冯铁城 [17615007230@163.com] 2025-08-07 10:00:00
package hash

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Client Redis哈希操作客户端
type Client struct {
	rdb *redis.Client
}

// New 创建哈希操作客户端
func New(rdb *redis.Client) *Client {
	return &Client{rdb: rdb}
}

// HSet 写入键值对（存在则覆盖）
func (c *Client) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return c.rdb.HSet(ctx, key, values...).Result()
}

// HSetNX 写入键值对（存在不覆盖）
func (c *Client) HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	return c.rdb.HSetNX(ctx, key, field, value).Result()
}

// HGet 获取键值对
func (c *Client) HGet(ctx context.Context, key, field string) (string, error) {
	return c.rdb.HGet(ctx, key, field).Result()
}

// HMGet 获取多个键值对
func (c *Client) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return c.rdb.HMGet(ctx, key, fields...).Result()
}

// HGetAll 获取所有键值对
func (c *Client) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return c.rdb.HGetAll(ctx, key).Result()
}

// HKeys 获取所有键
func (c *Client) HKeys(ctx context.Context, key string) ([]string, error) {
	return c.rdb.HKeys(ctx, key).Result()
}

// HVals 获取所有值
func (c *Client) HVals(ctx context.Context, key string) ([]string, error) {
	return c.rdb.HVals(ctx, key).Result()
}

// HDel 删除键值对
func (c *Client) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return c.rdb.HDel(ctx, key, fields...).Result()
}

// HExists 判断字段是否存在
func (c *Client) HExists(ctx context.Context, key, field string) (bool, error) {
	return c.rdb.HExists(ctx, key, field).Result()
}

// HLen 获取键值对数量
func (c *Client) HLen(ctx context.Context, key string) (int64, error) {
	return c.rdb.HLen(ctx, key).Result()
}

// HStrLen 获取值的长度
func (c *Client) HStrLen(ctx context.Context, key, field string) (int64, error) {
	return c.rdb.HStrLen(ctx, key, field).Result()
}

// HIncrBy 给字段的值加上一个整数（负数即为减法）
func (c *Client) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	return c.rdb.HIncrBy(ctx, key, field, incr).Result()
}

// HIncrByFloat 给字段的值加上一个数（可以是浮点数）
func (c *Client) HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error) {
	return c.rdb.HIncrByFloat(ctx, key, field, incr).Result()
}
