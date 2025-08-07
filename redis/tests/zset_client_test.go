// @Author:冯铁城 [17615007230@163.com] 2025-08-07 11:30:00
package redis_test

import (
	"context"
	"strconv"
	"testing"

	redisv9 "github.com/redis/go-redis/v9"
	"go-redis-demo/redis"
	zsetpkg "go-redis-demo/redis/zset"
)

func Test_zSetClient(t *testing.T) {

	//1.初始化链接
	config := redis.DefaultConfig()
	redis.InitClient(config)
	defer redis.CloseClient()

	//2.运行测试
	t.Run("redis zset客户端测试", func(t *testing.T) {
		z := redis.Client.ZSet
		ctx := context.Background()

		//1.添加单个元素到有序集合
		count, err := z.ZAdd(ctx, "zset_key", redisv9.Z{Score: 1.0, Member: "member1"})
		if count != 1 || err != nil {
			t.Error(err)
		}

		//2.添加多个元素到有序集合
		count, err = z.ZAdd(ctx, "zset_key", redisv9.Z{Score: 2.0, Member: "member2"}, redisv9.Z{Score: 3.0, Member: "member3"})
		if count != 2 || err != nil {
			t.Error(err)
		}

		//3.添加已存在的元素（更新分数）
		count, err = z.ZAdd(ctx, "zset_key", redisv9.Z{Score: 4.0, Member: "member1"})
		if count != 0 || err != nil {
			t.Error(err)
		}

		//4.验证有序集合元素数量
		card, err := z.ZCard(ctx, "zset_key")
		if card != 3 || err != nil {
			t.Error("ZCard结果不符合预期")
		}

		//5.获取有序集合所有元素（从小到大）
		members, err := z.ZRange(ctx, "zset_key", 0, -1)
		if len(members) != 3 || err != nil {
			t.Error("ZRange结果不符合预期")
		}
		testSliceEqualsZSet(t, members, []string{"member2", "member3", "member1"})

		//6.获取有序集合所有元素及分数（从小到大）
		membersWithScores, err := z.ZRangeWithScores(ctx, "zset_key", 0, -1)
		if len(membersWithScores) != 3 || err != nil {
			t.Error("ZRangeWithScores结果不符合预期")
		}
		testZSetElementsWithScores(t, membersWithScores, []redisv9.Z{
			{Score: 2.0, Member: "member2"},
			{Score: 3.0, Member: "member3"},
			{Score: 4.0, Member: "member1"},
		})

		//7.获取有序集合所有元素（从大到小）
		members, err = z.ZRevRange(ctx, "zset_key", 0, -1)
		if len(members) != 3 || err != nil {
			t.Error("ZRevRange结果不符合预期")
		}
		testSliceEqualsZSet(t, members, []string{"member1", "member3", "member2"})

		//8.获取有序集合所有元素及分数（从大到小）
		membersWithScores, err = z.ZRevRangeWithScores(ctx, "zset_key", 0, -1)
		if len(membersWithScores) != 3 || err != nil {
			t.Error("ZRevRangeWithScores结果不符合预期")
		}
		testZSetElementsWithScores(t, membersWithScores, []redisv9.Z{
			{Score: 4.0, Member: "member1"},
			{Score: 3.0, Member: "member3"},
			{Score: 2.0, Member: "member2"},
		})

		//9.获取有序集合部分元素（从小到大）
		members, err = z.ZRange(ctx, "zset_key", 0, 1)
		if len(members) != 2 || err != nil {
			t.Error("ZRange部分结果不符合预期")
		}
		testSliceEqualsZSet(t, members, []string{"member2", "member3"})

		//10.获取有序集合部分元素（从大到小）
		members, err = z.ZRevRange(ctx, "zset_key", 0, 1)
		if len(members) != 2 || err != nil {
			t.Error("ZRevRange部分结果不符合预期")
		}
		testSliceEqualsZSet(t, members, []string{"member1", "member3"})

		//11.对有序集合中元素的分数增加增量
		newScore, err := z.ZIncrBy(ctx, "zset_key", 2.5, "member2")
		if newScore != 4.5 || err != nil {
			t.Error("ZIncrBy结果不符合预期")
		}

		//12.对有序集合中元素的分数减少增量（负数增量）
		newScore, err = z.ZIncrBy(ctx, "zset_key", -1.5, "member3")
		if newScore != 1.5 || err != nil {
			t.Error("ZIncrBy负数结果不符合预期")
		}

		//13.对不存在的元素进行增量操作（自动添加）
		newScore, err = z.ZIncrBy(ctx, "zset_key", 1.0, "member4")
		if newScore != 1.0 || err != nil {
			t.Error("ZIncrBy不存在元素结果不符合预期")
		}

		//14.验证有序集合元素数量增加
		card, err = z.ZCard(ctx, "zset_key")
		if card != 4 || err != nil {
			t.Error("ZIncrBy后ZCard结果不符合预期")
		}

		//15.获取有序集合元素的排名（从小到大）
		rank, err := z.ZRank(ctx, "zset_key", "member3")
		if rank != 1 || err != nil {
			t.Error("ZRank结果不符合预期")
		}

		//16.获取有序集合元素的排名（从大到小）
		rank, err = z.ZRevRank(ctx, "zset_key", "member3")
		if rank != 2 || err != nil {
			t.Error("ZRevRank结果不符合预期")
		}

		//17.获取不存在元素的排名
		_, err = z.ZRank(ctx, "zset_key", "nonexistent")
		if err == nil || err.Error() != "redis: nil" {
			t.Error("ZRank不存在元素结果不符合预期")
		}

		//18.按分数范围获取元素（从小到大）
		members, err = z.ZRangeByScore(ctx, "zset_key", "1", "2", 0, 0)
		if len(members) != 2 || err != nil {
			t.Error("ZRangeByScore结果不符合预期")
		}
		testSetContainsAllZSet(t, members, []string{"member3", "member4"})

		//19.按分数范围获取元素及分数（从小到大）
		membersWithScores, err = z.ZRangeByScoreWithScores(ctx, "zset_key", "1", "2", 0, 0)
		if len(membersWithScores) != 2 || err != nil {
			t.Error("ZRangeByScoreWithScores结果不符合预期")
		}
		for _, m := range membersWithScores {
			if m.Member == "member3" && m.Score != 1.5 {
				t.Error("ZRangeByScoreWithScores member3分数不符合预期")
			}
			if m.Member == "member4" && m.Score != 1.0 {
				t.Error("ZRangeByScoreWithScores member4分数不符合预期")
			}
		}

		//20.按分数范围获取元素（从大到小）
		members, err = z.ZRevRangeByScore(ctx, "zset_key", "5", "3", 0, 0)
		if len(members) != 2 || err != nil {
			t.Error("ZRevRangeByScore结果不符合预期")
		}
		testSetContainsAllZSet(t, members, []string{"member1", "member2"})

		//21.按分数范围获取元素及分数（从大到小）
		membersWithScores, err = z.ZRevRangeByScoreWithScores(ctx, "zset_key", "5", "3", 0, 0)
		if len(membersWithScores) != 2 || err != nil {
			t.Error("ZRevRangeByScoreWithScores结果不符合预期")
		}
		for _, m := range membersWithScores {
			if m.Member == "member1" && m.Score != 4.0 {
				t.Error("ZRevRangeByScoreWithScores member1分数不符合预期")
			}
			if m.Member == "member2" && m.Score != 4.5 {
				t.Error("ZRevRangeByScoreWithScores member2分数不符合预期")
			}
		}

		//22.按分数范围获取元素（带偏移和限制）
		members, err = z.ZRangeByScore(ctx, "zset_key", "0", "5", 1, 2)
		if len(members) != 2 || err != nil {
			t.Error("ZRangeByScore带偏移和限制结果不符合预期")
		}

		//23.获取指定分数范围内的元素数量
		count, err = z.ZCount(ctx, "zset_key", "1", "2")
		if count != 2 || err != nil {
			t.Error("ZCount结果不符合预期")
		}

		//24.删除有序集合中的元素
		count, err = z.ZRem(ctx, "zset_key", "member3")
		if count != 1 || err != nil {
			t.Error("ZRem结果不符合预期")
		}

		//25.删除不存在的元素
		count, err = z.ZRem(ctx, "zset_key", "nonexistent")
		if count != 0 || err != nil {
			t.Error("ZRem不存在元素结果不符合预期")
		}

		//26.验证有序集合元素数量减少
		card, err = z.ZCard(ctx, "zset_key")
		if card != 3 || err != nil {
			t.Error("ZRem后ZCard结果不符合预期")
		}

		//27.删除有序集合中指定排名范围的元素
		count, err = z.ZRemRangeByRank(ctx, "zset_key", 0, 0)
		if count != 1 || err != nil {
			t.Error("ZRemRangeByRank结果不符合预期")
		}

		//28.验证有序集合元素数量减少
		card, err = z.ZCard(ctx, "zset_key")
		if card != 2 || err != nil {
			t.Error("ZRemRangeByRank后ZCard结果不符合预期")
		}

		//29.准备测试数据
		_, err = z.ZAdd(ctx, "zset_score_key", redisv9.Z{Score: 1.0, Member: "a"}, redisv9.Z{Score: 2.0, Member: "b"}, redisv9.Z{Score: 3.0, Member: "c"}, redisv9.Z{Score: 4.0, Member: "d"}, redisv9.Z{Score: 5.0, Member: "e"})
		if err != nil {
			t.Error(err)
		}

		//30.删除有序集合中指定分数范围的元素
		count, err = z.ZRemRangeByScore(ctx, "zset_score_key", "2", "4")
		if count != 3 || err != nil {
			t.Error("ZRemRangeByScore结果不符合预期")
		}

		//31.验证有序集合元素数量减少
		card, err = z.ZCard(ctx, "zset_score_key")
		if card != 2 || err != nil {
			t.Error("ZRemRangeByScore后ZCard结果不符合预期")
		}

		//32.验证剩余元素
		members, err = z.ZRange(ctx, "zset_score_key", 0, -1)
		if len(members) != 2 || err != nil {
			t.Error("ZRemRangeByScore后ZRange结果不符合预期")
		}
		testSetContainsAllZSet(t, members, []string{"a", "e"})

		//33.清理测试数据
		cleanupZSets(t, z, ctx, []string{"zset_key", "zset_score_key"})

		//34.测试空有序集合的各种操作
		testEmptyZSetOperations(t, z, ctx, "empty_zset")

		//35.测试大批量数据
		testBulkZSetOperations(t, z, ctx, "bulk_zset", 100)
	})
}

// 测试有序集合元素及分数
func testZSetElementsWithScores(t *testing.T, actual []redisv9.Z, expected []redisv9.Z) {

	//1.验证长度是否相等
	if len(actual) != len(expected) {
		t.Errorf("有序集合长度不符合预期: expected=%d, actual=%d", len(expected), len(actual))
		return
	}

	//2.验证每个元素及其分数
	for i, exp := range expected {
		act := actual[i]
		if act.Member != exp.Member || act.Score != exp.Score {
			t.Errorf("有序集合元素不符合预期: index=%d, expected={%v, %v}, actual={%v, %v}",
				i, exp.Member, exp.Score, act.Member, act.Score)
		}
	}
}

// 测试集合是否包含所有期望的元素
func testSetContainsAllZSet(t *testing.T, actual []string, expected []string) {

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

// 测试两个切片是否相等
func testSliceEqualsZSet(t *testing.T, actual []string, expected []string) {

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

// 清理测试有序集合
func cleanupZSets(t *testing.T, z *zsetpkg.Client, ctx context.Context, keys []string) {

	//1.遍历所有key
	for _, key := range keys {

		//2.获取有序集合所有元素
		members, err := z.ZRange(ctx, key, 0, -1)
		if err != nil {
			t.Error(err)
			continue
		}

		//3.如果有序集合不为空，则删除所有元素
		if len(members) > 0 {
			interfaceMembers := make([]interface{}, len(members))
			for i, m := range members {
				interfaceMembers[i] = m
			}
			_, err = z.ZRem(ctx, key, interfaceMembers...)
			if err != nil {
				t.Error(err)
			}
		}

		//4.验证有序集合已清空
		card, err := z.ZCard(ctx, key)
		if card != 0 || err != nil {
			t.Error("有序集合未完全清空")
		}
	}
}

// 测试空有序集合的各种操作
func testEmptyZSetOperations(t *testing.T, z *zsetpkg.Client, ctx context.Context, key string) {

	//1.获取有序集合元素数量
	card, err := z.ZCard(ctx, key)
	if card != 0 || err != nil {
		t.Error("空有序集合的ZCard结果不符合预期")
	}

	//2.获取有序集合所有元素
	members, err := z.ZRange(ctx, key, 0, -1)
	if len(members) != 0 || err != nil {
		t.Error("空有序集合的ZRange结果不符合预期")
	}

	//3.获取有序集合所有元素及分数
	membersWithScores, err := z.ZRangeWithScores(ctx, key, 0, -1)
	if len(membersWithScores) != 0 || err != nil {
		t.Error("空有序集合的ZRangeWithScores结果不符合预期")
	}

	//4.获取元素排名
	_, err = z.ZRank(ctx, key, "nonexistent")
	if err == nil || err.Error() != "redis: nil" {
		t.Error("空有序集合的ZRank结果不符合预期")
	}

	//5.删除元素
	count, err := z.ZRem(ctx, key, "nonexistent")
	if count != 0 || err != nil {
		t.Error("空有序集合的ZRem结果不符合预期")
	}

	//6.按分数范围获取元素数量
	count, err = z.ZCount(ctx, key, "-inf", "+inf")
	if count != 0 || err != nil {
		t.Error("空有序集合的ZCount结果不符合预期")
	}
}

// 测试大批量数据
func testBulkZSetOperations(t *testing.T, z *zsetpkg.Client, ctx context.Context, key string, count int) {

	//1.准备批量添加的数据
	members := make([]redisv9.Z, count)
	for i := 0; i < count; i++ {
		members[i] = redisv9.Z{
			Score:  float64(i),
			Member: "member" + strconv.Itoa(i),
		}
	}

	//2.批量添加数据
	addCount, err := z.ZAdd(ctx, key, members...)
	if int(addCount) != count || err != nil {
		t.Error("批量ZAdd结果不符合预期")
	}

	//3.验证有序集合元素数量
	card, err := z.ZCard(ctx, key)
	if int(card) != count || err != nil {
		t.Error("批量添加后ZCard结果不符合预期")
	}

	//4.获取部分元素
	partMembers, err := z.ZRange(ctx, key, 0, 9)
	if len(partMembers) != 10 || err != nil {
		t.Error("批量数据ZRange结果不符合预期")
	}

	//5.按分数范围获取元素数量
	rangeCount, err := z.ZCount(ctx, key, "10", "20")
	if rangeCount != 11 || err != nil {
		t.Error("批量数据ZCount结果不符合预期")
	}

	//6.清理测试数据
	cleanupZSets(t, z, ctx, []string{key})
}
