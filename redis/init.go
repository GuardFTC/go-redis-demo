// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:08:55
package redis

import (
	"log"
)

// Client redis客户端
var Client *client

// InitClient 初始化redis客户端
func InitClient(config *Config) {

	//1.初始化Redis客户端
	c, err := newClient(config)
	if err != nil {
		log.Fatalf("redis connection error: %v", err)
	} else {
		log.Println("redis connection success")
	}

	//2.全局客户端赋值
	Client = c
}

// CloseClient 关闭redis客户端
func CloseClient() {
	if err := Client.Close(); err != nil {
		log.Fatalf("redis connection closed error: %v", err)
	} else {
		log.Println("redis connection closed success")
	}
}
