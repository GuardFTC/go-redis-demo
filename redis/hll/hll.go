// Package hll 提供Redis HyperLogLog操作的封装
// @Author:冯铁城 [17615007230@163.com] 2025-08-07 11:00:00
package hll

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Client Redis HyperLogLog操作客户端
type Client struct {
	rdb *redis.Client
}

// New 创建HyperLogLog操作客户端
func New(rdb *redis.Client) *Client {
	return &Client{rdb: rdb}
}

// PFAdd 添加指定元素到HyperLogLog中
func (c *Client) PFAdd(ctx context.Context, key string, els ...interface{}) (int64, error) {
	return c.rdb.PFAdd(ctx, key, els...).Result()
}

// PFCount 返回给定HyperLogLog的基数估算值
func (c *Client) PFCount(ctx context.Context, keys ...string) (int64, error) {
	return c.rdb.PFCount(ctx, keys...).Result()
}

// PFMerge 将多个HyperLogLog合并为一个HyperLogLog
func (c *Client) PFMerge(ctx context.Context, dest string, keys ...string) error {
	return c.rdb.PFMerge(ctx, dest, keys...).Err()
}
