// @Author:冯铁城 [17615007230@163.com] 2025-08-05 19:30:35
package redis

import (
	"testing"
)

func Test_bitMapClient(t *testing.T) {

	//1.初始化链接
	InitRedis()

	//2.运行测试
	t.Run("redis bitmap客户端测试", func(t *testing.T) {
		b := BitMapClient

		//1.模拟用户签到场景 - 创建一个用户1月份的签到记录
		//1.1 用户在1月1日签到
		val, err := b.SetBit("user:sign:1:202501", 0, 1)
		if err != nil || val != 0 {
			t.Error(err)
		}

		//1.2 用户在1月3日签到
		val, err = b.SetBit("user:sign:1:202501", 2, 1)
		if err != nil || val != 0 {
			t.Error(err)
		}

		//1.3 用户在1月5日签到
		val, err = b.SetBit("user:sign:1:202501", 4, 1)
		if err != nil || val != 0 {
			t.Error(err)
		}

		//1.4 检查用户1月1日是否签到
		val, err = b.GetBit("user:sign:1:202501", 0)
		if err != nil || val != 1 {
			t.Error("GetBit结果不符合预期，用户1月1日应该已签到")
		}

		//1.5 检查用户1月2日是否签到
		val, err = b.GetBit("user:sign:1:202501", 1)
		if err != nil || val != 0 {
			t.Error("GetBit结果不符合预期，用户1月2日应该未签到")
		}

		//1.6 检查用户1月3日是否签到
		val, err = b.GetBit("user:sign:1:202501", 2)
		if err != nil || val != 1 {
			t.Error("GetBit结果不符合预期，用户1月3日应该已签到")
		}

		//1.7 统计用户1月份签到次数
		count, err := b.BitCount("user:sign:1:202501", 0, -1)
		if err != nil || count != 3 {
			t.Errorf("BitCount结果不符合预期，用户1月份应签到3次，实际为%d次", count)
		}

		//1.8 查找用户1月份第一次签到日期
		pos, err := b.BitPos("user:sign:1:202501", 1, 0, -1)
		if err != nil || pos != 0 {
			t.Errorf("BitPos结果不符合预期，用户1月份第一次签到应为1月1日，实际为1月%d日", pos+1)
		}

		//1.9 查找用户1月份第一次未签到日期
		pos, err = b.BitPos("user:sign:1:202501", 0, 0, -1)
		if err != nil || pos != 1 {
			t.Errorf("BitPos结果不符合预期，用户1月份第一次未签到应为1月2日，实际为1月%d日", pos+1)
		}

		//2.模拟权限管理场景 - 创建两个角色的权限位图
		//2.1 设置管理员角色权限(权限1:读取、权限2:写入、权限3:删除)
		_, err = b.SetBit("role:permissions:admin", 0, 1) // 读取权限
		if err != nil {
			t.Error(err)
		}
		_, err = b.SetBit("role:permissions:admin", 1, 1) // 写入权限
		if err != nil {
			t.Error(err)
		}
		_, err = b.SetBit("role:permissions:admin", 2, 1) // 删除权限
		if err != nil {
			t.Error(err)
		}

		//2.2 设置普通用户角色权限(只有读取权限)
		_, err = b.SetBit("role:permissions:user", 0, 1) // 读取权限
		if err != nil {
			t.Error(err)
		}

		//2.3 检查管理员是否有读取权限
		val, err = b.GetBit("role:permissions:admin", 0)
		if err != nil || val != 1 {
			t.Error("GetBit结果不符合预期，管理员应该有读取权限")
		}

		//2.4 检查管理员是否有删除权限
		val, err = b.GetBit("role:permissions:admin", 2)
		if err != nil || val != 1 {
			t.Error("GetBit结果不符合预期，管理员应该有删除权限")
		}

		//2.5 检查普通用户是否有读取权限
		val, err = b.GetBit("role:permissions:user", 0)
		if err != nil || val != 1 {
			t.Error("GetBit结果不符合预期，普通用户应该有读取权限")
		}

		//2.6 检查普通用户是否有写入权限
		val, err = b.GetBit("role:permissions:user", 1)
		if err != nil || val != 0 {
			t.Error("GetBit结果不符合预期，普通用户应该没有写入权限")
		}

		//2.7 计算管理员拥有的权限数量
		count, err = b.BitCount("role:permissions:admin", 0, -1)
		if err != nil || count != 3 {
			t.Errorf("BitCount结果不符合预期，管理员应有3个权限，实际为%d个", count)
		}

		//2.8 计算普通用户拥有的权限数量
		count, err = b.BitCount("role:permissions:user", 0, -1)
		if err != nil || count != 1 {
			t.Errorf("BitCount结果不符合预期，普通用户应有1个权限，实际为%d个", count)
		}

		//3.使用位操作合并权限
		//3.1 计算两个角色的权限并集(拥有任一角色的所有权限)
		len, err := b.BitOpOR("role:permissions:combined", "role:permissions:admin", "role:permissions:user")
		if err != nil || len < 1 {
			t.Error(err)
		}

		//3.2 检查合并后的权限
		val, err = b.GetBit("role:permissions:combined", 0)
		if err != nil || val != 1 {
			t.Error("BitOpOR结果不符合预期，合并后应该有读取权限")
		}

		val, err = b.GetBit("role:permissions:combined", 1)
		if err != nil || val != 1 {
			t.Error("BitOpOR结果不符合预期，合并后应该有写入权限")
		}

		val, err = b.GetBit("role:permissions:combined", 2)
		if err != nil || val != 1 {
			t.Error("BitOpOR结果不符合预期，合并后应该有删除权限")
		}

		//3.3 计算两个角色的权限交集(同时拥有的权限)
		len, err = b.BitOpAND("role:permissions:common", "role:permissions:admin", "role:permissions:user")
		if err != nil || len < 1 {
			t.Error(err)
		}

		//3.4 检查交集权限
		val, err = b.GetBit("role:permissions:common", 0)
		if err != nil || val != 1 {
			t.Error("BitOpAND结果不符合预期，共同权限应该包含读取权限")
		}

		val, err = b.GetBit("role:permissions:common", 1)
		if err != nil || val != 0 {
			t.Error("BitOpAND结果不符合预期，共同权限不应该包含写入权限")
		}

		//3.5 计算管理员独有的权限(管理员有但普通用户没有的权限)
		len, err = b.BitOpXOR("role:permissions:unique", "role:permissions:admin", "role:permissions:user")
		if err != nil || len < 1 {
			t.Error(err)
		}

		//3.6 检查独有权限
		val, err = b.GetBit("role:permissions:unique", 0)
		if err != nil || val != 0 {
			t.Error("BitOpXOR结果不符合预期，独有权限不应该包含读取权限")
		}

		val, err = b.GetBit("role:permissions:unique", 1)
		if err != nil || val != 1 {
			t.Error("BitOpXOR结果不符合预期，独有权限应该包含写入权限")
		}

		val, err = b.GetBit("role:permissions:unique", 2)
		if err != nil || val != 1 {
			t.Error("BitOpXOR结果不符合预期，独有权限应该包含删除权限")
		}

		//3.7 计算权限的补集(取反)
		len, err = b.BitOpNOT("role:permissions:not_user", "role:permissions:user")
		if err != nil || len < 1 {
			t.Error(err)
		}

		//3.8 检查补集权限
		val, err = b.GetBit("role:permissions:not_user", 0)
		if err != nil || val != 0 {
			t.Error("BitOpNOT结果不符合预期，普通用户权限的补集不应该包含读取权限")
		}

		val, err = b.GetBit("role:permissions:not_user", 1)
		if err != nil || val != 1 {
			t.Error("BitOpNOT结果不符合预期，普通用户权限的补集应该包含写入权限")
		}

		//4.测试边界情况
		//4.1 测试大偏移量(模拟稀疏数据)
		_, err = b.SetBit("sparse:bitmap", 10000, 1)
		if err != nil {
			t.Error(err)
		}

		val, err = b.GetBit("sparse:bitmap", 10000)
		if err != nil || val != 1 {
			t.Error("大偏移量GetBit结果不符合预期")
		}

		//4.2 测试不存在的key
		val, err = b.GetBit("nonexistent:bitmap", 0)
		if err != nil || val != 0 {
			t.Error("不存在的key GetBit结果不符合预期")
		}

		//5.模拟在线用户统计场景
		testOnlineUserStatistics(t, b)

		//6.清理测试数据
		cleanupKeys(t, []string{
			"user:sign:1:202501",
			"role:permissions:admin", "role:permissions:user",
			"role:permissions:combined", "role:permissions:common",
			"role:permissions:unique", "role:permissions:not_user",
			"sparse:bitmap", "online:users:20250105",
		})
	})
}

// 测试在线用户统计场景
func testOnlineUserStatistics(t *testing.T, b *bitMapClient) {

	//1.模拟记录一天内用户在线情况(用户ID作为偏移量)
	key := "online:users:20250105"

	//2.记录100个用户的在线情况(偶数ID的用户在线)
	for i := 0; i < 100; i += 2 {
		_, err := b.SetBit(key, int64(i), 1)
		if err != nil {
			t.Error("记录用户在线状态失败")
			return
		}
	}

	//3.统计在线用户数量
	onlineCount, err := b.BitCount(key, 0, -1)
	if err != nil {
		t.Error(err)
		return
	}

	//4.验证在线用户数量
	expectedOnline := 50 // 0-99的偶数共50个
	if onlineCount != int64(expectedOnline) {
		t.Errorf("在线用户统计结果不符合预期: expected=%d, actual=%d", expectedOnline, onlineCount)
	}

	//5.检查特定用户是否在线
	for i := 0; i < 10; i++ {
		userID := int64(i)
		val, err := b.GetBit(key, userID)
		if err != nil {
			t.Error(err)
			continue
		}

		expectedStatus := int64(0)
		if i%2 == 0 {
			expectedStatus = 1
		}

		if val != expectedStatus {
			t.Errorf("用户%d的在线状态不符合预期: expected=%d, actual=%d", userID, expectedStatus, val)
		}
	}

	//6.查找第一个在线用户
	firstOnline, err := b.BitPos(key, 1, 0, -1)
	if err != nil || firstOnline != 0 {
		t.Errorf("第一个在线用户ID不符合预期: expected=0, actual=%d", firstOnline)
	}
}
