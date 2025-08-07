// Package set 提供Redis集合操作的封装
// @Author:冯铁城 [17615007230@163.com] 2025-08-07 10:00:00
package set

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Client Redis集合操作客户端
type Client struct {
	rdb *redis.Client
}

// New 创建集合操作客户端
func New(rdb *redis.Client) *Client {
	return &Client{rdb: rdb}
}

// SAdd 添加若干指定元素member到key集合中，并返回成功添加元素个数
func (c *Client) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return c.rdb.SAdd(ctx, key, members...).Result()
}

// SPop 随机移除并返回集合key中若干随机元素
func (c *Client) SPop(ctx context.Context, key string, count ...int64) ([]string, error) {
	if len(count) > 0 {
		return c.rdb.SPopN(ctx, key, count[0]).Result()
	}
	result, err := c.rdb.SPop(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []string{result}, nil
}

// SRem 在集合key中移除指定元素，并返回成功移除元素个数
func (c *Client) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return c.rdb.SRem(ctx, key, members...).Result()
}

// SCard 返回指定集合key中的元素数
func (c *Client) SCard(ctx context.Context, key string) (int64, error) {
	return c.rdb.SCard(ctx, key).Result()
}

// SIsMember 返回集合key中是否存在指定元素member
func (c *Client) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return c.rdb.SIsMember(ctx, key, member).Result()
}

// SMembers 返回集合key的所有元素
func (c *Client) SMembers(ctx context.Context, key string) ([]string, error) {
	return c.rdb.SMembers(ctx, key).Result()
}

// SRandMember 随机返回集合key中的一个元素，或随机返回集合key中的count的元素
func (c *Client) SRandMember(ctx context.Context, key string, count ...int64) ([]string, error) {
	if len(count) > 0 {
		return c.rdb.SRandMemberN(ctx, key, count[0]).Result()
	}
	result, err := c.rdb.SRandMember(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []string{result}, nil
}

// SMove 将指定元素member从集合source中移动到集合destination中
func (c *Client) SMove(ctx context.Context, source, destination string, member interface{}) (bool, error) {
	return c.rdb.SMove(ctx, source, destination, member).Result()
}

// SInter 返回所有指定集合中元素的交集
func (c *Client) SInter(ctx context.Context, keys ...string) ([]string, error) {
	return c.rdb.SInter(ctx, keys...).Result()
}

// SInterStore 返回所有指定集合中元素的交集，并将结果保存在集合destination中
func (c *Client) SInterStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return c.rdb.SInterStore(ctx, destination, keys...).Result()
}

// SUnion 返回所有指定集合中元素的并集
func (c *Client) SUnion(ctx context.Context, keys ...string) ([]string, error) {
	return c.rdb.SUnion(ctx, keys...).Result()
}

// SUnionStore 返回所有指定集合中元素的并集，并将结果保存在集合destination中
func (c *Client) SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return c.rdb.SUnionStore(ctx, destination, keys...).Result()
}

// SDiff 返回一个集合与其余指定集合的差集
func (c *Client) SDiff(ctx context.Context, keys ...string) ([]string, error) {
	return c.rdb.SDiff(ctx, keys...).Result()
}

// SDiffStore 返回一个集合与其余指定集合的差集，并将结果保存在集合destination中
func (c *Client) SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return c.rdb.SDiffStore(ctx, destination, keys...).Result()
}
