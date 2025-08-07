// Package bitmap 提供Redis位图操作的封装
// @Author:冯铁城 [17615007230@163.com] 2025-08-07 11:00:00
package bitmap

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Client Redis位图操作客户端
type Client struct {
	rdb *redis.Client
}

// New 创建位图操作客户端
func New(rdb *redis.Client) *Client {
	return &Client{rdb: rdb}
}

// SetBit 设置或清除指定偏移量上的位(bit)
func (c *Client) SetBit(ctx context.Context, key string, offset int64, value int) (int64, error) {
	return c.rdb.SetBit(ctx, key, offset, value).Result()
}

// GetBit 获取指定偏移量上的位(bit)
func (c *Client) GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	return c.rdb.GetBit(ctx, key, offset).Result()
}

// BitCount 计算给定字符串中，被设置为1的比特位的数量
func (c *Client) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) (int64, error) {
	return c.rdb.BitCount(ctx, key, bitCount).Result()
}

// BitOpAnd 对一个或多个保存二进制位的字符串key进行位元操作，并将结果保存到destkey上
// 进行AND运算
func (c *Client) BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return c.rdb.BitOpAnd(ctx, destKey, keys...).Result()
}

// BitOpOr 对一个或多个保存二进制位的字符串key进行位元操作，并将结果保存到destkey上
// 进行OR运算
func (c *Client) BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return c.rdb.BitOpOr(ctx, destKey, keys...).Result()
}

// BitOpXor 对一个或多个保存二进制位的字符串key进行位元操作，并将结果保存到destkey上
// 进行XOR运算
func (c *Client) BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return c.rdb.BitOpXor(ctx, destKey, keys...).Result()
}

// BitOpNot 对给定key进行位元操作，并将结果保存到destkey上
// 进行NOT运算
func (c *Client) BitOpNot(ctx context.Context, destKey string, key string) (int64, error) {
	return c.rdb.BitOpNot(ctx, destKey, key).Result()
}

// BitPos 返回位图中第一个值为bit的二进制位的位置
func (c *Client) BitPos(ctx context.Context, key string, bit int64, pos ...int64) (int64, error) {
	return c.rdb.BitPos(ctx, key, bit, pos...).Result()
}

// BitField 对字符串进行任意位长度和偏移量的位域操作
func (c *Client) BitField(ctx context.Context, key string, args ...interface{}) ([]int64, error) {
	return c.rdb.BitField(ctx, key, args...).Result()
}
