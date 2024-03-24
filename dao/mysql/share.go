package mysql

import (
	"GoBlog/modules"

	"go.uber.org/zap"
)

func CreateShare(share *modules.ShareInfo) (err error) {
	sqlStr := "INSERT INTO share (title, url, date, type)VALUES (?, ?, ?, ?);"
	_, err = db.Exec(sqlStr, share.Title, share.Url, share.Date, share.Type)
	if err != nil {
		zap.L().Error("执行mysql失败", zap.Error(err))
		return
	}
	return
}

func GetShare() (share []modules.ShareInfo, err error) {
	sqlStr := "select `title`,`url`,`date`,`type` from share"
	err = db.Select(&share, sqlStr)
	if err != nil {
		zap.L().Error("执行mysql失败", zap.Error(err))
		return
	}
	return
}
