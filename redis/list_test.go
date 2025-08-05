// @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:57:35
package redis

import (
	"testing"
	"time"
)

func Test_listClient(t *testing.T) {

	//1.初始化链接
	InitRedis()

	//2.运行测试
	t.Run("redis list客户端测试", func(t *testing.T) {
		l := ListClient

		//1.左端推入元素（Key不存在创建Key）
		count, err := l.LPush("list_key", "value1")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//2.左端推入多个元素
		count, err = l.LPush("list_key", "value2", "value3")
		if count != 3 || err != nil {
			t.Error(err)
		}

		//3.验证列表长度
		length, err := l.LLen("list_key")
		if length != 3 || err != nil {
			t.Error("LLen结果不符合预期")
		}

		//4.验证列表元素顺序（左端推入，元素应该是倒序）
		testListElements(t, l, "list_key", []string{"value3", "value2", "value1"})

		//5.右端推入元素（Key不存在创建Key）
		count, err = l.RPush("list_key2", "value1")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//6.右端推入多个元素
		count, err = l.RPush("list_key2", "value2", "value3")
		if count != 3 || err != nil {
			t.Error(err)
		}

		//7.验证列表长度
		length, err = l.LLen("list_key2")
		if length != 3 || err != nil {
			t.Error("LLen结果不符合预期")
		}

		//8.验证列表元素顺序（右端推入，元素应该是正序）
		testListElements(t, l, "list_key2", []string{"value1", "value2", "value3"})

		//9.左端推入元素（Key不存在不做操作）
		count, err = l.LPushX("nonexistent_key", "value1")
		if count != 0 || err != nil {
			t.Error("LPushX对不存在的key操作结果不符合预期")
		}

		//10.左端推入元素到已存在的key
		count, err = l.LPushX("list_key", "value4")
		if count != 4 || err != nil {
			t.Error("LPushX结果不符合预期")
		}

		//11.右端推入元素（Key不存在不做操作）
		count, err = l.RPushX("nonexistent_key", "value1")
		if count != 0 || err != nil {
			t.Error("RPushX对不存在的key操作结果不符合预期")
		}

		//12.右端推入元素到已存在的key
		count, err = l.RPushX("list_key", "value5")
		if count != 5 || err != nil {
			t.Error("RPushX结果不符合预期")
		}

		//13.验证列表元素顺序
		testListElements(t, l, "list_key", []string{"value4", "value3", "value2", "value1", "value5"})

		//14.左端弹出元素
		value, err := l.LPop("list_key")
		if value != "value4" || err != nil {
			t.Error("LPop结果不符合预期")
		}

		//15.右端弹出元素
		value, err = l.RPop("list_key")
		if value != "value5" || err != nil {
			t.Error("RPop结果不符合预期")
		}

		//16.验证列表长度
		length, err = l.LLen("list_key")
		if length != 3 || err != nil {
			t.Error("弹出后LLen结果不符合预期")
		}

		//17.验证列表元素顺序
		testListElements(t, l, "list_key", []string{"value3", "value2", "value1"})

		//18.获取索引处的元素
		value, err = l.LIndex("list_key", 0)
		if value != "value3" || err != nil {
			t.Error("LIndex结果不符合预期")
		}

		//19.获取索引处的元素（负数索引）
		value, err = l.LIndex("list_key", -1)
		if value != "value1" || err != nil {
			t.Error("LIndex负数索引结果不符合预期")
		}

		//20.获取不存在的索引
		value, err = l.LIndex("list_key", 100)
		if value != "" || err == nil || err.Error() != "redis: nil" {
			t.Error("LIndex不存在索引结果不符合预期")
		}

		//21.在目标元素前插入元素
		count, err = l.LInsert("list_key", "BEFORE", "value2", "value2.5")
		if count != 4 || err != nil {
			t.Error("LInsert BEFORE结果不符合预期")
		}

		//22.在目标元素后插入元素
		count, err = l.LInsert("list_key", "AFTER", "value2", "value1.5")
		if count != 5 || err != nil {
			t.Error("LInsert AFTER结果不符合预期")
		}

		//23.验证列表元素顺序
		testListElements(t, l, "list_key", []string{"value3", "value2.5", "value2", "value1.5", "value1"})

		//24.在不存在的目标元素前插入元素
		count, err = l.LInsert("list_key", "BEFORE", "nonexistent", "value")
		if count != -1 || err != nil {
			t.Error("LInsert不存在元素结果不符合预期")
		}

		//25.获取指定范围的元素
		values, err := l.LRange("list_key", 1, 3)
		if len(values) != 3 || values[0] != "value2.5" || values[1] != "value2" || values[2] != "value1.5" || err != nil {
			t.Error("LRange结果不符合预期")
		}

		//26.获取全部元素
		values, err = l.LRange("list_key", 0, -1)
		if len(values) != 5 || err != nil {
			t.Error("LRange全部元素结果不符合预期")
		}

		//27.删除指定元素（删除1个）
		count, err = l.LRem("list_key", 1, "value2.5")
		if count != 1 || err != nil {
			t.Error("LRem结果不符合预期")
		}

		//28.验证列表元素顺序
		testListElements(t, l, "list_key", []string{"value3", "value2", "value1.5", "value1"})

		//29.更新指定下标的值
		err = l.LSet("list_key", 1, "new_value2")
		if err != nil {
			t.Error("LSet结果不符合预期")
		}

		//30.验证列表元素顺序
		testListElements(t, l, "list_key", []string{"value3", "new_value2", "value1.5", "value1"})

		//31.更新不存在的下标
		err = l.LSet("list_key", 100, "value")
		if err == nil {
			t.Error("LSet不存在下标应该返回错误")
		}

		//32.裁剪list
		err = l.LTrim("list_key", 1, 2)
		if err != nil {
			t.Error("LTrim结果不符合预期")
		}

		//33.验证列表元素顺序
		testListElements(t, l, "list_key", []string{"new_value2", "value1.5"})

		//34.右边弹出，左边推入
		// 先准备源list和目标list
		_, err = l.RPush("source_list", "s1", "s2", "s3")
		if err != nil {
			t.Error(err)
		}
		_, err = l.RPush("dest_list", "d1", "d2")
		if err != nil {
			t.Error(err)
		}

		//35.执行右边弹出，左边推入
		value, err = l.RPopLPush("source_list", "dest_list")
		if value != "s3" || err != nil {
			t.Error("RPopLPush结果不符合预期")
		}

		//36.验证源list和目标list
		testListElements(t, l, "source_list", []string{"s1", "s2"})
		testListElements(t, l, "dest_list", []string{"s3", "d1", "d2"})

		//37.测试阻塞操作（使用较短的超时时间）
		// 注意：这些操作在实际环境中可能会阻塞，所以使用短超时
		timeout := 1 * time.Second

		//38.阻塞式右边弹出，左边推入（有数据）
		value, err = l.BRPopLPush("source_list", "dest_list", timeout)
		if value != "s2" || err != nil {
			t.Error("BRPopLPush结果不符合预期")
		}

		//39.阻塞式右边弹出，左边推入（无数据，超时）
		_, err = l.RPush("empty_source", "temp")
		if err != nil {
			t.Error(err)
		}
		_, err = l.LPop("empty_source")
		if err != nil {
			t.Error(err)
		}
		value, err = l.BRPopLPush("empty_source", "dest_list", 1*time.Second)
		if err == nil {
			t.Error("BRPopLPush空list应该超时")
		}

		//40.阻塞式左端弹出（有数据）
		result, err := l.BLPop(timeout, "source_list")
		if len(result) != 2 || result[0] != "source_list" || result[1] != "s1" || err != nil {
			t.Error("BLPop结果不符合预期")
		}

		//41.阻塞式右端弹出（有数据）
		_, err = l.RPush("brpop_test", "b1", "b2")
		if err != nil {
			t.Error(err)
		}
		result, err = l.BRPop(timeout, "brpop_test")
		if len(result) != 2 || result[0] != "brpop_test" || result[1] != "b2" || err != nil {
			t.Error("BRPop结果不符合预期")
		}

		//42.阻塞式左端弹出（无数据，超时）
		result, err = l.BLPop(1*time.Second, "nonexistent_key")
		if err == nil {
			t.Error("BLPop空list应该超时")
		}

		//43.阻塞式右端弹出（无数据，超时）
		result, err = l.BRPop(1*time.Second, "nonexistent_key")
		if err == nil {
			t.Error("BRPop空list应该超时")
		}

		//44.清理测试数据
		cleanupLists(t, l, []string{"list_key", "list_key2", "source_list", "dest_list", "empty_source", "brpop_test"})

		//45.测试空list的各种操作
		testEmptyListOperations(t, l, "empty_list")
	})
}

// 测试列表元素
func testListElements(t *testing.T, l *listClient, key string, expectedValues []string) {

	//1.获取列表所有元素
	values, err := l.LRange(key, 0, -1)
	if err != nil {
		t.Error(err)
		return
	}

	//2.验证列表长度
	if len(values) != len(expectedValues) {
		t.Errorf("列表长度不符合预期: key=%s, expected=%d, actual=%d", key, len(expectedValues), len(values))
		return
	}

	//3.验证每个元素
	for i, expected := range expectedValues {
		if values[i] != expected {
			t.Errorf("列表元素不符合预期: key=%s, index=%d, expected=%s, actual=%s", key, i, expected, values[i])
		}
	}
}

// 清理测试列表
func cleanupLists(t *testing.T, l *listClient, keys []string) {

	//1.遍历所有key
	for _, key := range keys {

		//2.获取列表长度
		length, err := l.LLen(key)
		if err != nil {
			t.Error(err)
			continue
		}

		//3.如果列表不为空，则清空列表
		if length > 0 {
			err = l.LTrim(key, 1, 0) // 裁剪为空列表
			if err != nil {
				t.Error(err)
			}
		}

		//4.验证列表已清空
		length, err = l.LLen(key)
		if length != 0 || err != nil {
			t.Error("列表未完全清空")
		}
	}
}

// 测试空list的各种操作
func testEmptyListOperations(t *testing.T, l *listClient, key string) {

	//1.获取列表长度
	length, err := l.LLen(key)
	if length != 0 || err != nil {
		t.Error("空list的LLen结果不符合预期")
	}

	//2.获取列表元素
	values, err := l.LRange(key, 0, -1)
	if len(values) != 0 || err != nil {
		t.Error("空list的LRange结果不符合预期")
	}

	//3.获取索引处的元素
	value, err := l.LIndex(key, 0)
	if value != "" || err == nil || err.Error() != "redis: nil" {
		t.Error("空list的LIndex结果不符合预期")
	}

	//4.左端弹出元素
	value, err = l.LPop(key)
	if value != "" || err == nil || err.Error() != "redis: nil" {
		t.Error("空list的LPop结果不符合预期")
	}

	//5.右端弹出元素
	value, err = l.RPop(key)
	if value != "" || err == nil || err.Error() != "redis: nil" {
		t.Error("空list的RPop结果不符合预期")
	}

	//6.更新指定下标的值
	err = l.LSet(key, 0, "value")
	if err == nil {
		t.Error("空list的LSet应该返回错误")
	}

	//7.删除指定元素
	count, err := l.LRem(key, 1, "value")
	if count != 0 || err != nil {
		t.Error("空list的LRem结果不符合预期")
	}

	//8.右边弹出，左边推入
	value, err = l.RPopLPush(key, "dest_list")
	if value != "" || err == nil || err.Error() != "redis: nil" {
		t.Error("空list的RPopLPush结果不符合预期")
	}
}
