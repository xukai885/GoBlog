package logic

import (
	"GoBlog/dao/mysql"
	"GoBlog/modules"
)

// 写入数据库
func AddBoke(boke *modules.Boke) (err error) {
	return mysql.AddBoke(boke)
}
