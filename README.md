# Go Redis Demo - 重构后的模块化Redis客户端

这是一个经过重构的 Go Redis 客户端库，采用模块化设计，将不同的 Redis 数据类型操作分离到各自的子包中。

## 项目结构

```
redis/
├── client.go          # 统一的客户端入口
├── config.go          # Redis配置管理
├── init.go            # 客户端初始化
├── tests/             # 单元测试目录
│   ├── string_client_test.go
│   ├── hash_client_test.go
│   ├── list_client_test.go
│   ├── set_client_test.go
│   ├── zset_client_test.go
│   ├── geo_client_test.go
│   ├── bitmap_client_test.go
│   └── hll_client_test.go
├── string/            # 字符串操作
│   └── string.go
├── hash/              # 哈希操作
│   └── hash.go
├── list/              # 列表操作
│   └── list.go
├── set/               # 集合操作
│   └── set.go
├── zset/              # 有序集合操作
│   └── zset.go
├── geo/               # 地理位置操作
│   └── geo.go
├── bitmap/            # 位图操作
│   └── bitmap.go
└── hll/               # HyperLogLog操作
    └── hll.go
```

## 主要特性

1. **模块化设计**: 每种数据类型都有独立的子包，职责清晰
2. **统一入口**: 通过 `Client` 结构体提供统一的访问接口
3. **灵活配置**: 支持自定义配置和默认配置
4. **上下文支持**: 所有操作都支持 `context.Context`
5. **类型安全**: 充分利用 Go 的类型系统

## 使用方法

### 1. 初始化客户端

```go
package main

import (
    "context"
    "log"
    "time"
    
    "go-redis-demo/redis"
)

func main() {
    // 使用默认配置初始化
    config := redis.DefaultConfig()
    redis.InitClient(config)
    defer redis.CloseClient()
    
    // 或者使用自定义配置
    config = redis.DefaultConfig()
    config.Addr = "localhost:6379"
    config.Password = "your-password"
    config.DB = 1
    
    redis.InitClient(config)
    defer redis.CloseClient()
}
```

### 2. 字符串操作

```go
ctx := context.Background()

// 设置字符串
err := redis.Client.String.Set(ctx, "name", "CodeBuddy", time.Hour)

// 获取字符串
name, err := redis.Client.String.Get(ctx, "name")

// 自增操作
count, err := redis.Client.String.Incr(ctx, "counter")
```

### 3. 哈希操作

```go
// 设置哈希字段
_, err := redis.Client.Hash.HSet(ctx, "user:1", "name", "张三", "age", 25)

// 获取单个字段
name, err := redis.Client.Hash.HGet(ctx, "user:1", "name")

// 获取所有字段
userInfo, err := redis.Client.Hash.HGetAll(ctx, "user:1")
```

### 4. 列表操作

```go
// 左侧推入元素
_, err := redis.Client.List.LPush(ctx, "tasks", "任务1", "任务2")

// 获取列表范围
tasks, err := redis.Client.List.LRange(ctx, "tasks", 0, -1)

// 弹出元素
task, err := redis.Client.List.LPop(ctx, "tasks")
```

### 5. 集合操作

```go
// 添加元素到集合
_, err := redis.Client.Set.SAdd(ctx, "tags", "Go", "Redis", "数据库")

// 获取集合所有成员
members, err := redis.Client.Set.SMembers(ctx, "tags")

// 判断元素是否存在
exists, err := redis.Client.Set.SIsMember(ctx, "tags", "Go")
```

### 6. 有序集合操作

```go
import "github.com/redis/go-redis/v9"

// 添加元素到有序集合
_, err := redis.Client.ZSet.ZAdd(ctx, "scores", 
    redis.Z{Score: 95.5, Member: "张三"},
    redis.Z{Score: 87.0, Member: "李四"},
)

// 获取排名（从高到低）
topScores, err := redis.Client.ZSet.ZRevRangeWithScores(ctx, "scores", 0, 2)
```

### 7. 地理位置操作

```go
// 添加地理位置
_, err := redis.Client.Geo.GeoAdd(ctx, "cities", 116.397128, 39.916527, "北京")

// 计算距离
distance, err := redis.Client.Geo.GeoDist(ctx, "cities", "北京", "上海", "km")

// 半径搜索
locations, err := redis.Client.Geo.GeoRadius(ctx, "cities", 116.397128, 39.916527, 1000, "km", true, true, true, 10)
```

### 8. 位图操作

```go
import "github.com/redis/go-redis/v9"

// 设置位
_, err := redis.Client.Bitmap.SetBit(ctx, "user_sign", 0, 1)

// 获取位
bit, err := redis.Client.Bitmap.GetBit(ctx, "user_sign", 0)

// 统计位数
count, err := redis.Client.Bitmap.BitCount(ctx, "user_sign", &redis.BitCount{Start: 0, End: -1})
```

### 9. HyperLogLog操作

```go
// 添加元素
_, err := redis.Client.HLL.PFAdd(ctx, "unique_visitors", "user1", "user2", "user3")

// 获取基数估算
count, err := redis.Client.HLL.PFCount(ctx, "unique_visitors")

// 合并HyperLogLog
err = redis.Client.HLL.PFMerge(ctx, "merged_visitors", "visitors1", "visitors2")
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

## 测试

项目包含完整的单元测试，所有测试文件位于 `redis/tests/` 目录中：

```bash
# 运行所有测试
go test ./redis/tests

# 运行特定测试
go test ./redis/tests -run Test_stringClient

# 运行测试并显示详细输出
go test -v ./redis/tests
```

测试覆盖了所有Redis数据类型的操作：
- 字符串操作测试 (`string_client_test.go`)
- 哈希操作测试 (`hash_client_test.go`)
- 列表操作测试 (`list_client_test.go`)
- 集合操作测试 (`set_client_test.go`)
- 有序集合操作测试 (`zset_client_test.go`)
- 地理位置操作测试 (`geo_client_test.go`)
- 位图操作测试 (`bitmap_client_test.go`)
- HyperLogLog操作测试 (`hll_client_test.go`)

## 迁移指南

如果您之前使用的是旧版本的客户端，迁移步骤如下：

### 旧版本调用方式:
```go
redis.InitRedis()
redis.StringClient.Set("key", "value", time.Hour)
```

### 新版本调用方式:
```go
config := redis.DefaultConfig()
redis.InitClient(config)
defer redis.CloseClient()

ctx := context.Background()
redis.Client.String.Set(ctx, "key", "value", time.Hour)
```

主要变化：
1. 使用 `redis.InitClient()` 和 `redis.CloseClient()` 进行初始化和清理
2. 所有操作都需要传入 `context.Context`
3. API 调用方式从 `redis.StringClient.Set` 改为 `redis.Client.String.Set`
4. 支持更多Redis数据类型（地理位置、位图、HyperLogLog等）
