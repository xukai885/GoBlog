package mysql

import (
	"GoBlog/modules"
	"GoBlog/pkg/PassAddrsa"
	"encoding/base64"
	"errors"

	"go.uber.org/zap"
)

// InsertUser 新增用户
func InsertUser(user *modules.AppUser) (err error) {
	// 密码加密

	pass, err := PassAddrsa.RsaEncrypt([]byte(user.Password))

	if err != nil {
		zap.L().Error("加密错误", zap.Error(err))
		return
	}
	// sql
	sqlStr := "insert into Appuser(`user_id`,`username`,`password`) values (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, base64.StdEncoding.EncodeToString(pass))
	return
}

// AppUserLogin 从数据库验证账号密码是否正确
func AppUserLogin(u *modules.AppUser) (err error) {
	userinputpass := u.Password
	sqlStr := "select `user_id`,`username`,`password` from Appuser where username=?"
	err = db.Get(u, sqlStr, u.Username)
	if err != nil {
		zap.L().Error("db.get err", zap.Error(err))
		return
	}

	// 密码解密
	b, _ := base64.StdEncoding.DecodeString(u.Password)
	passb, err := PassAddrsa.RsaDecrypt(b)
	if err != nil {
		zap.L().Error("PassAddrsa.RsaDecrypt", zap.Error(err))
		return
	}
	if userinputpass != string(passb) {
		zap.L().Error("密码错误", zap.Error(err))
		return errors.New("密码错误")
	}
	return

}
