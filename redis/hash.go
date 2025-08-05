// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:29:24
package redis

// HashClient redis哈希操作
var HashClient = new(hashClient)

// hashClient redis哈希操作
type hashClient struct {
}

// HSet 写入键值对（存在则覆盖）
func (h *hashClient) HSet(key string, values ...interface{}) (int64, error) {
	return rdb.HSet(ctx, key, values...).Result()
}

// HSetNX 写入键值对（存在不覆盖）
func (h *hashClient) HSetNX(key, field string, value interface{}) (bool, error) {
	return rdb.HSetNX(ctx, key, field, value).Result()
}

// HGet 获取键值对
func (h *hashClient) HGet(key, field string) (string, error) {
	return rdb.HGet(ctx, key, field).Result()
}

// HMGet 获取多个键值对
func (h *hashClient) HMGet(key string, fields ...string) ([]interface{}, error) {
	return rdb.HMGet(ctx, key, fields...).Result()
}

// HGetAll 获取所有键值对
func (h *hashClient) HGetAll(key string) (map[string]string, error) {
	return rdb.HGetAll(ctx, key).Result()
}

// HKeys 获取所有键
func (h *hashClient) HKeys(key string) ([]string, error) {
	return rdb.HKeys(ctx, key).Result()
}

// HVals 获取所有值
func (h *hashClient) HVals(key string) ([]string, error) {
	return rdb.HVals(ctx, key).Result()
}

// HDel 删除键值对
func (h *hashClient) HDel(key string, fields ...string) (int64, error) {
	return rdb.HDel(ctx, key, fields...).Result()
}

// HExists 判断字段是否存在
func (h *hashClient) HExists(key, field string) (bool, error) {
	return rdb.HExists(ctx, key, field).Result()
}

// HLen 获取键值对数量
func (h *hashClient) HLen(key string) (int64, error) {
	return rdb.HLen(ctx, key).Result()
}

// HStrLen 获取值的长度
func (h *hashClient) HStrLen(key, field string) (int64, error) {
	return rdb.HStrLen(ctx, key, field).Result()
}

// HIncrBy 给字段的值加上一个整数（负数即为减法）
func (h *hashClient) HIncrBy(key, field string, incr int64) (int64, error) {
	return rdb.HIncrBy(ctx, key, field, incr).Result()
}

// HIncrByFloat 给字段的值加上一个数（可以是浮点数）
func (h *hashClient) HIncrByFloat(key, field string, incr float64) (float64, error) {
	return rdb.HIncrByFloat(ctx, key, field, incr).Result()
}
