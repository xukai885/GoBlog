package logic

import (
	"GoBlog/dao/mysql"
	"GoBlog/dao/redis"

	"go.uber.org/zap"
)

// GetYiYanWriteRedis 获取一言写入到redis
func GetYiYanWriteRedis() {
	// 拿取mysql数据
	yiyan, err := mysql.GetYiYan()
	if err != nil {
		zap.L().Error("mysql读取yiyan失败", zap.Error(err))
		return
	}
	// redis写入
	redis.GetYiYanWriteRedis(yiyan)

}
