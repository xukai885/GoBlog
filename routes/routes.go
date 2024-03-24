package routes

import (
	"GoBlog/controllers"
	"GoBlog/logger"
	"GoBlog/middlewares"
	"GoBlog/settings"
	"net/http"

	"github.com/penglongli/gin-metrics/ginmetrics"

	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	// 这个文件夹必须导入 否则会报错
)

func SetUp() *gin.Engine {
	r := gin.New()
	// get global Monitor object
	m := ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	// set middleware for gin
	m.Use(r)

	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	//v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	RegistrationsSwitch := settings.Conf.RegistrationsSwitch
	if RegistrationsSwitch == true {
		// 注册
		v1.POST("/register", controllers.RegisterHandler)
	} else if RegistrationsSwitch == false {
		zap.L().Info("关闭用户注册功能")
	}
	// 登陆
	v1.POST("/login", controllers.LoginHandler)
	v1.GET("/yiyan", controllers.GetYiYan)                                        // 获取一言
	v1.GET("/getboke", controllers.GetBoke)                                       // 查看博客详情 根据id
	v1.GET("/gettypeboke", controllers.GetTypeBoke)                               // 查看博客列表 按照分类id
	v1.GET("/getbokelist", controllers.GetBokeList)                               // 默认查看博客列表（id,名字，时间，描述）
	v1.GET("/gettype", controllers.GetType)                                       // 获取分类
	v1.POST("/search", controllers.Search)                                        // 搜索文章标题
	v1.GET("/getTimeArchive", controllers.GetTimeArchiveNum)                      // 获取时间归档-数量
	v1.GET("/getTimeArchiveInfo/:param1/:param2", controllers.GetTimeArchiveInfo) // 获取时间归档-详情
	v1.GET("/look_images/:param1", controllers.LookImages)                        // 获取图片-展示
	v1.GET("/share", controllers.GetShare)                                        // 获取分享的信息
	v1.Use(middlewares.JWTAuthMiddleware())                                       // 应用jwt认证中间件
	{
		v1.GET("/ping", controllers.Ping)
		v1.POST("/createtype", controllers.CreateType) // 新增分类
		v1.POST("/addboke", controllers.AddBoke)       // 新增博文
		v1.POST("/deleteboke", controllers.DeleteBoke) // 删除博客 根据id
		v1.POST("/deletetype", controllers.DeleteType) // 删除分类 根据id
		v1.POST("/update", controllers.UpdateBoke)     // 更新博客
		v1.POST("/upload", controllers.UploadImage)    // 上传图片
		v1.GET("/images", controllers.GetImages)       // 获取所有图片
		v1.POST("/share", controllers.CreateShare)     // 创建分享的信息
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
