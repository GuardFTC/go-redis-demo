// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:29:24
package redis

import "github.com/redis/go-redis/v9"

// ZSetClient redis有序集合操作
var ZSetClient = new(zSetClient)

// zSetClient redis有序集合操作
type zSetClient struct {
}

// ZAdd 添加多个元素到有序集合中
func (z *zSetClient) ZAdd(key string, members ...redis.Z) (int64, error) {
	return rdb.ZAdd(ctx, key, members...).Result()
}

// ZIncrBy 对有序集合中指定成员的分数加上增量increment
func (z *zSetClient) ZIncrBy(key string, increment float64, member string) (float64, error) {
	return rdb.ZIncrBy(ctx, key, increment, member).Result()
}

// ZRange 查询有序集合，指定区间内的元素（从小到大排序）
func (z *zSetClient) ZRange(key string, start, stop int64) ([]string, error) {
	return rdb.ZRange(ctx, key, start, stop).Result()
}

// ZRangeWithScores 查询有序集合，指定区间内的元素及其分数（从小到大排序）
func (z *zSetClient) ZRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	result, err := rdb.ZRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZRevRange 查询有序集合，指定区间内的元素（从大到小排序）
func (z *zSetClient) ZRevRange(key string, start, stop int64) ([]string, error) {
	return rdb.ZRevRange(ctx, key, start, stop).Result()
}

// ZRevRangeWithScores 查询有序集合，指定区间内的元素及其分数（从大到小排序）
func (z *zSetClient) ZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	result, err := rdb.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZRem 删除有序集合中的一个或多个成员
func (z *zSetClient) ZRem(key string, members ...interface{}) (int64, error) {
	return rdb.ZRem(ctx, key, members...).Result()
}

// ZCard 获取有序集合的成员数
func (z *zSetClient) ZCard(key string) (int64, error) {
	return rdb.ZCard(ctx, key).Result()
}

// ZRangeByScore 获取有序集合中指定分数区间的成员（从小到大排序）
func (z *zSetClient) ZRangeByScore(key string, min, max string, offset, count int64) ([]string, error) {
	opt := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	return rdb.ZRangeByScore(ctx, key, &opt).Result()
}

// ZRangeByScoreWithScores 获取有序集合中指定分数区间的成员及其分数（从小到大排序）
func (z *zSetClient) ZRangeByScoreWithScores(key string, min, max string, offset, count int64) ([]redis.Z, error) {
	opt := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	result, err := rdb.ZRangeByScoreWithScores(ctx, key, &opt).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZRevRangeByScore 获取有序集合中指定分数区间的成员（从大到小排序）
func (z *zSetClient) ZRevRangeByScore(key string, max, min string, offset, count int64) ([]string, error) {
	opt := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	return rdb.ZRevRangeByScore(ctx, key, &opt).Result()
}

// ZRevRangeByScoreWithScores 获取有序集合中指定分数区间的成员及其分数（从大到小排序）
func (z *zSetClient) ZRevRangeByScoreWithScores(key string, max, min string, offset, count int64) ([]redis.Z, error) {
	opt := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}
	result, err := rdb.ZRevRangeByScoreWithScores(ctx, key, &opt).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZCount 获取有序集合中指定分数区间的成员数量
func (z *zSetClient) ZCount(key, min, max string) (int64, error) {
	return rdb.ZCount(ctx, key, min, max).Result()
}

// ZRemRangeByRank 删除有序集合中指定排名区间的成员
func (z *zSetClient) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	return rdb.ZRemRangeByRank(ctx, key, start, stop).Result()
}

// ZRemRangeByScore 删除有序集合中指定分数区间的成员
func (z *zSetClient) ZRemRangeByScore(key, min, max string) (int64, error) {
	return rdb.ZRemRangeByScore(ctx, key, min, max).Result()
}

// ZRank 获取有序集合中成员的排名（从小到大，0表示第一个元素）
func (z *zSetClient) ZRank(key, member string) (int64, error) {
	return rdb.ZRank(ctx, key, member).Result()
}

// ZRevRank 获取有序集合中成员的排名（从大到小，0表示第一个元素）
func (z *zSetClient) ZRevRank(key, member string) (int64, error) {
	return rdb.ZRevRank(ctx, key, member).Result()
}
