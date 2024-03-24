package mysql

import "GoBlog/modules"

func UploadImage(image *modules.Images) (err error) {
	sqlStr := "insert into images(`name`,`createtime`,`path`,`size`,`imagesUrl`) values (?,?,?,?,?)"
	_, err = db.Exec(sqlStr, image.Name, image.CreateTime, image.Path, image.ImageSize, image.ImageUrl)
	return
}

func GetImages() (images []*modules.Images, err error) {
	sqlStr := "select `id`,`name`,`createtime`,`size`,`imagesUrl` from images"
	err = db.Select(&images, sqlStr)
	return
}
