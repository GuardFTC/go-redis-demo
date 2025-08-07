package main

import "go-redis-demo/redis"

func main() {

	//1.初始化Redis客户端
	redis.InitClient(redis.DefaultConfig())
	defer redis.CloseClient()
}
