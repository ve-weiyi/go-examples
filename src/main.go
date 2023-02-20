package main

import (
	"context"
	"github.com/ve-weiyi/go-examplse/src/core/database"
	"github.com/ve-weiyi/go-examplse/src/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("init invoke")
}

func main() {
	database.Connect()
	defer database.CloseDB()

	routers := router.InitRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routers,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("server run success on", server.Addr)
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("监听端口发生错误: %s\n", err)
		}
	}()

	// 创建无缓冲的通道。并等待中断信号。
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	//一直阻塞，直到收到信号。
	<-quit // 从quit渠道中接收值，忽略结果
	log.Printf("关闭服务器...")

	// 关闭服务器（设置 5 秒的超时时间），如果在5s内未关闭则会报错。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("服务器关闭:", err)
	}
	log.Printf("服务器已关闭")
}
