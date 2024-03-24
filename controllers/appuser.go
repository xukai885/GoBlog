package controllers

import (
	"GoBlog/logic"
	"GoBlog/modules"

	_ "GoBlog/docs"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags 用户
// @Summary  用户注册
// @Description 用户注册
// @Accept  json
// @Product json
// @Param data body modules.AddAppUser true "请求参数"
// @Success 1000 {object} modules.AppUser
// @Router /api/v1/register [POST]
// RegisterHandler 注册函数
func RegisterHandler(c *gin.Context) {
	// 1.参数校验
	var user modules.AddAppUser
	if err := c.ShouldBindJSON(&user); err != nil {
		zap.L().Error("ShouldBindJSON err", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.业务处理
	if err := logic.Register(&user); err != nil {
		zap.L().Error("register err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, CodeSuccess.Msg())
}

// @Tags 用户
// @Summary  用户登陆
// @Accept  json
// @Product json
// @Param data body modules.AppUser true "请求参数"
// @Success 1000 {object} string "{"code": 1000, "msg": sucess,"data": token}"
// @Router /api/v1/login [POST]
func LoginHandler(c *gin.Context) {
	// 1.参数校验
	var user modules.AppUser
	if err := c.ShouldBindJSON(&user); err != nil {
		zap.L().Error("ShouldBindJSON err", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.业务处理
	token, err := logic.Login(&user)
	if err != nil {
		zap.L().Error("logic login err", zap.String("username", user.Username), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"accessToken": token,
		"userid":      user.UserID,
		"username":    user.Username,
	})
}
