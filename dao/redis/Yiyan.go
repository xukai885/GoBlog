package redis

import (
	"GoBlog/modules"
	"math/rand"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// GetYiYanWriteRedis 获取一言写入到redis
func GetYiYanWriteRedis(yiyan []modules.YiYan) {
	// 存储名言和名言出处数据到哈希表
	for i, yiyan := range yiyan {
		keyi := strconv.Itoa(i)
		// 1.写入到redis
		rdb.HSet(keyi, "hitokoto", yiyan.Hitokoto)
		rdb.HSet(keyi, "from_source", yiyan.FromSource)
		rdb.HSet(keyi, "from_who", yiyan.FromWho)
		zap.L().Info("写入redis成功", zap.String("hitokoto", yiyan.Hitokoto))
	}
}

// GetYiYan 从redis获取一言
func GetYiYan() (yiyan modules.YiYan) {
	num := strconv.Itoa(GetRand(int(rdb.DBSize().Val())))
	// 获取redis数据
	yiyan = modules.YiYan{
		Hitokoto:   rdb.HGet(num, "hitokoto").Val(),
		FromSource: rdb.HGet(num, "from_source").Val(),
		FromWho:    rdb.HGet(num, "from_who").Val(),
	}
	return
}

func GetRand(max int) (num int) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 生成 0 到 max之间的随机数
	num = rand.Intn(max)
	return
}
