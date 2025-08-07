// Package redis 提供了Redis客户端的统一封装
// @Author:冯铁城 [17615007230@163.com] 2025-08-07 10:00:00
package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	bitmappkg "go-redis-demo/redis/bitmap"
	geopkg "go-redis-demo/redis/geo"
	hashpkg "go-redis-demo/redis/hash"
	hllpkg "go-redis-demo/redis/hll"
	listpkg "go-redis-demo/redis/list"
	setpkg "go-redis-demo/redis/set"
	stringpkg "go-redis-demo/redis/string"
	zsetpkg "go-redis-demo/redis/zset"
)

// client 是统一的Redis客户端，提供所有数据类型操作的入口
type client struct {
	rdb    *redis.Client     // 底层go-redis客户端
	String *stringpkg.Client // 字符串操作客户端
	Hash   *hashpkg.Client   // 哈希操作客户端
	List   *listpkg.Client   // 列表操作客户端
	Set    *setpkg.Client    // 集合操作客户端
	ZSet   *zsetpkg.Client   // 有序集合操作客户端
	Geo    *geopkg.Client    // 地理位置操作客户端
	Bitmap *bitmappkg.Client // 位图操作客户端
	HLL    *hllpkg.Client    // HyperLogLog操作客户端
}

// NewClient 创建一个新的Redis客户端实例
func newClient(config *Config) (*client, error) {

	//1.创建底层go-redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		MaxRetries:   config.MaxRetries,
	})

	//2.测试连接
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	//3.创建统一客户端，组装各个数据类型的操作客户端
	redisClient := &client{
		rdb:    rdb,
		String: stringpkg.New(rdb),
		Hash:   hashpkg.New(rdb),
		List:   listpkg.New(rdb),
		Set:    setpkg.New(rdb),
		ZSet:   zsetpkg.New(rdb),
		Geo:    geopkg.New(rdb),
		Bitmap: bitmappkg.New(rdb),
		HLL:    hllpkg.New(rdb),
	}

	//4.返回
	return redisClient, nil
}

// Close 关闭Redis连接
func (c *client) Close() error {
	if c.rdb != nil {
		return c.rdb.Close()
	}
	return nil
}

// Ping 测试Redis连接是否正常
func (c *client) Ping(ctx context.Context) error {
	return c.rdb.Ping(ctx).Err()
}

// GetRawClient 获取底层的go-redis客户端（高级用法）
func (c *client) GetRawClient() *redis.Client {
	return c.rdb
}
