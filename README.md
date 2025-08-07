# Go Redis Demo - 重构后的模块化Redis客户端

这是一个经过重构的 Go Redis 客户端库，采用模块化设计，将不同的 Redis 数据类型操作分离到各自的子包中。

## 项目结构

```
redis/
├── client.go          # 统一的客户端入口
├── config.go          # Redis配置管理
├── string/            # 字符串操作
│   └── string.go
├── hash/              # 哈希操作
│   └── hash.go
├── list/              # 列表操作
│   └── list.go
├── set/               # 集合操作
│   └── set.go
└── zset/              # 有序集合操作
    └── zset.go
```

## 主要特性

1. **模块化设计**: 每种数据类型都有独立的子包，职责清晰
2. **统一入口**: 通过 `Client` 结构体提供统一的访问接口
3. **灵活配置**: 支持自定义配置和默认配置
4. **上下文支持**: 所有操作都支持 `context.Context`
5. **类型安全**: 充分利用 Go 的类型系统

## 使用方法

### 1. 创建客户端

```go
package main

import (
    "context"
    "log"
    "time"
    
    "go-redis-demo/redis"
)

func main() {
    // 使用默认配置
    client, err := redis.NewClientWithDefault()
    if err != nil {
        log.Fatalf("创建Redis客户端失败: %v", err)
    }
    defer client.Close()
    
    // 或者使用自定义配置
    config := redis.DefaultConfig()
    config.Addr = "localhost:6379"
    config.Password = "your-password"
    config.DB = 1
    
    client, err = redis.NewClient(config)
    if err != nil {
        log.Fatalf("创建Redis客户端失败: %v", err)
    }
    defer client.Close()
}
```

### 2. 字符串操作

```go
ctx := context.Background()

// 设置字符串
err := client.String.Set(ctx, "name", "CodeBuddy", time.Hour)

// 获取字符串
name, err := client.String.Get(ctx, "name")

// 自增操作
count, err := client.String.Incr(ctx, "counter")
```

### 3. 哈希操作

```go
// 设置哈希字段
_, err := client.Hash.HSet(ctx, "user:1", "name", "张三", "age", 25)

// 获取单个字段
name, err := client.Hash.HGet(ctx, "user:1", "name")

// 获取所有字段
userInfo, err := client.Hash.HGetAll(ctx, "user:1")
```

### 4. 列表操作

```go
// 左侧推入元素
_, err := client.List.LPush(ctx, "tasks", "任务1", "任务2")

// 获取列表范围
tasks, err := client.List.LRange(ctx, "tasks", 0, -1)

// 弹出元素
task, err := client.List.LPop(ctx, "tasks")
```

### 5. 集合操作

```go
// 添加元素到集合
_, err := client.Set.SAdd(ctx, "tags", "Go", "Redis", "数据库")

// 获取集合所有成员
members, err := client.Set.SMembers(ctx, "tags")

// 判断元素是否存在
exists, err := client.Set.SIsMember(ctx, "tags", "Go")
```

### 6. 有序集合操作

```go
import "github.com/redis/go-redis/v9"

// 添加元素到有序集合
_, err := client.ZSet.ZAdd(ctx, "scores", 
    redis.Z{Score: 95.5, Member: "张三"},
    redis.Z{Score: 87.0, Member: "李四"},
)

// 获取排名（从高到低）
topScores, err := client.ZSet.ZRevRangeWithScores(ctx, "scores", 0, 2)
```

## 配置选项

```go
type Config struct {
    Addr         string        // Redis服务器地址
    Password     string        // 密码
    DB           int           // 数据库编号
    PoolSize     int           // 连接池大小
    MinIdleConns int           // 最小空闲连接数
    DialTimeout  time.Duration // 连接超时时间
    ReadTimeout  time.Duration // 读取超时时间
    WriteTimeout time.Duration // 写入超时时间
    MaxRetries   int           // 最大重试次数
}
```

## 优势

1. **清晰的代码组织**: 每种数据类型的操作都在独立的包中，便于维护
2. **易于扩展**: 添加新的数据类型操作只需创建新的子包
3. **统一的API风格**: 所有操作都遵循相同的命名和参数约定
4. **更好的测试支持**: 可以针对每个子包编写独立的测试
5. **减少依赖**: 使用者可以根据需要导入特定的功能模块

## 迁移指南

如果您之前使用的是旧版本的客户端，迁移步骤如下：

### 旧版本调用方式:
```go
redis.InitRedis()
redis.StringClient.Set("key", "value", time.Hour)
```

### 新版本调用方式:
```go
client, _ := redis.NewClientWithDefault()
defer client.Close()
client.String.Set(ctx, "key", "value", time.Hour)
```

主要变化：
1. 需要显式创建和关闭客户端实例
2. 所有操作都需要传入 `context.Context`
3. API 调用方式从 `redis.StringClient.Set` 改为 `client.String.Set`