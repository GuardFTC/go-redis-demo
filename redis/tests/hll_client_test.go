// @Author:冯铁城 [17615007230@163.com] 2025-08-07 11:30:00
package redis_test

import (
	"context"
	"testing"

	"go-redis-demo/redis"
	hllpkg "go-redis-demo/redis/hll"
)

func Test_hyperLogLogClient(t *testing.T) {

	//1.初始化链接
	config := redis.DefaultConfig()
	redis.InitClient(config)
	defer redis.CloseClient()

	//2.运行测试
	t.Run("redis hll客户端测试", func(t *testing.T) {
		h := redis.Client.HLL
		ctx := context.Background()

		//1.添加单个元素到HyperLogLog
		count, err := h.PFAdd(ctx, "hll_key", "element1")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//2.添加多个元素到HyperLogLog
		count, err = h.PFAdd(ctx, "hll_key", "element2", "element3", "element4")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//3.添加重复元素到HyperLogLog
		count, err = h.PFAdd(ctx, "hll_key", "element1", "element2", "element5")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//4.获取HyperLogLog的基数估算值
		card, err := h.PFCount(ctx, "hll_key")
		if card != 5 || err != nil {
			t.Error("PFCount结果不符合预期")
		}

		//5.创建另一个HyperLogLog
		count, err = h.PFAdd(ctx, "hll_key2", "element4", "element5", "element6", "element7")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//6.获取第二个HyperLogLog的基数估算值
		card, err = h.PFCount(ctx, "hll_key2")
		if card != 4 || err != nil {
			t.Error("PFCount结果不符合预期")
		}

		//7.获取多个HyperLogLog的合并基数估算值
		card, err = h.PFCount(ctx, "hll_key", "hll_key2")
		if card != 7 || err != nil {
			t.Error("多个HyperLogLog的PFCount结果不符合预期")
		}

		//8.合并多个HyperLogLog
		err = h.PFMerge(ctx, "hll_merged", "hll_key", "hll_key2")
		if err != nil {
			t.Error(err)
		}

		//9.获取合并后HyperLogLog的基数估算值
		card, err = h.PFCount(ctx, "hll_merged")
		if card != 7 || err != nil {
			t.Error("合并后PFCount结果不符合预期")
		}

		//10.添加新元素到合并后的HyperLogLog
		count, err = h.PFAdd(ctx, "hll_merged", "element8", "element9")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//11.获取添加新元素后的基数估算值
		card, err = h.PFCount(ctx, "hll_merged")
		if card != 9 || err != nil {
			t.Error("添加新元素后PFCount结果不符合预期")
		}

		//12.测试不存在的HyperLogLog
		card, err = h.PFCount(ctx, "nonexistent_hll")
		if card != 0 || err != nil {
			t.Error("不存在的HyperLogLog的PFCount结果不符合预期")
		}

		//13.合并包含不存在的HyperLogLog
		err = h.PFMerge(ctx, "hll_merged2", "hll_key", "nonexistent_hll")
		if err != nil {
			t.Error(err)
		}

		//14.获取包含不存在HyperLogLog的合并结果
		card, err = h.PFCount(ctx, "hll_merged2")
		if card != 5 || err != nil {
			t.Error("包含不存在HyperLogLog的合并结果不符合预期")
		}

		//15.测试大量元素
		testBulkHyperLogLog(t, h, ctx, "hll_bulk", 1000)

		//16.清理测试数据
		cleanupHLLKeys(t, ctx, []string{"hll_key", "hll_key2", "hll_merged", "hll_merged2", "hll_bulk"})
	})
}

// 测试大量元素的HyperLogLog
func testBulkHyperLogLog(t *testing.T, h *hllpkg.Client, ctx context.Context, key string, count int) {

	//1.准备批量添加的数据
	elements := make([]interface{}, count)
	for i := 0; i < count; i++ {
		elements[i] = "bulk_element_" + string(rune(i))
	}

	//2.批量添加数据
	_, err := h.PFAdd(ctx, key, elements...)
	if err != nil {
		t.Error("批量PFAdd失败")
	}

	//3.验证基数估算值
	// 注意：HyperLogLog是一种概率数据结构，有一定的误差
	// 这里我们只验证基数估算值是否在合理范围内
	card, err := h.PFCount(ctx, key)
	if err != nil {
		t.Error(err)
	}

	//4.验证基数估算值在合理范围内
	// HyperLogLog的标准误差约为0.81%
	errorMargin := float64(count) * 0.02 // 使用2%作为容错范围
	if float64(card) < float64(count)-errorMargin || float64(card) > float64(count)+errorMargin {
		t.Errorf("批量数据PFCount结果不在合理范围内: expected=%d±%f, actual=%d", count, errorMargin, card)
	}
}

// 清理测试数据
func cleanupHLLKeys(t *testing.T, ctx context.Context, keys []string) {

	//1.删除所有测试使用的key
	_, err := redis.Client.String.Del(ctx, keys...)
	if err != nil {
		t.Error("清理测试数据失败")
	}

	//2.验证key已被删除
	for _, key := range keys {
		exists, err := redis.Client.String.Exists(ctx, key)
		if exists != 0 || err != nil {
			t.Errorf("key未被成功删除: %s", key)
		}
	}
}
