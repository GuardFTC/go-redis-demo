// @Author:冯铁城 [17615007230@163.com] 2025-08-07 11:30:00
package redis_test

import (
	"context"
	"testing"
	"time"

	"go-redis-demo/redis"
	stringpkg "go-redis-demo/redis/string"
)

func Test_stringClient(t *testing.T) {

	//1.初始化链接
	config := redis.DefaultConfig()
	redis.InitClient(config)
	defer redis.CloseClient()

	//2.运行测试
	t.Run("redis string客户端测试", func(t *testing.T) {
		s := redis.Client.String
		ctx := context.Background()

		//1.设置 key,使用默认过期时间
		err := s.SetWithDefaultExpire(ctx, "key", "value")
		if err != nil {
			t.Error(err)
		}

		//2.测试
		testTTLAndGetAndDel(t, s, ctx, func(ttl time.Duration) bool {
			return ttl != -1
		}, "key", "value")

		//3.设置 key 不过期
		err = s.Set(ctx, "key", "value", 0)
		if err != nil {
			t.Error(err)
		}

		//4.测试
		testTTLAndGetAndDel(t, s, ctx, func(ttl time.Duration) bool {
			return ttl == -1
		}, "key", "value")

		//5.设置 key,使用默认过期时间
		ok, err := s.SetNXWithDefaultExpire(ctx, "key", "value")
		if !ok || err != nil {
			t.Error(err)
		}

		//6.再次设置key
		ok, err = s.SetNXWithDefaultExpire(ctx, "key", "value")
		if ok || err != nil {
			t.Error(err)
		}

		//7.测试
		testTTLAndGetAndDel(t, s, ctx, func(ttl time.Duration) bool {
			return ttl != -1
		}, "key", "value")

		//8.设置 key,不过期
		ok, err = s.SetNX(ctx, "key", "value", 0)
		if !ok || err != nil {
			t.Error(err)
		}

		//9.再次设置key
		ok, err = s.SetNX(ctx, "key", "value", 0)
		if ok || err != nil {
			t.Error(err)
		}

		//10.测试
		testTTLAndGetAndDel(t, s, ctx, func(ttl time.Duration) bool {
			return ttl == -1
		}, "key", "value")

		//11.设置多个key
		err = s.MSet(ctx, "key1", "value1", "key2", "value2")
		if err != nil {
			t.Error(err)
		}

		//12.测试
		testTTLAndGetAndDel(t, s, ctx, func(ttl time.Duration) bool {
			return ttl == -1
		}, "key1", "value1")
		testTTLAndGetAndDel(t, s, ctx, func(ttl time.Duration) bool {
			return ttl == -1
		}, "key2", "value2")

		//13.设置多个key
		ok, err = s.MSetNX(ctx, "key1", "value1", "key2", "value2")
		if !ok || err != nil {
			t.Error(err)
		}

		//14.再次设置多个Key
		ok, err = s.MSetNX(ctx, "key1", "value1", "key2", "value2")
		if ok || err != nil {
			t.Error(err)
		}

		//15.获取keys
		keys, err := s.Keys(ctx, "*")
		if keys == nil || len(keys) != 2 || err != nil {
			t.Error(err)
		}

		//16.设置过期时间
		err = s.Expire(ctx, "key1", 60*time.Second)
		err = s.Expire(ctx, "key2", 60*time.Second)

		//17.测试
		testTTLAndGetAndDel(t, s, ctx, func(ttl time.Duration) bool {
			return ttl != -1
		}, "key1", "value1")
		testTTLAndGetAndDel(t, s, ctx, func(ttl time.Duration) bool {
			return ttl != -1
		}, "key2", "value2")

		//18.写入数据
		err = s.Set(ctx, "key", "0", 60*time.Second)
		if err != nil {
			t.Error(err)
		}

		//19.递增+1
		incr, err := s.Incr(ctx, "key")
		if incr != 1 || err != nil {
			t.Error(err)
		}

		//20.递增+10
		by, err := s.IncrBy(ctx, "key", 10)
		if by != 11 || err != nil {
			t.Error(err)
		}

		//21.递减-1
		decr, err := s.Decr(ctx, "key")
		if decr != 10 || err != nil {
			t.Error(err)
		}

		//22.递减-10
		decr, err = s.DecrBy(ctx, "key", 10)
		if decr != 0 || err != nil {
			t.Error(err)
		}

		//23.删除key
		_, err = s.Del(ctx, "key")
		if err != nil {
			t.Error(err)
		}

		//24.判定key不存在
		exists, err := s.Exists(ctx, "key")
		if exists != 0 || err != nil {
			t.Error(err)
		}
	})
}

// 测试TTL、Get、Del
func testTTLAndGetAndDel(t *testing.T, s *stringpkg.Client, ctx context.Context, expectTTL func(time.Duration) bool, expKey string, expValue string) {

	//1.获取过期时间
	ttl, err := s.TTL(ctx, expKey)
	if err != nil {
		t.Error(err)
	}
	if !expectTTL(ttl) {
		t.Errorf("ttl 不符合预期: %d", ttl)
	}

	//2.判定key是否存在
	exists, err := s.Exists(ctx, expKey)
	if exists != 1 || err != nil {
		t.Error(err)
	}

	//3.查询 key
	value, err := s.Get(ctx, expKey)
	if value != expValue || err != nil {
		t.Error(err)
	}

	//4.删除 key
	_, err = s.Del(ctx, expKey)
	if err != nil {
		t.Error(err)
	}

	//5.再次查询key
	value, err = s.Get(ctx, expKey)
	if value != "" || err == nil || err.Error() != "redis: nil" {
		t.Error(err)
	}

	//6.再次判定key是否存在
	exists, err = s.Exists(ctx, expKey)
	if exists != 0 || err != nil {
		t.Error(err)
	}
}
