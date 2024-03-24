package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登陆")

const CtxUserIDKey = "userID"

// getCurrentUser 获取当前用户登陆的id
func getCurrentUser(c *gin.Context) (userID uint64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
