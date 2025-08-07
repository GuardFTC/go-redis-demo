// @Author:冯铁城 [17615007230@163.com] 2025-08-07 11:30:00
package redis_test

import (
	"context"
	"sort"
	"testing"

	"go-redis-demo/redis"
	setpkg "go-redis-demo/redis/set"
)

func Test_setClient(t *testing.T) {

	//1.初始化链接
	config := redis.DefaultConfig()
	redis.InitClient(config)
	defer redis.CloseClient()

	//2.运行测试
	t.Run("redis set客户端测试", func(t *testing.T) {
		s := redis.Client.Set
		ctx := context.Background()

		//1.添加单个元素
		count, err := s.SAdd(ctx, "set_key", "value1")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//2.添加多个元素
		count, err = s.SAdd(ctx, "set_key", "value2", "value3", "value4")
		if count != 3 || err != nil {
			t.Error(err)
		}

		//3.添加重复元素
		count, err = s.SAdd(ctx, "set_key", "value1", "value2", "value5")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//4.验证集合元素数量
		card, err := s.SCard(ctx, "set_key")
		if card != 5 || err != nil {
			t.Error("SCard结果不符合预期")
		}

		//5.验证元素是否存在
		exists, err := s.SIsMember(ctx, "set_key", "value1")
		if !exists || err != nil {
			t.Error("SIsMember结果不符合预期")
		}

		//6.验证元素不存在
		exists, err = s.SIsMember(ctx, "set_key", "nonexistent")
		if exists || err != nil {
			t.Error("SIsMember不存在元素结果不符合预期")
		}

		//7.获取所有元素
		members, err := s.SMembers(ctx, "set_key")
		if len(members) != 5 || err != nil {
			t.Error("SMembers结果不符合预期")
		}
		testSetContainsAll(t, members, []string{"value1", "value2", "value3", "value4", "value5"})

		//8.随机获取一个元素
		randMembers, err := s.SRandMember(ctx, "set_key")
		if len(randMembers) != 1 || err != nil {
			t.Error("SRandMember结果不符合预期")
		}
		testSetContains(t, []string{"value1", "value2", "value3", "value4", "value5"}, randMembers[0])

		//9.随机获取多个元素
		randMembers, err = s.SRandMember(ctx, "set_key", 3)
		if len(randMembers) != 3 || err != nil {
			t.Error("SRandMember多个元素结果不符合预期")
		}
		for _, member := range randMembers {
			testSetContains(t, []string{"value1", "value2", "value3", "value4", "value5"}, member)
		}

		//10.随机弹出一个元素
		popMembers, err := s.SPop(ctx, "set_key")
		if len(popMembers) != 1 || err != nil {
			t.Error("SPop结果不符合预期")
		}

		//11.验证集合元素数量减少
		card, err = s.SCard(ctx, "set_key")
		if card != 4 || err != nil {
			t.Error("SPop后SCard结果不符合预期")
		}

		//12.随机弹出多个元素
		popMembers, err = s.SPop(ctx, "set_key", 2)
		if len(popMembers) != 2 || err != nil {
			t.Error("SPop多个元素结果不符合预期")
		}

		//13.验证集合元素数量减少
		card, err = s.SCard(ctx, "set_key")
		if card != 2 || err != nil {
			t.Error("SPop多个元素后SCard结果不符合预期")
		}

		//14.移除元素
		count, err = s.SRem(ctx, "set_key", "nonexistent")
		if count != 0 || err != nil {
			t.Error("SRem不存在元素结果不符合预期")
		}

		//15.获取剩余元素
		members, err = s.SMembers(ctx, "set_key")
		if len(members) != 2 || err != nil {
			t.Error("SRem后SMembers结果不符合预期")
		}

		//16.移除存在的元素
		count, err = s.SRem(ctx, "set_key", members[0])
		if count != 1 || err != nil {
			t.Error("SRem结果不符合预期")
		}

		//17.验证集合元素数量减少
		card, err = s.SCard(ctx, "set_key")
		if card != 1 || err != nil {
			t.Error("SRem后SCard结果不符合预期")
		}

		//18.准备集合操作测试
		_, err = s.SAdd(ctx, "set_key1", "a", "b", "c", "d")
		if err != nil {
			t.Error(err)
		}
		_, err = s.SAdd(ctx, "set_key2", "c", "d", "e", "f")
		if err != nil {
			t.Error(err)
		}

		//19.测试集合交集
		inter, err := s.SInter(ctx, "set_key1", "set_key2")
		if err != nil {
			t.Error(err)
		}
		sort.Strings(inter)
		testSliceEquals(t, inter, []string{"c", "d"})

		//20.测试集合交集并存储
		count, err = s.SInterStore(ctx, "set_inter", "set_key1", "set_key2")
		if count != 2 || err != nil {
			t.Error("SInterStore结果不符合预期")
		}
		inter, err = s.SMembers(ctx, "set_inter")
		if err != nil {
			t.Error(err)
		}
		sort.Strings(inter)
		testSliceEquals(t, inter, []string{"c", "d"})

		//21.测试集合并集
		union, err := s.SUnion(ctx, "set_key1", "set_key2")
		if err != nil {
			t.Error(err)
		}
		sort.Strings(union)
		testSliceEquals(t, union, []string{"a", "b", "c", "d", "e", "f"})

		//22.测试集合并集并存储
		count, err = s.SUnionStore(ctx, "set_union", "set_key1", "set_key2")
		if count != 6 || err != nil {
			t.Error("SUnionStore结果不符合预期")
		}
		union, err = s.SMembers(ctx, "set_union")
		if err != nil {
			t.Error(err)
		}
		sort.Strings(union)
		testSliceEquals(t, union, []string{"a", "b", "c", "d", "e", "f"})

		//23.测试集合差集
		diff, err := s.SDiff(ctx, "set_key1", "set_key2")
		if err != nil {
			t.Error(err)
		}
		sort.Strings(diff)
		testSliceEquals(t, diff, []string{"a", "b"})

		//24.测试集合差集并存储
		count, err = s.SDiffStore(ctx, "set_diff", "set_key1", "set_key2")
		if count != 2 || err != nil {
			t.Error("SDiffStore结果不符合预期")
		}
		diff, err = s.SMembers(ctx, "set_diff")
		if err != nil {
			t.Error(err)
		}
		sort.Strings(diff)
		testSliceEquals(t, diff, []string{"a", "b"})

		//25.测试元素移动
		_, err = s.SAdd(ctx, "set_source", "item1", "item2")
		if err != nil {
			t.Error(err)
		}
		_, err = s.SAdd(ctx, "set_dest", "item3")
		if err != nil {
			t.Error(err)
		}

		//26.移动元素
		moved, err := s.SMove(ctx, "set_source", "set_dest", "item1")
		if !moved || err != nil {
			t.Error("SMove结果不符合预期")
		}

		//27.验证源集合
		members, err = s.SMembers(ctx, "set_source")
		if len(members) != 1 || members[0] != "item2" || err != nil {
			t.Error("SMove后源集合结果不符合预期")
		}

		//28.验证目标集合
		members, err = s.SMembers(ctx, "set_dest")
		if err != nil {
			t.Error(err)
		}
		testSetContainsAll(t, members, []string{"item1", "item3"})

		//29.移动不存在的元素
		moved, err = s.SMove(ctx, "set_source", "set_dest", "nonexistent")
		if moved || err != nil {
			t.Error("SMove不存在元素结果不符合预期")
		}

		//30.清理测试数据
		cleanupSets(t, s, ctx, []string{"set_key", "set_key1", "set_key2", "set_inter", "set_union", "set_diff", "set_source", "set_dest"})

		//31.测试空集合的各种操作
		testEmptySetOperations(t, s, ctx, "empty_set")
	})
}

// 测试集合是否包含所有期望的元素
func testSetContainsAll(t *testing.T, actual []string, expected []string) {

	//1.创建一个映射，用于存储实际结果
	actualMap := make(map[string]bool)
	for _, item := range actual {
		actualMap[item] = true
	}

	//2.验证所有期望的元素都在实际结果中
	for _, item := range expected {
		if !actualMap[item] {
			t.Errorf("集合中缺少期望的元素: %s", item)
		}
	}

	//3.验证实际结果中没有多余的元素
	if len(actual) != len(expected) {
		t.Errorf("集合元素数量不符合预期: expected=%d, actual=%d", len(expected), len(actual))
	}
}

// 测试集合是否包含指定元素
func testSetContains(t *testing.T, set []string, element string) {

	//1.遍历集合查找元素
	for _, item := range set {
		if item == element {
			return
		}
	}

	//2.如果没有找到元素，则报错
	t.Errorf("集合中不包含期望的元素: %s", element)
}

// 测试两个切片是否相等
func testSliceEquals(t *testing.T, actual []string, expected []string) {

	//1.验证长度是否相等
	if len(actual) != len(expected) {
		t.Errorf("切片长度不符合预期: expected=%d, actual=%d", len(expected), len(actual))
		return
	}

	//2.验证每个元素是否相等
	for i, item := range actual {
		if item != expected[i] {
			t.Errorf("切片元素不符合预期: index=%d, expected=%s, actual=%s", i, expected[i], item)
		}
	}
}

// 清理测试集合
func cleanupSets(t *testing.T, s *setpkg.Client, ctx context.Context, keys []string) {

	//1.遍历所有key
	for _, key := range keys {

		//2.获取集合所有元素
		members, err := s.SMembers(ctx, key)
		if err != nil {
			t.Error(err)
			continue
		}

		//3.如果集合不为空，则删除所有元素
		if len(members) > 0 {
			_, err = s.SRem(ctx, key, interfaceSlice(members)...)
			if err != nil {
				t.Error(err)
			}
		}

		//4.验证集合已清空
		card, err := s.SCard(ctx, key)
		if card != 0 || err != nil {
			t.Error("集合未完全清空")
		}
	}
}

// 将字符串切片转换为接口切片
func interfaceSlice(slice []string) []interface{} {

	//1.创建接口切片
	result := make([]interface{}, len(slice))

	//2.将字符串转换为接口
	for i, v := range slice {
		result[i] = v
	}

	//3.返回结果
	return result
}

// 测试空集合的各种操作
func testEmptySetOperations(t *testing.T, s *setpkg.Client, ctx context.Context, key string) {

	//1.获取集合元素数量
	card, err := s.SCard(ctx, key)
	if card != 0 || err != nil {
		t.Error("空集合的SCard结果不符合预期")
	}

	//2.获取集合所有元素
	members, err := s.SMembers(ctx, key)
	if len(members) != 0 || err != nil {
		t.Error("空集合的SMembers结果不符合预期")
	}

	//3.检查元素是否存在
	exists, err := s.SIsMember(ctx, key, "nonexistent")
	if exists || err != nil {
		t.Error("空集合的SIsMember结果不符合预期")
	}

	//4.随机获取元素
	randMembers, err := s.SRandMember(ctx, key)
	if len(randMembers) != 0 || err == nil || err.Error() != "redis: nil" {
		t.Error("空集合的SRandMember结果不符合预期")
	}

	//5.随机弹出元素
	popMembers, err := s.SPop(ctx, key)
	if len(popMembers) != 0 || err == nil || err.Error() != "redis: nil" {
		t.Error("空集合的SPop结果不符合预期")
	}

	//6.移除元素
	count, err := s.SRem(ctx, key, "nonexistent")
	if count != 0 || err != nil {
		t.Error("空集合的SRem结果不符合预期")
	}
}
