// @Author:冯铁城 [17615007230@163.com] 2025-08-05 19:45:35
package redis

import (
	"github.com/redis/go-redis/v9"
	"testing"
)

func Test_geoClient(t *testing.T) {

	//1.初始化链接
	InitRedis()

	//2.运行测试
	t.Run("redis geo客户端测试", func(t *testing.T) {
		g := GeoClient

		//1.添加单个地理位置
		count, err := g.GeoAdd("geo_key", 116.397128, 39.916527, "北京")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//2.再次添加相同位置（覆盖）
		count, err = g.GeoAdd("geo_key", 116.397128, 39.916527, "北京")
		if count != 0 || err != nil {
			t.Error(err)
		}

		//3.添加更多地理位置
		count, err = g.GeoAdd("geo_key", 121.473701, 31.230416, "上海")
		if count != 1 || err != nil {
			t.Error(err)
		}

		//4.批量添加地理位置
		locations := []*redis.GeoLocation{
			{Longitude: 113.264434, Latitude: 23.129162, Name: "广州"},
			{Longitude: 114.085947, Latitude: 22.547, Name: "深圳"},
			{Longitude: 104.065735, Latitude: 30.659462, Name: "成都"},
		}
		count, err = g.GeoBatchAdd("geo_key", locations...)
		if count != 3 || err != nil {
			t.Error(err)
		}

		//5.获取地理位置的经纬度
		positions, err := g.GeoPos("geo_key", "北京", "上海", "广州", "深圳", "成都", "不存在的城市")
		if err != nil {
			t.Error(err)
		}
		if len(positions) != 6 || positions[5] != nil {
			t.Error("GeoPos结果不符合预期")
		}
		// 验证北京的经纬度
		if positions[0] == nil || !floatEquals(positions[0].Longitude, 116.397128, 0.001) || !floatEquals(positions[0].Latitude, 39.916527, 0.001) {
			t.Error("北京的经纬度不符合预期")
		}

		//6.计算两个位置之间的距离（单位：千米）
		distance, err := g.GeoDist("geo_key", "北京", "上海", "km")
		if err != nil {
			t.Error(err)
		}
		// 北京到上海的距离约为1068公里，允许有误差
		if distance < 1000 || distance > 1100 {
			t.Errorf("北京到上海的距离不符合预期: %f", distance)
		}

		//7.测试无效的距离单位
		_, err = g.GeoDist("geo_key", "北京", "上海", "invalid")
		if err == nil || err.Error() != "无效的距离单位，必须是m、km、mi或ft之一" {
			t.Error("无效距离单位未返回预期错误")
		}

		//8.获取地理位置的Geohash表示
		hashes, err := g.GeoHash("geo_key", "北京", "上海", "不存在的城市")
		if err != nil {
			t.Error(err)
		}
		if len(hashes) != 3 || hashes[0] == "" || hashes[1] == "" || hashes[2] != "" {
			t.Error("GeoHash结果不符合预期")
		}

		//9.以坐标为中心进行半径搜索
		radius, err := g.GeoRadius("geo_key", 116.397128, 39.916527, 1500, "km", true, true, true, 0)
		if err != nil {
			t.Error(err)
		}
		// 应该返回所有添加的城市
		if len(radius) != 2 {
			t.Errorf("GeoRadius结果不符合预期，应返回5个城市，实际返回%d个", len(radius))
		}

		//10.测试无效的距离单位进行半径搜索
		_, err = g.GeoRadius("geo_key", 116.397128, 39.916527, 1500, "invalid", true, true, true, 0)
		if err == nil || err.Error() != "无效的距离单位，必须是m、km、mi或ft之一" {
			t.Error("无效距离单位未返回预期错误")
		}

		//11.以成员为中心进行半径搜索
		radius, err = g.GeoRadiusByMember("geo_key", "北京", 1500, "km", true, true, true, 0)
		if err != nil {
			t.Error(err)
		}
		// 应该返回所有添加的城市
		if len(radius) != 2 {
			t.Errorf("GeoRadiusByMember结果不符合预期，应返回5个城市，实际返回%d个", len(radius))
		}

		//12.测试无效的距离单位进行成员半径搜索
		_, err = g.GeoRadiusByMember("geo_key", "北京", 1500, "invalid", true, true, true, 0)
		if err == nil || err.Error() != "无效的距离单位，必须是m、km、mi或ft之一" {
			t.Error("无效距离单位未返回预期错误")
		}

		//13.限制返回数量的半径搜索
		radius, err = g.GeoRadius("geo_key", 116.397128, 39.916527, 1500, "km", true, true, true, 2)
		if err != nil {
			t.Error(err)
		}
		// 应该只返回2个城市
		if len(radius) != 2 {
			t.Errorf("限制数量的GeoRadius结果不符合预期，应返回2个城市，实际返回%d个", len(locations))
		}

		//14.测试不存在的成员进行半径搜索
		radius, err = g.GeoRadiusByMember("geo_key", "不存在的城市", 1500, "km", true, true, true, 0)
		if err == nil || len(radius) != 0 {
			t.Error("不存在的成员进行半径搜索应返回错误或空结果")
		}

		//15.清理测试数据
		_, err = rdb.Del(ctx, "geo_key").Result()
		if err != nil {
			t.Error("清理数据失败")
		}

		//16.测试空geo的操作
		testEmptyGeoOperations(t, g, "empty_geo")
	})
}

// 测试空geo的各种操作
func testEmptyGeoOperations(t *testing.T, g *geoClient, key string) {

	//1.获取不存在的位置的经纬度
	positions, err := g.GeoPos(key, "nonexistent")
	if err != nil || len(positions) != 1 || positions[0] != nil {
		t.Error("空geo的GeoPos结果不符合预期")
	}

	//2.计算不存在的位置之间的距离
	_, err = g.GeoDist(key, "pos1", "pos2", "km")
	if err == nil || err.Error() != "redis: nil" {
		t.Error("空geo的GeoDist结果不符合预期")
	}

	//3.获取不存在的位置的Geohash
	hashes, err := g.GeoHash(key, "nonexistent")
	if err != nil || len(hashes) != 1 || hashes[0] != "" {
		t.Error("空geo的GeoHash结果不符合预期")
	}

	//4.以坐标为中心进行半径搜索
	locations, err := g.GeoRadius(key, 0, 0, 100, "km", true, true, true, 0)
	if err != nil || len(locations) != 0 {
		t.Error("空geo的GeoRadius结果不符合预期")
	}

	//5.以成员为中心进行半径搜索
	locations, err = g.GeoRadiusByMember(key, "nonexistent", 100, "km", true, true, true, 0)
	if err != nil || len(locations) != 0 {
		t.Error("空geo的GeoRadiusByMember结果不符合预期")
	}
}

// 浮点数比较，考虑误差
func floatEquals(a, b, epsilon float64) bool {
	return (a-b) < epsilon && (b-a) < epsilon
}
