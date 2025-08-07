// Package redis 提供了Redis客户端的核心功能，包括配置和初始化。
// @Author:冯铁城 [17615007230@163.com] 2025-08-07 10:00:00
package redis

import "time"

// Config 定义了Redis客户端的所有配置选项。
type Config struct {
	Addr         string        // Redis服务器地址，格式为 "host:port"
	Password     string        // 密码
	DB           int           // 使用的数据库编号
	PoolSize     int           // 连接池大小
	MinIdleConns int           // 最小空闲连接数
	DialTimeout  time.Duration // 连接超时时间
	ReadTimeout  time.Duration // 读取超时时间
	WriteTimeout time.Duration // 写入超时时间
	MaxRetries   int           // 最大重试次数
}

// DefaultConfig 返回一个包含推荐默认值的配置实例。
// 这些默认值适用于大多数本地开发环境。
func DefaultConfig() *Config {
	return &Config{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     100,             // 默认连接池大小
		MinIdleConns: 10,              // 默认最小空闲连接数
		DialTimeout:  5 * time.Second, // 5秒连接超时
		ReadTimeout:  3 * time.Second, // 3秒读取超时
		WriteTimeout: 3 * time.Second, // 3秒写入超时
		MaxRetries:   3,               // 失败时重试3次
	}
}
