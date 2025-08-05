// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 19:30:35
package redis

import (
	"errors"
	"github.com/redis/go-redis/v9"
)

// BitMapClient redis位图操作
var BitMapClient = new(bitMapClient)

// bitMapClient redis位图操作
type bitMapClient struct {
}

// SetBit 设置指定key在offset处的bit值
func (b *bitMapClient) SetBit(key string, offset int64, value int) (int64, error) {
	if value != 0 && value != 1 {
		return 0, errors.New("value必须为0或1")
	}

	return rdb.SetBit(ctx, key, offset, value).Result()
}

// GetBit 获取指定key在offset处的bit值
func (b *bitMapClient) GetBit(key string, offset int64) (int64, error) {
	return rdb.GetBit(ctx, key, offset).Result()
}

// BitCount 统计指定key中值为1的bit数量
func (b *bitMapClient) BitCount(key string, start, end int64) (int64, error) {
	return rdb.BitCount(ctx, key, &redis.BitCount{
		Start: start,
		End:   end,
	}).Result()
}

// BitPos 返回指定key中第一个值为targetBit(0或1)的bit位置
func (b *bitMapClient) BitPos(key string, targetBit int64, start, end int64) (int64, error) {
	if targetBit != 0 && targetBit != 1 {
		return 0, errors.New("targetBit必须为0或1")
	}

	return rdb.BitPos(ctx, key, targetBit, start, end).Result()
}

// BitOpAND 对一个或多个key执行按位与操作，并将结果存储到destKey
func (b *bitMapClient) BitOpAND(destKey string, keys ...string) (int64, error) {
	return rdb.BitOpAnd(ctx, destKey, keys...).Result()
}

// BitOpOR 对一个或多个key执行按位或操作，并将结果存储到destKey
func (b *bitMapClient) BitOpOR(destKey string, keys ...string) (int64, error) {
	return rdb.BitOpOr(ctx, destKey, keys...).Result()
}

// BitOpXOR 对一个或多个key执行按位异或操作，并将结果存储到destKey
func (b *bitMapClient) BitOpXOR(destKey string, keys ...string) (int64, error) {
	return rdb.BitOpXor(ctx, destKey, keys...).Result()
}

// BitOpNOT 对指定key执行按位非操作，并将结果存储到destKey
func (b *bitMapClient) BitOpNOT(destKey string, key string) (int64, error) {
	return rdb.BitOpNot(ctx, destKey, key).Result()
}
