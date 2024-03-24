package mysql

import "GoBlog/modules"

// GetYiYan 从mysql拿出所有的yiyan
func GetYiYan() (yiyan []modules.YiYan, err error) {
	sqlStr := "select `id`,`hitokoto`,`from_source`,`from_who` from yiyan"
	err = db.Select(&yiyan, sqlStr)
	return
}

func GetYiYanOne() (yiyan modules.YiYan, err error) {
	sqlStr := "select `id`,`hitokoto`,`from_source`,`from_who` from yiyan ORDER BY RAND() LIMIT 1"
	err = db.Get(&yiyan, sqlStr)
	return
}
