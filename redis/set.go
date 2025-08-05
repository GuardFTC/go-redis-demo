// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:29:24
package redis

// SetClient redis集合操作
var SetClient = new(setClient)

// setClient redis集合操作
type setClient struct {
}

// SAdd 添加若干指定元素member到key集合中，并返回成功添加元素个数
func (s *setClient) SAdd(key string, members ...interface{}) (int64, error) {
	return rdb.SAdd(ctx, key, members...).Result()
}

// SPop 随机移除并返回集合key中若干随机元素
func (s *setClient) SPop(key string, count ...int64) ([]string, error) {
	if len(count) > 0 {
		return rdb.SPopN(ctx, key, count[0]).Result()
	}
	result, err := rdb.SPop(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []string{result}, nil
}

// SRem 在集合key中移除指定元素，并返回成功移除元素个数
func (s *setClient) SRem(key string, members ...interface{}) (int64, error) {
	return rdb.SRem(ctx, key, members...).Result()
}

// SCard 返回指定集合key中的元素数
func (s *setClient) SCard(key string) (int64, error) {
	return rdb.SCard(ctx, key).Result()
}

// SIsMember 返回集合key中是否存在指定元素member
func (s *setClient) SIsMember(key string, member interface{}) (bool, error) {
	return rdb.SIsMember(ctx, key, member).Result()
}

// SMembers 返回集合key的所有元素
func (s *setClient) SMembers(key string) ([]string, error) {
	return rdb.SMembers(ctx, key).Result()
}

// SRandMember 随机返回集合key中的一个元素，或随机返回集合key中的count的元素
func (s *setClient) SRandMember(key string, count ...int64) ([]string, error) {
	if len(count) > 0 {
		return rdb.SRandMemberN(ctx, key, count[0]).Result()
	}
	result, err := rdb.SRandMember(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []string{result}, nil
}

// SMove 将指定元素member从集合source中移动到集合destination中
func (s *setClient) SMove(source, destination string, member interface{}) (bool, error) {
	return rdb.SMove(ctx, source, destination, member).Result()
}

// SInter 返回所有指定集合中元素的交集
func (s *setClient) SInter(keys ...string) ([]string, error) {
	return rdb.SInter(ctx, keys...).Result()
}

// SInterStore 返回所有指定集合中元素的交集，并将结果保存在集合destination中
func (s *setClient) SInterStore(destination string, keys ...string) (int64, error) {
	return rdb.SInterStore(ctx, destination, keys...).Result()
}

// SUnion 返回所有指定集合中元素的并集
func (s *setClient) SUnion(keys ...string) ([]string, error) {
	return rdb.SUnion(ctx, keys...).Result()
}

// SUnionStore 返回所有指定集合中元素的并集，并将结果保存在集合destination中
func (s *setClient) SUnionStore(destination string, keys ...string) (int64, error) {
	return rdb.SUnionStore(ctx, destination, keys...).Result()
}

// SDiff 返回一个集合与其余指定集合的差集
func (s *setClient) SDiff(keys ...string) ([]string, error) {
	return rdb.SDiff(ctx, keys...).Result()
}

// SDiffStore 返回一个集合与其余指定集合的差集，并将结果保存在集合destination中
func (s *setClient) SDiffStore(destination string, keys ...string) (int64, error) {
	return rdb.SDiffStore(ctx, destination, keys...).Result()
}
