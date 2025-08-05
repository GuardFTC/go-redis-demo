// @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:08:55
package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

// rdb redis客户端
var rdb *redis.Client

// ctx 空白上下文
var ctx = context.Background()

// InitRedis 初始化redis
func InitRedis() {

	//1.初始化Redis客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//2.测试链接是否建立成功
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("redis connection error: %v", err)
	} else {
		log.Printf("redis connection success")
	}
}
