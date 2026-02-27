package main

import (
	"fmt"
	"go-chat/internal/config"
	"go-chat/internal/https_server"
	"go-chat/pkg/zlog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.GetConfig()
	host := conf.MainConfig.Host
	port := conf.MainConfig.Port
	addr := fmt.Sprintf("%s:%d", host, port)

	// 启动 HTTP 服务器
	go func() {
		if err := https_server.GE.Run(addr); err != nil {
			zlog.Fatal("server running fault: " + err.Error())
			return
		}
	}()

	zlog.Info("服务器启动成功，监听地址: " + addr)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zlog.Info("正在关闭服务器...")
	// 如果有需要清理的资源（如数据库连接、Redis等），在此添加
	// myredis.Close() 等

	zlog.Info("服务器已关闭")
}
