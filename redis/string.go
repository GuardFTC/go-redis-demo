// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:29:24
package redis

import (
	"time"
)

// StringClient redis字符串操作
var StringClient = new(stringClient)

// string redis字符串操作
type stringClient struct {
}

// SetWithDefaultExpire 设置 key,使用默认过期时间（存在则覆盖）
func (s *stringClient) SetWithDefaultExpire(key string, value interface{}) error {
	return s.Set(key, value, DefaultExpire)
}

// Set 设置 key（存在则覆盖）
func (s *stringClient) Set(key string, value interface{}, expiration time.Duration) error {
	return rdb.Set(ctx, key, value, expiration).Err()
}

// SetNXWithDefaultExpire 设置 key,使用默认过期时间（存在不覆盖）
func (s *stringClient) SetNXWithDefaultExpire(key string, value interface{}) (bool, error) {
	return s.SetNX(key, value, DefaultExpire)
}

// SetNX 设置 key（存在不覆盖）
func (s *stringClient) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return rdb.SetNX(ctx, key, value, expiration).Result()
}

// Get 获取 key
func (s *stringClient) Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

// MSet 设置多个 key-value（存在则覆盖）
func (s *stringClient) MSet(pairs ...interface{}) error {
	return rdb.MSet(ctx, pairs...).Err()
}

// MSetNX 设置多个 key-value（存在不覆盖）
func (s *stringClient) MSetNX(pairs ...interface{}) (bool, error) {
	return rdb.MSetNX(ctx, pairs...).Result()
}

// Keys 获取所有匹配的 key（慎用）
func (s *stringClient) Keys(pattern string) ([]string, error) {
	return rdb.Keys(ctx, pattern).Result()
}

// Exists 判断 key 是否存在（返回匹配个数）
func (s *stringClient) Exists(keys ...string) (int64, error) {
	return rdb.Exists(ctx, keys...).Result()
}

// TTL 获取 key 剩余 TTL（秒级）
func (s *stringClient) TTL(key string) (time.Duration, error) {
	return rdb.TTL(ctx, key).Result()
}

// Expire 设置过期时间（秒）
func (s *stringClient) Expire(key string, expiration time.Duration) error {
	return rdb.Expire(ctx, key, expiration).Err()
}

// PExpire 设置过期时间（毫秒）
func (s *stringClient) PExpire(key string, expiration time.Duration) error {
	return rdb.PExpire(ctx, key, expiration).Err()
}

// Incr 增加 key 的值（整数）
func (s *stringClient) Incr(key string) (int64, error) {
	return rdb.Incr(ctx, key).Result()
}

// IncrBy 按指定值增加
func (s *stringClient) IncrBy(key string, increment int64) (int64, error) {
	return rdb.IncrBy(ctx, key, increment).Result()
}

// Decr 减少 key 的值
func (s *stringClient) Decr(key string) (int64, error) {
	return rdb.Decr(ctx, key).Result()
}

// DecrBy 按指定值减少
func (s *stringClient) DecrBy(key string, decrement int64) (int64, error) {
	return rdb.DecrBy(ctx, key, decrement).Result()
}

// Del 删除 key
func (s *stringClient) Del(keys ...string) (int64, error) {
	return rdb.Del(ctx, keys...).Result()
}
