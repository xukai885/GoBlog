package mysql

import (
	"GoBlog/modules"
)

// AddBoke 新增博文
func AddBoke(boke *modules.Boke) (err error) {
	sqlStr := "insert into boke(`titlename`,`createtime`,`filename`,`md5filename`,`filepath`,`introduction`,`type`) values (?,?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, boke.TitleName, boke.CreateTime, boke.FileName, boke.Md5FileName, boke.FilePath, boke.Introduction, boke.Type)
	return
}

// UpdateBoke 更新博客
func UpdateBoke(boke *modules.Boke) (err error) {
	sqlStr := "UPDATE boke SET titlename =?,createtime =?,filename =?,md5filename =?,filepath =?,introduction = ?,type =? WHERE id =?"
	_, err = db.Exec(sqlStr, boke.TitleName, boke.CreateTime, boke.FileName, boke.Md5FileName, boke.FilePath, boke.Introduction, boke.Type, boke.Id)
	return
}

// GetBokeList 默认查看博客列表（id, 名字，时间，描述）
func GetBokeList() (boke []modules.Boke, err error) {
	sqlStr := "select `id`,`titlename`,`createtime`,`introduction`,`type` from boke order by createtime desc"
	err = db.Select(&boke, sqlStr)
	return
}

// GetTypeBoke 查看博客列表 按照分类id, (id, 名字，时间，描述)
func GetTypeBoke(typeid int) (boke []modules.Boke, err error) {
	sqlStr := "select `id`,`titlename`,`createtime`,`introduction`,`type` from boke where type=? order by createtime desc"
	err = db.Select(&boke, sqlStr, typeid)
	return
}

// GetBoke 查看博客详情 根据id
func GetBoke(id int) (boke modules.Boke, err error) {
	sqlStr := "select `titlename`,`createtime`,`filepath` from boke where id = ?"
	err = db.Get(&boke, sqlStr, id)
	return
}

// DeleteBoke 删除博客 根据id
func DeleteBoke(id int) (err error) {
	sqlStr := "delete from boke where id = ?"
	_, err = db.Exec(sqlStr, id)
	return
}

// DeleteType 删除分类 根据id
func DeleteType(id int) (err error) {
	sqlStr := "delete from boketype where id = ?"
	_, err = db.Exec(sqlStr, id)
	return
}

// Search 搜索标题
func Search(search string) (boke []modules.Boke, err error) {
	sqlStr := "select `id`,`titlename`,`createtime`,`introduction`,`type` from boke where titlename LIKE ?"
	err = db.Select(&boke, sqlStr, "%"+search+"%")
	return
}

// GetTimeArchiveNum 获取时间归档-数量
func GetTimeArchiveNum() (date []modules.TimeArchive, err error) {
	sqlStr := "SELECT CONCAT(YEAR(createtime), '-', LPAD(MONTH(createtime), 2, '0')) AS yearmounth,\n       COUNT(*) AS sum\nFROM boke\nGROUP BY yearmounth\nORDER BY yearmounth desc;"
	err = db.Select(&date, sqlStr)
	return
}

// GetTimeArchiveInfo 获取时间归档-详情
func GetTimeArchiveInfo(ym string) (boke []modules.Boke, err error) {
	sqlStr := "select `id`,`titlename`,`createtime`,`introduction`,`type` from boke where createtime LIKE ? order by createtime desc"
	err = db.Select(&boke, sqlStr, ym)
	return
}
