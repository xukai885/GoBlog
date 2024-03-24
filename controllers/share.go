package controllers

import (
	"GoBlog/dao/mysql"
	"GoBlog/modules"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CreateShare(c *gin.Context) {
	// 获取输入信息
	var share modules.ShareInfo
	if err := c.ShouldBindJSON(&share); err != nil {
		zap.L().Error("ShouldBindJSON err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	share.Date = time.Now()
	// 写入数据
	if err := mysql.CreateShare(&share); err != nil {
		zap.L().Error("数据写入失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, CodeSuccess)

}

func GetShare(c *gin.Context) {
	var share modules.Share
	shareinfo, err := mysql.GetShare()
	if err != nil {
		zap.L().Error("数据查询失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	share.Share = shareinfo
	share.Sum = len(shareinfo)
	ResponseSuccess(c, share)
}
