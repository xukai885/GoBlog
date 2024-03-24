package controllers

import (
	"GoBlog/dao/mysql"
	"GoBlog/dao/redis"
	"GoBlog/settings"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// GetYiYan 获取一言
func GetYiYan(c *gin.Context) {
	// 判断是否写入一言到redis
	if settings.Conf.YiYanReadToRedis == true {
		yiyan := redis.GetYiYan()
		ResponseSuccess(c, yiyan)
	} else {
		// 从mysql拿数据
		yiyan, err := mysql.GetYiYanOne()
		if err != nil {
			zap.L().Error("操作获取一言失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		ResponseSuccess(c, yiyan)
	}

}
