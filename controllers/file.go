package controllers

import (
	"GoBlog/dao/mysql"
	"GoBlog/modules"
	"GoBlog/settings"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadImage(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("image")
	if err != nil {
		zap.L().Error("获取文件失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 判断是否是图片
	if !isImageFile(file.Filename) {
		zap.L().Error("上传的不是图片")
		ResponseError(c, CodeServerBusy)
		return
	}

	newname := strTomd5(file.Filename+time.Now().String()) + filepath.Ext(file.Filename)
	// 将文件保存到指定路径
	dst := settings.Conf.FileConfig.ImagePath + newname
	if err := c.SaveUploadedFile(file, dst); err != nil {
		zap.L().Error("保存文件失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 写入数据库
	image_obj := modules.Images{
		CreateTime: time.Now(),
		Name:       newname,
		Path:       dst,
		ImageUrl:   settings.Conf.FileConfig.CtUrl + newname,
		ImageSize:  file.Size,
	}

	if err := mysql.UploadImage(&image_obj); err != nil {
		zap.L().Error("写入数据库失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回url
	ResponseSuccess(c, image_obj.ImageUrl)

}

// LookImages 展示图片
func LookImages(c *gin.Context) {
	image_name := c.Param("param1")
	// 判断是否是图片
	if !isImageFile(image_name) {
		zap.L().Error("获取的非图片")
		ResponseError(c, CodeServerBusy)
		return
	}

	// 判断图片是否存在
	image_path := settings.Conf.FileConfig.ImagePath + image_name
	_, err := os.Stat(image_path)
	if err != nil {
		zap.L().Error("图片不存在", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回图片展示
	imageBytes, err := os.ReadFile(image_path)
	if err != nil {
		zap.L().Error("打开图片失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Data(http.StatusOK, "image/jpeg", imageBytes)
}

// GetImages 获取所有图片
func GetImages(c *gin.Context) {
	images, err := mysql.GetImages()
	if err != nil {
		zap.L().Error("数据库查询失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, images)
}

func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".bmp"
}
