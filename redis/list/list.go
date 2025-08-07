// Package list 提供Redis列表操作的封装
// @Author:冯铁城 [17615007230@163.com] 2025-08-07 10:00:00
package list

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// Client Redis列表操作客户端
type Client struct {
	rdb *redis.Client
}

// New 创建列表操作客户端
func New(rdb *redis.Client) *Client {
	return &Client{rdb: rdb}
}

// LPush 左端推入元素（Key不存在创建Key）
func (c *Client) LPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return c.rdb.LPush(ctx, key, values...).Result()
}

// LPushX 左端推入元素（Key不存在不做操作）
func (c *Client) LPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return c.rdb.LPushX(ctx, key, values...).Result()
}

// RPush 右端推入元素（Key不存在创建Key）
func (c *Client) RPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return c.rdb.RPush(ctx, key, values...).Result()
}

// RPushX 右端推入元素（Key不存在不做操作）
func (c *Client) RPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return c.rdb.RPushX(ctx, key, values...).Result()
}

// LPop 左端弹出
func (c *Client) LPop(ctx context.Context, key string) (string, error) {
	return c.rdb.LPop(ctx, key).Result()
}

// RPop 右端弹出
func (c *Client) RPop(ctx context.Context, key string) (string, error) {
	return c.rdb.RPop(ctx, key).Result()
}

// LIndex 返回索引处的元素
func (c *Client) LIndex(ctx context.Context, key string, index int64) (string, error) {
	return c.rdb.LIndex(ctx, key, index).Result()
}

// LInsert 在目标元素前或后插入元素
func (c *Client) LInsert(ctx context.Context, key, op string, pivot, value interface{}) (int64, error) {
	return c.rdb.LInsert(ctx, key, op, pivot, value).Result()
}

// LRange 获取指定范围的元素
func (c *Client) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return c.rdb.LRange(ctx, key, start, stop).Result()
}

// LLen 获取集合长度
func (c *Client) LLen(ctx context.Context, key string) (int64, error) {
	return c.rdb.LLen(ctx, key).Result()
}

// LRem 删除n个指定元素
func (c *Client) LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error) {
	return c.rdb.LRem(ctx, key, count, value).Result()
}

// LSet 更新指定下标的值
func (c *Client) LSet(ctx context.Context, key string, index int64, value interface{}) error {
	return c.rdb.LSet(ctx, key, index, value).Err()
}

// LTrim 裁剪list
func (c *Client) LTrim(ctx context.Context, key string, start, stop int64) error {
	return c.rdb.LTrim(ctx, key, start, stop).Err()
}

// RPopLPush 右边弹出，左边推入
func (c *Client) RPopLPush(ctx context.Context, source, destination string) (string, error) {
	return c.rdb.RPopLPush(ctx, source, destination).Result()
}

// BRPopLPush 阻塞式右边弹出，左边推入
func (c *Client) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (string, error) {
	return c.rdb.BRPopLPush(ctx, source, destination, timeout).Result()
}

// BLPop 阻塞式左端弹出
func (c *Client) BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return c.rdb.BLPop(ctx, timeout, keys...).Result()
}

// BRPop 阻塞式右端弹出
func (c *Client) BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return c.rdb.BRPop(ctx, timeout, keys...).Result()
}
