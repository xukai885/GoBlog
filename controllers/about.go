package controllers

import (
	"GoBlog/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"go.uber.org/zap"
	"os"
)

func About(c *gin.Context) {
	// 获取about的md文件，并返回
	// 读取文件内容
	content, err := os.ReadFile(fmt.Sprintf("%s/about.md", settings.Conf.FileConfig.MdfilePath))

	unsafe := blackfriday.MarkdownCommon(content)
	// 清除不信任的内容
	output := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	if err != nil {
		zap.L().Error("无法读取文件", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, string(output))

}
