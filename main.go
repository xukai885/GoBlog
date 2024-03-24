package main

import (
	"GoBlog/dao/mysql"
	"GoBlog/dao/redis"
	"GoBlog/logger"
	"GoBlog/logic"
	"GoBlog/pkg/snowflake"
	"GoBlog/routes"
	"GoBlog/settings"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 1.配置文件初始化
	if err := settings.Init(); err != nil {
		fmt.Printf("初始化配置文件失败", err)
		return
	}
	// 2.日志初始化
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("初始化日志失败", err)
		return
	}
	defer zap.L().Sync()
	// 3.mysql初始化
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("初始化mysql失败", err)
		return
	}
	defer mysql.Close()

	// 4.redis初始化
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("初始化redis失败", err)
		return
	}
	defer redis.Close()

	// 5.读取yiyan到redis
	if settings.Conf.YiYanSyncMysqlToRedis == true {
		logic.GetYiYanWriteRedis()
	}

	// 6.初始化雪花id
	if err := snowflake.Init(settings.Conf); err != nil {
		fmt.Printf("初始化雪花失败", err)
		return
	}

	// 7.注册路由

	r := routes.SetUp()

	// 6、启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v\n", err)
		}
	}()

	// 等待中断信号来优雅关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)                      // 创建一个接收信号的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGALRM) //此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号是才会向下执行
	zap.L().Info("shutdown server.....")
	// 创建一个5秒的超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//5秒内优雅关闭服务
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("server shutdown:", zap.Error(err))
	}
	zap.L().Info("server exiting")

}
