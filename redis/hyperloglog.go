// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:29:24
package redis

// HyperLogLogClient redis基数统计操作
var HyperLogLogClient = new(hyperLogLogClient)

// hyperLogLogClient redis基数统计操作
type hyperLogLogClient struct {
}

// PFAdd 添加指定元素到HyperLogLog中
func (h *hyperLogLogClient) PFAdd(key string, elements ...interface{}) (int64, error) {
	return rdb.PFAdd(ctx, key, elements...).Result()
}

// PFCount 返回给定HyperLogLog的基数估算值
func (h *hyperLogLogClient) PFCount(keys ...string) (int64, error) {
	return rdb.PFCount(ctx, keys...).Result()
}

// PFMerge 将多个HyperLogLog合并为一个HyperLogLog
func (h *hyperLogLogClient) PFMerge(destKey string, sourceKeys ...string) error {
	return rdb.PFMerge(ctx, destKey, sourceKeys...).Err()
}
