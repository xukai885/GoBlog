package mysql

import (
	"GoBlog/modules"
)

// CreateType 新增分类
func CreateType(bt *modules.BokeType) (err error) {
	sqlStr := "insert into boketype(typename) values (?)"
	_, err = db.Exec(sqlStr, bt.Typename)
	return
}

// GetType 获取分类
func GetType() (tyl []modules.BokeType, err error) {
	sqlStr := "SELECT boketype.id, boketype.typename, IFNULL(count_table.count_sum, 0) AS count_sum\nFROM boketype\nLEFT JOIN (\n    SELECT type, COUNT(*) AS count_sum\n    FROM boke\n    GROUP BY type\n) count_table ON count_table.type = boketype.id;"
	err = db.Select(&tyl, sqlStr)
	return
}
