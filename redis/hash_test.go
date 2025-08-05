// @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:57:35
package redis

import (
	"testing"
)

func Test_hashClient(t *testing.T) {

	//1.初始化链接
	InitRedis()

	//2.运行测试
	t.Run("redis hash客户端测试", func(t *testing.T) {
		h := HashClient

		//1.写入单个键值对（存在则覆盖）
		count, err := h.HSet("hash_key", "field1", "value1")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//2.再次写入相同字段（覆盖）
		count, err = h.HSet("hash_key", "field1", "new_value1")
		if count != 0 || err != nil {
			t.Error(err)
		}

		//3.写入多个键值对
		count, err = h.HSet("hash_key", "field2", "value2", "field3", "value3")
		if count != 2 || err != nil {
			t.Error(err)
		}

		//4.测试HGet获取单个字段
		testHGetAndExists(t, h, "hash_key", "field1", "new_value1", true)
		testHGetAndExists(t, h, "hash_key", "field2", "value2", true)
		testHGetAndExists(t, h, "hash_key", "field3", "value3", true)

		//5.写入键值对（存在不覆盖）
		ok, err := h.HSetNX("hash_key", "field4", "value4")
		if !ok || err != nil {
			t.Error(err)
		}

		//6.再次写入相同字段（不覆盖）
		ok, err = h.HSetNX("hash_key", "field4", "new_value4")
		if ok || err != nil {
			t.Error(err)
		}

		//7.验证字段值未被覆盖
		testHGetAndExists(t, h, "hash_key", "field4", "value4", true)

		//8.获取多个键值对
		values, err := h.HMGet("hash_key", "field1", "field2", "field4", "nonexistent")
		if len(values) != 4 || values[0] != "new_value1" || values[1] != "value2" || values[2] != "value4" || values[3] != nil || err != nil {
			t.Error("HMGet结果不符合预期")
		}

		//9.获取所有键值对
		allFields, err := h.HGetAll("hash_key")
		if len(allFields) != 4 || allFields["field1"] != "new_value1" || allFields["field2"] != "value2" || allFields["field3"] != "value3" || allFields["field4"] != "value4" || err != nil {
			t.Error("HGetAll结果不符合预期")
		}

		//10.获取所有键
		keys, err := h.HKeys("hash_key")
		if len(keys) != 4 || err != nil {
			t.Error("HKeys结果不符合预期")
		}
		testContainsAll(t, keys, []string{"field1", "field2", "field3", "field4"})

		//11.获取所有值
		vals, err := h.HVals("hash_key")
		if len(vals) != 4 || err != nil {
			t.Error("HVals结果不符合预期")
		}
		testContainsAll(t, vals, []string{"new_value1", "value2", "value3", "value4"})

		//12.获取键值对数量
		length, err := h.HLen("hash_key")
		if length != 4 || err != nil {
			t.Error("HLen结果不符合预期")
		}

		//13.获取值的长度
		strLen, err := h.HStrLen("hash_key", "field1")
		if strLen != int64(len("new_value1")) || err != nil {
			t.Error("HStrLen结果不符合预期")
		}

		//14.测试不存在字段的长度
		strLen, err = h.HStrLen("hash_key", "nonexistent")
		if strLen != 0 || err != nil {
			t.Error("不存在字段的HStrLen结果不符合预期")
		}

		//15.删除单个键值对
		delCount, err := h.HDel("hash_key", "field3")
		if delCount != 1 || err != nil {
			t.Error(err)
		}

		//16.验证字段已删除
		testHGetAndExists(t, h, "hash_key", "field3", "", false)

		//17.删除多个键值对
		delCount, err = h.HDel("hash_key", "field1", "field2", "nonexistent")
		if delCount != 2 || err != nil {
			t.Error(err)
		}

		//18.验证字段已删除
		testHGetAndExists(t, h, "hash_key", "field1", "", false)
		testHGetAndExists(t, h, "hash_key", "field2", "", false)

		//19.验证剩余字段数量
		length, err = h.HLen("hash_key")
		if length != 1 || err != nil {
			t.Error("删除后HLen结果不符合预期")
		}

		//20.测试数值操作 - 设置初始值
		count, err = h.HSet("hash_key", "counter", "10")
		if err != nil {
			t.Error(err)
		}

		//21.整数递增
		incrResult, err := h.HIncrBy("hash_key", "counter", 5)
		if incrResult != 15 || err != nil {
			t.Error("HIncrBy结果不符合预期")
		}

		//22.整数递减（负数递增）
		incrResult, err = h.HIncrBy("hash_key", "counter", -3)
		if incrResult != 12 || err != nil {
			t.Error("HIncrBy负数结果不符合预期")
		}

		//23.设置浮点数初始值
		count, err = h.HSet("hash_key", "float_counter", "10.5")
		if err != nil {
			t.Error(err)
		}

		//24.浮点数递增
		floatResult, err := h.HIncrByFloat("hash_key", "float_counter", 2.3)
		if floatResult != 12.8 || err != nil {
			t.Error("HIncrByFloat结果不符合预期")
		}

		//25.浮点数递减
		floatResult, err = h.HIncrByFloat("hash_key", "float_counter", -1.8)
		if floatResult != 11.0 || err != nil {
			t.Error("HIncrByFloat负数结果不符合预期")
		}

		//26.对不存在的字段进行数值操作
		incrResult, err = h.HIncrBy("hash_key", "new_counter", 100)
		if incrResult != 100 || err != nil {
			t.Error("不存在字段的HIncrBy结果不符合预期")
		}

		//27.清理测试数据 - 删除整个hash
		delCount, err = h.HDel("hash_key", "field4", "counter", "float_counter", "new_counter")
		if delCount != 4 || err != nil {
			t.Error("清理数据失败")
		}

		//28.验证hash已清空
		length, err = h.HLen("hash_key")
		if length != 0 || err != nil {
			t.Error("hash未完全清空")
		}

		//29.测试空hash的各种操作
		testEmptyHashOperations(t, h, "empty_hash")
	})
}

// 测试HGet和HExists
func testHGetAndExists(t *testing.T, h *hashClient, key, field, expectedValue string, shouldExist bool) {

	//1.测试字段是否存在
	exists, err := h.HExists(key, field)
	if exists != shouldExist || err != nil {
		t.Errorf("HExists结果不符合预期: field=%s, expected=%v, actual=%v", field, shouldExist, exists)
	}

	//2.获取字段值
	value, err := h.HGet(key, field)
	if shouldExist {
		if value != expectedValue || err != nil {
			t.Errorf("HGet结果不符合预期: field=%s, expected=%s, actual=%s", field, expectedValue, value)
		}
	} else {
		if value != "" || err == nil || err.Error() != "redis: nil" {
			t.Errorf("不存在字段的HGet结果不符合预期: field=%s", field)
		}
	}
}

// 测试切片是否包含所有期望的元素
func testContainsAll(t *testing.T, actual []string, expected []string) {

	//1.如果长度不相等，则返回错误
	if len(actual) != len(expected) {
		t.Errorf("长度不匹配: expected=%d, actual=%d", len(expected), len(actual))
		return
	}

	//2.创建一个映射，用于存储期望的元素
	expectedMap := make(map[string]bool)
	for _, item := range expected {
		expectedMap[item] = true
	}

	//3.遍历实际结果，验证是否包含所有期望的元素
	for _, item := range actual {
		if !expectedMap[item] {
			t.Errorf("未期望的元素: %s", item)
		}
	}
}

// 测试空hash的各种操作
func testEmptyHashOperations(t *testing.T, h *hashClient, key string) {

	//1.获取不存在的字段
	value, err := h.HGet(key, "nonexistent")
	if value != "" || err == nil || err.Error() != "redis: nil" {
		t.Error("空hash的HGet结果不符合预期")
	}

	//2.检查字段是否存在
	exists, err := h.HExists(key, "nonexistent")
	if exists || err != nil {
		t.Error("空hash的HExists结果不符合预期")
	}

	//3.获取hash长度
	length, err := h.HLen(key)
	if length != 0 || err != nil {
		t.Error("空hash的HLen结果不符合预期")
	}

	//4.获取所有键
	keys, err := h.HKeys(key)
	if len(keys) != 0 || err != nil {
		t.Error("空hash的HKeys结果不符合预期")
	}

	//5.获取所有值
	vals, err := h.HVals(key)
	if len(vals) != 0 || err != nil {
		t.Error("空hash的HVals结果不符合预期")
	}

	//6.获取所有键值对
	allFields, err := h.HGetAll(key)
	if len(allFields) != 0 || err != nil {
		t.Error("空hash的HGetAll结果不符合预期")
	}

	//7.删除不存在的字段
	delCount, err := h.HDel(key, "nonexistent")
	if delCount != 0 || err != nil {
		t.Error("空hash的HDel结果不符合预期")
	}
}
