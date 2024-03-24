package logic

import (
	"GoBlog/dao/mysql"
	"GoBlog/modules"
	"GoBlog/pkg/jwt"
	"GoBlog/pkg/snowflake"
)

// Register 用户注册
func Register(user *modules.AddAppUser) (err error) {
	// 生成user id
	userID, err := snowflake.GetID()
	if err != nil {
		return err
	}
	// 构造一个user
	newuser := &modules.AppUser{
		UserID:   userID,
		Username: user.Username,
		Password: user.Password,
	}
	// 保存在数据库
	return mysql.InsertUser(newuser)
}

// Login 用户登陆
func Login(user *modules.AppUser) (token string, err error) {
	//
	if err := mysql.AppUserLogin(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}
