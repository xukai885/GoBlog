package logic

import (
	"GoBlog/dao/mysql"
	"GoBlog/modules"
)

// CreateType 增加分类
func CreateType(bt *modules.BokeType) (err error) {
	//增加数据库
	return mysql.CreateType(bt)
}
