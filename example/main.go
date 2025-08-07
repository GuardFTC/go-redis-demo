package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-redis-demo/redis"

	redisv9 "github.com/redis/go-redis/v9"
)

func main() {

	//1.初始化Redis客户端
	config := redis.DefaultConfig()
	redis.InitClient(config)
	defer redis.CloseClient()

	//2.创建上下文
	ctx := context.Background()

	//3.字符串操作示例
	stringExample(ctx)

	//4.哈希操作示例
	hashExample(ctx)

	//5.列表操作示例
	listExample(ctx)

	//6.集合操作示例
	setExample(ctx)

	//7.有序集合操作示例
	zsetExample(ctx)

	//8.地理位置操作示例
	geoExample(ctx)

	//9.位图操作示例
	bitmapExample(ctx)

	//10.HyperLogLog操作示例
	hllExample(ctx)

	fmt.Println("所有示例执行完成！")
}

// 字符串操作示例
func stringExample(ctx context.Context) {
	fmt.Println("=== 字符串操作示例 ===")

	//1.设置字符串
	err := redis.Client.String.Set(ctx, "name", "CodeBuddy", time.Hour)
	if err != nil {
		log.Printf("设置字符串失败: %v", err)
		return
	}

	//2.获取字符串
	name, err := redis.Client.String.Get(ctx, "name")
	if err != nil {
		log.Printf("获取字符串失败: %v", err)
		return
	}
	fmt.Printf("获取到的名称: %s\n", name)

	//3.自增操作
	count, err := redis.Client.String.Incr(ctx, "counter")
	if err != nil {
		log.Printf("自增操作失败: %v", err)
		return
	}
	fmt.Printf("计数器值: %d\n", count)

	//4.清理数据
	redis.Client.String.Del(ctx, "name", "counter")
	fmt.Println()
}

// 哈希操作示例
func hashExample(ctx context.Context) {
	fmt.Println("=== 哈希操作示例 ===")

	//1.设置哈希字段
	_, err := redis.Client.Hash.HSet(ctx, "user:1", "name", "张三", "age", "25", "city", "北京")
	if err != nil {
		log.Printf("设置哈希字段失败: %v", err)
		return
	}

	//2.获取单个字段
	name, err := redis.Client.Hash.HGet(ctx, "user:1", "name")
	if err != nil {
		log.Printf("获取哈希字段失败: %v", err)
		return
	}
	fmt.Printf("用户姓名: %s\n", name)

	//3.获取所有字段
	userInfo, err := redis.Client.Hash.HGetAll(ctx, "user:1")
	if err != nil {
		log.Printf("获取所有哈希字段失败: %v", err)
		return
	}
	fmt.Printf("用户信息: %v\n", userInfo)

	//4.清理数据
	redis.Client.Hash.HDel(ctx, "user:1", "name", "age", "city")
	fmt.Println()
}

// 列表操作示例
func listExample(ctx context.Context) {
	fmt.Println("=== 列表操作示例 ===")

	//1.左侧推入元素
	_, err := redis.Client.List.LPush(ctx, "tasks", "任务1", "任务2", "任务3")
	if err != nil {
		log.Printf("推入列表元素失败: %v", err)
		return
	}

	//2.获取列表范围
	tasks, err := redis.Client.List.LRange(ctx, "tasks", 0, -1)
	if err != nil {
		log.Printf("获取列表范围失败: %v", err)
		return
	}
	fmt.Printf("任务列表: %v\n", tasks)

	//3.弹出元素
	task, err := redis.Client.List.LPop(ctx, "tasks")
	if err != nil {
		log.Printf("弹出列表元素失败: %v", err)
		return
	}
	fmt.Printf("弹出的任务: %s\n", task)

	//4.清理数据
	redis.Client.List.LTrim(ctx, "tasks", 1, 0) // 清空列表
	fmt.Println()
}

// 集合操作示例
func setExample(ctx context.Context) {
	fmt.Println("=== 集合操作示例 ===")

	//1.添加元素到集合
	_, err := redis.Client.Set.SAdd(ctx, "tags", "Go", "Redis", "数据库", "缓存")
	if err != nil {
		log.Printf("添加集合元素失败: %v", err)
		return
	}

	//2.获取集合所有成员
	members, err := redis.Client.Set.SMembers(ctx, "tags")
	if err != nil {
		log.Printf("获取集合成员失败: %v", err)
		return
	}
	fmt.Printf("标签集合: %v\n", members)

	//3.判断元素是否存在
	exists, err := redis.Client.Set.SIsMember(ctx, "tags", "Go")
	if err != nil {
		log.Printf("判断集合成员失败: %v", err)
		return
	}
	fmt.Printf("Go标签是否存在: %v\n", exists)

	//4.清理数据
	redis.Client.Set.SRem(ctx, "tags", "Go", "Redis", "数据库", "缓存")
	fmt.Println()
}

// 有序集合操作示例
func zsetExample(ctx context.Context) {
	fmt.Println("=== 有序集合操作示例 ===")

	//1.添加元素到有序集合
	_, err := redis.Client.ZSet.ZAdd(ctx, "scores",
		redisv9.Z{Score: 95.5, Member: "张三"},
		redisv9.Z{Score: 87.0, Member: "李四"},
		redisv9.Z{Score: 92.5, Member: "王五"},
	)
	if err != nil {
		log.Printf("添加有序集合元素失败: %v", err)
		return
	}

	//2.获取排名（从高到低）
	topScores, err := redis.Client.ZSet.ZRevRangeWithScores(ctx, "scores", 0, 2)
	if err != nil {
		log.Printf("获取有序集合排名失败: %v", err)
		return
	}
	fmt.Printf("成绩排名: %v\n", topScores)

	//3.获取成员排名
	rank, err := redis.Client.ZSet.ZRevRank(ctx, "scores", "张三")
	if err != nil {
		log.Printf("获取成员排名失败: %v", err)
		return
	}
	fmt.Printf("张三的排名: %d\n", rank+1) // 排名从0开始，所以+1

	//4.清理数据
	redis.Client.ZSet.ZRem(ctx, "scores", "张三", "李四", "王五")
	fmt.Println()
}

// 地理位置操作示例
func geoExample(ctx context.Context) {
	fmt.Println("=== 地理位置操作示例 ===")

	//1.添加地理位置
	_, err := redis.Client.Geo.GeoAdd(ctx, "cities", 116.397128, 39.916527, "北京")
	if err != nil {
		log.Printf("添加地理位置失败: %v", err)
		return
	}

	_, err = redis.Client.Geo.GeoAdd(ctx, "cities", 121.473701, 31.230416, "上海")
	if err != nil {
		log.Printf("添加地理位置失败: %v", err)
		return
	}

	//2.计算距离
	distance, err := redis.Client.Geo.GeoDist(ctx, "cities", "北京", "上海", "km")
	if err != nil {
		log.Printf("计算距离失败: %v", err)
		return
	}
	fmt.Printf("北京到上海的距离: %.2f 公里\n", distance)

	//3.半径搜索
	locations, err := redis.Client.Geo.GeoRadius(ctx, "cities", 116.397128, 39.916527, 1500, "km", true, true, true, 10)
	if err != nil {
		log.Printf("半径搜索失败: %v", err)
		return
	}
	fmt.Printf("1500公里范围内的城市: %v\n", locations)

	//4.清理数据
	redis.Client.GetRawClient().Del(ctx, "cities")
	fmt.Println()
}

// 位图操作示例
func bitmapExample(ctx context.Context) {
	fmt.Println("=== 位图操作示例 ===")

	//1.模拟用户签到 - 用户在第1、3、5天签到
	_, err := redis.Client.Bitmap.SetBit(ctx, "user_sign:202501", 0, 1) // 第1天
	if err != nil {
		log.Printf("设置位图失败: %v", err)
		return
	}

	_, err = redis.Client.Bitmap.SetBit(ctx, "user_sign:202501", 2, 1) // 第3天
	if err != nil {
		log.Printf("设置位图失败: %v", err)
		return
	}

	_, err = redis.Client.Bitmap.SetBit(ctx, "user_sign:202501", 4, 1) // 第5天
	if err != nil {
		log.Printf("设置位图失败: %v", err)
		return
	}

	//2.统计签到天数
	count, err := redis.Client.Bitmap.BitCount(ctx, "user_sign:202501", &redisv9.BitCount{Start: 0, End: -1})
	if err != nil {
		log.Printf("统计位图失败: %v", err)
		return
	}
	fmt.Printf("用户1月份签到天数: %d\n", count)

	//3.检查特定日期是否签到
	signed, err := redis.Client.Bitmap.GetBit(ctx, "user_sign:202501", 2) // 检查第3天
	if err != nil {
		log.Printf("获取位图失败: %v", err)
		return
	}
	fmt.Printf("用户第3天是否签到: %v\n", signed == 1)

	//4.清理数据
	redis.Client.String.Del(ctx, "user_sign:202501")
	fmt.Println()
}

// HyperLogLog操作示例
func hllExample(ctx context.Context) {
	fmt.Println("=== HyperLogLog操作示例 ===")

	//1.添加访客
	_, err := redis.Client.HLL.PFAdd(ctx, "unique_visitors", "user1", "user2", "user3", "user1") // user1重复
	if err != nil {
		log.Printf("添加HyperLogLog元素失败: %v", err)
		return
	}

	//2.获取独立访客数估算
	count, err := redis.Client.HLL.PFCount(ctx, "unique_visitors")
	if err != nil {
		log.Printf("获取HyperLogLog基数失败: %v", err)
		return
	}
	fmt.Printf("独立访客数估算: %d\n", count)

	//3.添加更多访客到另一个HLL
	_, err = redis.Client.HLL.PFAdd(ctx, "unique_visitors_2", "user3", "user4", "user5")
	if err != nil {
		log.Printf("添加HyperLogLog元素失败: %v", err)
		return
	}

	//4.合并两个HLL
	err = redis.Client.HLL.PFMerge(ctx, "merged_visitors", "unique_visitors", "unique_visitors_2")
	if err != nil {
		log.Printf("合并HyperLogLog失败: %v", err)
		return
	}

	//5.获取合并后的独立访客数
	mergedCount, err := redis.Client.HLL.PFCount(ctx, "merged_visitors")
	if err != nil {
		log.Printf("获取合并后HyperLogLog基数失败: %v", err)
		return
	}
	fmt.Printf("合并后独立访客数估算: %d\n", mergedCount)

	//6.清理数据
	redis.Client.String.Del(ctx, "unique_visitors", "unique_visitors_2", "merged_visitors")
	fmt.Println()
}
