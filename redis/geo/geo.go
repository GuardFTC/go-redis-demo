// Package geo 提供Redis地理位置操作的封装
// @Author:冯铁城 [17615007230@163.com] 2025-08-07 11:00:00
package geo

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

// Client Redis地理位置操作客户端
type Client struct {
	rdb *redis.Client
}

// New 创建地理位置操作客户端
func New(rdb *redis.Client) *Client {
	return &Client{rdb: rdb}
}

// GeoAdd 将指定的地理空间位置（纬度、经度、名称）添加到指定的key中
// 参数:
//   - ctx: 上下文
//   - key: 键名
//   - longitude: 经度
//   - latitude: 纬度
//   - member: 位置名称
//
// 返回:
//   - 新添加的位置数量
//   - 错误信息
func (c *Client) GeoAdd(ctx context.Context, key string, longitude, latitude float64, member string) (int64, error) {
	return c.rdb.GeoAdd(ctx, key, &redis.GeoLocation{
		Longitude: longitude,
		Latitude:  latitude,
		Name:      member,
	}).Result()
}

// GeoBatchAdd 批量添加地理空间位置
// 参数:
//   - ctx: 上下文
//   - key: 键名
//   - locations: 位置信息数组，每个元素包含经度、纬度和名称
//
// 返回:
//   - 新添加的位置数量
//   - 错误信息
func (c *Client) GeoBatchAdd(ctx context.Context, key string, locations ...*redis.GeoLocation) (int64, error) {
	return c.rdb.GeoAdd(ctx, key, locations...).Result()
}

// GeoPos 从key里返回所有给定位置元素的位置（经度和纬度）
// 参数:
//   - ctx: 上下文
//   - key: 键名
//   - members: 位置名称列表
//
// 返回:
//   - 位置的经纬度列表，如果位置不存在则对应元素为nil
//   - 错误信息
func (c *Client) GeoPos(ctx context.Context, key string, members ...string) ([]*redis.GeoPos, error) {
	return c.rdb.GeoPos(ctx, key, members...).Result()
}

// GeoDist 返回两个给定位置之间的距离
// 参数:
//   - ctx: 上下文
//   - key: 键名
//   - member1: 第一个位置名称
//   - member2: 第二个位置名称
//   - unit: 距离单位，可选值："m"(米)、"km"(千米)、"mi"(英里)、"ft"(英尺)
//
// 返回:
//   - 两个位置之间的距离
//   - 错误信息
func (c *Client) GeoDist(ctx context.Context, key, member1, member2, unit string) (float64, error) {

	// 1.验证单位是否有效
	if unit != "m" && unit != "km" && unit != "mi" && unit != "ft" {
		return 0, errors.New("无效的距离单位，必须是m、km、mi或ft之一")
	}

	// 2.返回两个给定位置之间的距离
	return c.rdb.GeoDist(ctx, key, member1, member2, unit).Result()
}

// GeoHash 返回一个或多个位置元素的Geohash表示
// 参数:
//   - ctx: 上下文
//   - key: 键名
//   - members: 位置名称列表
//
// 返回:
//   - 位置的Geohash表示列表
//   - 错误信息
func (c *Client) GeoHash(ctx context.Context, key string, members ...string) ([]string, error) {
	return c.rdb.GeoHash(ctx, key, members...).Result()
}

// GeoRadius 以给定的经纬度为中心，返回键包含的位置元素当中，与中心的距离不超过给定最大距离的所有位置元素
// 参数:
//   - ctx: 上下文
//   - key: 键名
//   - longitude: 中心点经度
//   - latitude: 中心点纬度
//   - radius: 距离
//   - unit: 距离单位，可选值："m"(米)、"km"(千米)、"mi"(英里)、"ft"(英尺)
//   - withCoord: 是否返回位置的经纬度
//   - withDist: 是否返回位置与中心点的距离
//   - withHash: 是否返回位置的Geohash表示
//   - count: 返回的位置数量限制，0表示不限制
//
// 返回:
//   - 位置信息列表
//   - 错误信息
func (c *Client) GeoRadius(ctx context.Context, key string, longitude, latitude float64, radius float64, unit string, withCoord, withDist, withHash bool, count int) ([]redis.GeoLocation, error) {

	// 1.验证单位是否有效
	if unit != "m" && unit != "km" && unit != "mi" && unit != "ft" {
		return nil, errors.New("无效的距离单位，必须是m、km、mi或ft之一")
	}

	// 2.以给定的经纬度为中心，返回键包含的位置元素当中，与中心的距离不超过给定最大距离的所有位置元素
	return c.rdb.GeoRadius(ctx, key, longitude, latitude, &redis.GeoRadiusQuery{
		Radius:      radius,
		Unit:        unit,
		WithCoord:   withCoord,
		WithDist:    withDist,
		WithGeoHash: withHash,
		Count:       count,
	}).Result()
}

// GeoRadiusByMember 以给定的位置元素为中心，返回键包含的位置元素当中，与中心的距离不超过给定最大距离的所有位置元素
// 参数:
//   - ctx: 上下文
//   - key: 键名
//   - member: 中心点位置名称
//   - radius: 距离
//   - unit: 距离单位，可选值："m"(米)、"km"(千米)、"mi"(英里)、"ft"(英尺)
//   - withCoord: 是否返回位置的经纬度
//   - withDist: 是否返回位置与中心点的距离
//   - withHash: 是否返回位置的Geohash表示
//   - count: 返回的位置数量限制，0表示不限制
//
// 返回:
//   - 位置信息列表
//   - 错误信息
func (c *Client) GeoRadiusByMember(ctx context.Context, key, member string, radius float64, unit string, withCoord, withDist, withHash bool, count int) ([]redis.GeoLocation, error) {

	// 1.验证单位是否有效
	if unit != "m" && unit != "km" && unit != "mi" && unit != "ft" {
		return nil, errors.New("无效的距离单位，必须是m、km、mi或ft之一")
	}

	// 2.以给定的位置元素为中心，返回键包含的位置元素当中，与中心的距离不超过给定最大距离的所有位置元素
	return c.rdb.GeoRadiusByMember(ctx, key, member, &redis.GeoRadiusQuery{
		Radius:      radius,
		Unit:        unit,
		WithCoord:   withCoord,
		WithDist:    withDist,
		WithGeoHash: withHash,
		Count:       count,
	}).Result()
}

// GeoSearch 使用GEOSEARCH命令搜索地理位置
// 参数:
//   - ctx: 上下文
//   - key: 键名
//   - q: 搜索查询参数
//
// 返回:
//   - 位置名称列表
//   - 错误信息
func (c *Client) GeoSearch(ctx context.Context, key string, q *redis.GeoSearchQuery) ([]string, error) {
	return c.rdb.GeoSearch(ctx, key, q).Result()
}

// GeoSearchLocation 使用GEOSEARCH命令搜索地理位置，返回详细信息
// 参数:
//   - ctx: 上下文
//   - key: 键名
//   - q: 搜索查询参数
//
// 返回:
//   - 位置详细信息列表
//   - 错误信息
func (c *Client) GeoSearchLocation(ctx context.Context, key string, q *redis.GeoSearchLocationQuery) ([]redis.GeoLocation, error) {
	return c.rdb.GeoSearchLocation(ctx, key, q).Result()
}

// GeoSearchStore 使用GEOSEARCHSTORE命令搜索地理位置并存储结果
// 参数:
//   - ctx: 上下文
//   - key: 源键名
//   - store: 存储结果的键名
//   - q: 搜索存储查询参数
//
// 返回:
//   - 存储的位置数量
//   - 错误信息
func (c *Client) GeoSearchStore(ctx context.Context, key, store string, q *redis.GeoSearchStoreQuery) (int64, error) {
	return c.rdb.GeoSearchStore(ctx, key, store, q).Result()
}
