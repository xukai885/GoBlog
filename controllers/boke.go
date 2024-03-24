package controllers

import (
	"GoBlog/dao/mysql"
	"GoBlog/logic"
	"GoBlog/modules"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"go.uber.org/zap"
)

// CreateType 新增分类
func CreateType(c *gin.Context) {
	// 创建文章分类
	// 验证输入参数
	// 1.参数校验
	var boketype modules.BokeType
	if err := c.ShouldBindJSON(&boketype); err != nil {
		zap.L().Error("ShouldBindJSON err", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.业务处理
	if err := logic.CreateType(&boketype); err != nil {
		zap.L().Error("register err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	zap.L().Info("创建分类成功", zap.String("name", boketype.Typename))
	ResponseSuccess(c, boketype)
}

// GetType 获取分类
func GetType(c *gin.Context) {
	// 获取分类列表
	tyl, err := mysql.GetType()
	if err != nil {
		zap.L().Error("error sql", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	zap.L().Info("获取分类成功")
	ResponseSuccess(c, tyl)
}

// Search 搜索标题
func Search(c *gin.Context) {
	search := c.Query("search")
	boke, err := mysql.Search(search)
	if err != nil {
		zap.L().Error("查询数据库失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, boke)
}

// AddBoke 新增博文
func AddBoke(c *gin.Context) {
	// 校验 检查输入的信息，创建的时间，分类，简述
	CreateTimes := c.PostForm("createtime")
	Introductions := c.PostForm("introduction")
	Types := c.PostForm("type")

	file, err := c.FormFile("file")
	if err != nil {
		zap.L().Error("获取文件失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 修改文件名字,存储md5文件名字
	filenewname := strTomd5(file.Filename) + ".md"

	// 将文件保存到指定路径
	dst := "uploads/" + filenewname
	if err := c.SaveUploadedFile(file, dst); err != nil {
		zap.L().Error("保存文件失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 将字符串转换为整数
	typenum, err := strconv.Atoi(Types)
	if err != nil {
		zap.L().Error("字符转换整数失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	boke := modules.Boke{
		CreateTime:   StringChangeTime(CreateTimes),
		TitleName:    GetFileName(file.Filename),
		FileName:     file.Filename,
		Md5FileName:  filenewname,
		FilePath:     dst,
		Introduction: Introductions,
		Type:         typenum,
	}

	if err = logic.AddBoke(&boke); err != nil {
		zap.L().Error("数据库添加数据失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 业务操作
	zap.L().Info("创建博文成功", zap.String("titlename", boke.TitleName))
	ResponseSuccess(c, boke)
}

// GetBoke 查看博客详情 根据id
func GetBoke(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		zap.L().Error("ID获取错误")
		ResponseError(c, CodeInvalidParam)
		return
	} else {
		// 业务处理
		// 将字符串转换为整数
		idnum, err := strconv.Atoi(id)
		if err != nil {
			zap.L().Error("字符转换整数失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		// 根据id查询文件路径，打开文件并传递回去
		boke, err := mysql.GetBoke(idnum)
		if err != nil {
			zap.L().Error("数据库查询信息失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		// 读取文件内容
		content, err := ioutil.ReadFile(boke.FilePath)

		unsafe := blackfriday.MarkdownCommon(content)
		// 清除不信任的内容
		output := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

		if err != nil {
			zap.L().Error("无法读取文件", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}

		response := modules.BokeInfo{
			TitleName:  boke.TitleName,
			CreateTime: boke.CreateTime,
			Content:    string(output),
		}
		zap.L().Info("获取博文详情成功", zap.String("titlename", boke.TitleName))
		ResponseSuccess(c, response)

	}
}

// UpdateBoke 更新博客
func UpdateBoke(c *gin.Context) {
	id := c.PostForm("id")
	if len(id) == 0 {
		zap.L().Error("ID获取错误")
		ResponseError(c, CodeInvalidParam)
		return
	} else {
		// 业务处理
		// 将字符串转换为整数
		idnum, err := strconv.Atoi(id)
		if err != nil {
			zap.L().Error("字符转换整数失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		// 校验 检查输入的信息，创建的时间，分类，简述
		CreateTimes := c.PostForm("createtime")
		Introductions := c.PostForm("introduction")
		Types := c.PostForm("type")

		file, err := c.FormFile("file")
		if err != nil {
			zap.L().Error("获取文件失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}

		// 修改文件名字,存储md5文件名字
		filenewname := strTomd5(file.Filename) + ".md"

		// 将文件保存到指定路径
		dst := "uploads/" + filenewname
		if err := c.SaveUploadedFile(file, dst); err != nil {
			zap.L().Error("保存文件失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}

		// 将字符串转换为整数
		typenum, err := strconv.Atoi(Types)
		if err != nil {
			zap.L().Error("字符转换整数失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}

		boke := modules.Boke{
			Id:           int64(idnum),
			CreateTime:   StringChangeTime(CreateTimes),
			TitleName:    GetFileName(file.Filename),
			FileName:     file.Filename,
			Md5FileName:  filenewname,
			FilePath:     dst,
			Introduction: Introductions,
			Type:         typenum,
		}

		if err = mysql.UpdateBoke(&boke); err != nil {
			zap.L().Error("数据库添加数据失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		// 业务操作
		zap.L().Info("修改博文成功", zap.String("titlename", boke.TitleName))
		ResponseSuccess(c, boke)
	}
}

// GetTypeBoke 查看博客列表 按照分类id
func GetTypeBoke(c *gin.Context) {
	typeid := c.Query("id")
	if len(typeid) == 0 {
		zap.L().Error("ID获取错误")
		ResponseError(c, CodeInvalidParam)
		return
	} else {
		// 业务处理
		// 将字符串转换为整数
		typenum, err := strconv.Atoi(typeid)
		if err != nil {
			zap.L().Error("字符转换整数失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		boke, err := mysql.GetTypeBoke(typenum)
		if err != nil {
			zap.L().Error("数据库查询信息失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		zap.L().Info("获取博文列表成功，根据id", zap.String("typeid", typeid))
		ResponseSuccess(c, boke)
	}
}

// GetBokeList 默认查看博客列表（id, 名字，时间，描述）
func GetBokeList(c *gin.Context) {
	boke, err := mysql.GetBokeList()
	if err != nil {
		zap.L().Error("数据库查询数据失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	zap.L().Info("获取博文列表成功")
	ResponseSuccess(c, boke)
}

// DeleteBoke 删除博客 根据id
func DeleteBoke(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		zap.L().Error("ID获取错误")
		ResponseError(c, CodeInvalidParam)
		return
	} else {
		// 业务处理
		// 将字符串转换为整数
		idnum, err := strconv.Atoi(id)
		if err != nil {
			zap.L().Error("字符转换整数失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		// 删除操作
		if err = mysql.DeleteBoke(idnum); err != nil {
			zap.L().Error("数据库删除数据失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		zap.L().Info("博文删除成功", zap.String("id", id))
		ResponseSuccess(c, "删除成功")
	}
}

// DeleteType 删除分类 根据id
func DeleteType(c *gin.Context) {
	typeid := c.Query("id")
	if len(typeid) == 0 {
		zap.L().Error("ID获取错误")
		ResponseError(c, CodeInvalidParam)
		return
	} else {
		// 业务处理
		// 将字符串转换为整数
		typenum, err := strconv.Atoi(typeid)
		if err != nil {
			zap.L().Error("字符转换整数失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		// 数据库操作
		if err = mysql.DeleteType(typenum); err != nil {
			zap.L().Error("数据库删除数据失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		zap.L().Info("分类删除成功", zap.String("id", typeid))
		ResponseSuccess(c, "删除成功")
	}
}

// GetTimeArchiveNum 获取时间归档-数量
func GetTimeArchiveNum(c *gin.Context) {
	// 数据库操作
	date, err := mysql.GetTimeArchiveNum()
	if err != nil {
		zap.L().Error("获取时间分类数据失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回数据
	ResponseSuccess(c, date)
}

// GetTimeArchiveInfo 获取时间归档-详情
func GetTimeArchiveInfo(c *gin.Context) {
	YearParam := c.Param("param1")
	MounParam := c.Param("param2")
	// sql
	ym := fmt.Sprintf("%s-%s%%", YearParam, MounParam)

	bokelist, err := mysql.GetTimeArchiveInfo(ym)
	if err != nil {
		zap.L().Error("查询数据库失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if bokelist == nil {
		zap.L().Error("数据为空", zap.String("ym", ym))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, bokelist)
}

// 时间转换
func StringChangeTime(str string) (t time.Time) {
	layout := "2006-01-02"
	// 将字符串转换为time.Time类型
	t, err := time.Parse(layout, str)
	if err != nil {
		zap.L().Error("时间转换失败", zap.Error(err))
		return
	}
	return
}

// 当前项目路径
//func GetDir() (dir string) {
//	// 获取当前文件的绝对路径
//	absPath, err := filepath.Abs(os.Args[0])
//	if err != nil {
//		zap.L().Error("获取路径失败", zap.Error(err))
//		return
//	}
//
//	// 获取项目根目录路径
//	dir = filepath.Dir(filepath.Dir(absPath))
//	return
//
//}

// 获取文件名字，xxx.md 获取 xxx
func GetFileName(str string) (result string) {
	// 定义正则表达式
	reg := regexp.MustCompile(`^(.*)\.md$`)
	// 进行匹配和提取
	match := reg.FindStringSubmatch(str)
	if match == nil {
		zap.L().Error("未匹配到结果")
		return
	}
	result = match[1]
	return
}

// 对str进行md5 计算
func strTomd5(str string) (hashString string) {
	// 计算字符串的 MD5 哈希值
	hash := md5.Sum([]byte(str))
	// 将哈希值转换为十六进制字符串表示
	hashString = hex.EncodeToString(hash[:])
	return
}
