// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:29:24
package redis

import (
	"time"
)

// ListClient redis列表操作
var ListClient = new(listClient)

// listClient redis列表操作
type listClient struct {
}

// LPush 左端推入元素（Key不存在创建Key）
func (l *listClient) LPush(key string, values ...interface{}) (int64, error) {
	return rdb.LPush(ctx, key, values...).Result()
}

// LPushX 左端推入元素（Key不存在不做操作）
func (l *listClient) LPushX(key string, values ...interface{}) (int64, error) {
	return rdb.LPushX(ctx, key, values...).Result()
}

// RPush 右端推入元素（Key不存在创建Key）
func (l *listClient) RPush(key string, values ...interface{}) (int64, error) {
	return rdb.RPush(ctx, key, values...).Result()
}

// RPushX 右端推入元素（Key不存在不做操作）
func (l *listClient) RPushX(key string, values ...interface{}) (int64, error) {
	return rdb.RPushX(ctx, key, values...).Result()
}

// LPop 左端弹出
func (l *listClient) LPop(key string) (string, error) {
	return rdb.LPop(ctx, key).Result()
}

// RPop 右端弹出
func (l *listClient) RPop(key string) (string, error) {
	return rdb.RPop(ctx, key).Result()
}

// LIndex 返回索引处的元素
func (l *listClient) LIndex(key string, index int64) (string, error) {
	return rdb.LIndex(ctx, key, index).Result()
}

// LInsert 在目标元素前或后插入元素
func (l *listClient) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	return rdb.LInsert(ctx, key, op, pivot, value).Result()
}

// LRange 获取指定范围的元素
func (l *listClient) LRange(key string, start, stop int64) ([]string, error) {
	return rdb.LRange(ctx, key, start, stop).Result()
}

// LLen 获取集合长度
func (l *listClient) LLen(key string) (int64, error) {
	return rdb.LLen(ctx, key).Result()
}

// LRem 删除n个指定元素
func (l *listClient) LRem(key string, count int64, value interface{}) (int64, error) {
	return rdb.LRem(ctx, key, count, value).Result()
}

// LSet 更新指定下标的值
func (l *listClient) LSet(key string, index int64, value interface{}) error {
	return rdb.LSet(ctx, key, index, value).Err()
}

// LTrim 裁剪list
func (l *listClient) LTrim(key string, start, stop int64) error {
	return rdb.LTrim(ctx, key, start, stop).Err()
}

// RPopLPush 右边弹出，左边推入
func (l *listClient) RPopLPush(source, destination string) (string, error) {
	return rdb.RPopLPush(ctx, source, destination).Result()
}

// BRPopLPush 阻塞式右边弹出，左边推入
func (l *listClient) BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
	return rdb.BRPopLPush(ctx, source, destination, timeout).Result()
}

// BLPop 阻塞式左端弹出
func (l *listClient) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return rdb.BLPop(ctx, timeout, keys...).Result()
}

// BRPop 阻塞式右端弹出
func (l *listClient) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	return rdb.BRPop(ctx, timeout, keys...).Result()
}
