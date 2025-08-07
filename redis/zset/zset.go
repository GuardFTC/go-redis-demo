// Package zset 提供Redis有序集合操作的封装
// @Author:冯铁城 [17615007230@163.com] 2025-08-07 10:00:00
package zset

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Client Redis有序集合操作客户端
type Client struct {
	rdb *redis.Client
}

// New 创建有序集合操作客户端
func New(rdb *redis.Client) *Client {
	return &Client{rdb: rdb}
}

// ZAdd 添加多个元素到有序集合中
func (c *Client) ZAdd(ctx context.Context, key string, members ...redis.Z) (int64, error) {
	return c.rdb.ZAdd(ctx, key, members...).Result()
}

// ZIncrBy 对有序集合中指定成员的分数加上增量increment
func (c *Client) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	return c.rdb.ZIncrBy(ctx, key, increment, member).Result()
}

// ZRange 查询有序集合，指定区间内的元素（从小到大排序）
func (c *Client) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return c.rdb.ZRange(ctx, key, start, stop).Result()
}

// ZRangeWithScores 查询有序集合，指定区间内的元素及其分数（从小到大排序）
func (c *Client) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	result, err := c.rdb.ZRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZRevRange 查询有序集合，指定区间内的元素（从大到小排序）
func (c *Client) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return c.rdb.ZRevRange(ctx, key, start, stop).Result()
}

// ZRevRangeWithScores 查询有序集合，指定区间内的元素及其分数（从大到小排序）
func (c *Client) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	result, err := c.rdb.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZRem 删除有序集合中的一个或多个成员
func (c *Client) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return c.rdb.ZRem(ctx, key, members...).Result()
}

// ZCard 获取有序集合的成员数
func (c *Client) ZCard(ctx context.Context, key string) (int64, error) {
	return c.rdb.ZCard(ctx, key).Result()
}

// ZRangeByScore 获取有序集合中指定分数区间的成员（从小到大排序）
func (c *Client) ZRangeByScore(ctx context.Context, key string, min, max string, offset, count int64) ([]string, error) {
	opt := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	return c.rdb.ZRangeByScore(ctx, key, &opt).Result()
}

// ZRangeByScoreWithScores 获取有序集合中指定分数区间的成员及其分数（从小到大排序）
func (c *Client) ZRangeByScoreWithScores(ctx context.Context, key string, min, max string, offset, count int64) ([]redis.Z, error) {
	opt := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	result, err := c.rdb.ZRangeByScoreWithScores(ctx, key, &opt).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZRevRangeByScore 获取有序集合中指定分数区间的成员（从大到小排序）
func (c *Client) ZRevRangeByScore(ctx context.Context, key string, max, min string, offset, count int64) ([]string, error) {
	opt := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	return c.rdb.ZRevRangeByScore(ctx, key, &opt).Result()
}

// ZRevRangeByScoreWithScores 获取有序集合中指定分数区间的成员及其分数（从大到小排序）
func (c *Client) ZRevRangeByScoreWithScores(ctx context.Context, key string, max, min string, offset, count int64) ([]redis.Z, error) {
	opt := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	result, err := c.rdb.ZRevRangeByScoreWithScores(ctx, key, &opt).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZCount 获取有序集合中指定分数区间的成员数量
func (c *Client) ZCount(ctx context.Context, key, min, max string) (int64, error) {
	return c.rdb.ZCount(ctx, key, min, max).Result()
}

// ZRemRangeByRank 删除有序集合中指定排名区间的成员
func (c *Client) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	return c.rdb.ZRemRangeByRank(ctx, key, start, stop).Result()
}

// ZRemRangeByScore 删除有序集合中指定分数区间的成员
func (c *Client) ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error) {
	return c.rdb.ZRemRangeByScore(ctx, key, min, max).Result()
}

// ZRank 获取有序集合中成员的排名（从小到大，0表示第一个元素）
func (c *Client) ZRank(ctx context.Context, key, member string) (int64, error) {
	return c.rdb.ZRank(ctx, key, member).Result()
}

// ZRevRank 获取有序集合中成员的排名（从大到小，0表示第一个元素）
func (c *Client) ZRevRank(ctx context.Context, key, member string) (int64, error) {
	return c.rdb.ZRevRank(ctx, key, member).Result()
}
