/*
@Time : 2020/4/17 10:31 下午
@Author : mmy83
@File : app.go
@Software: GoLand
*/

package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"leonote/config"
	"leonote/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func AppRun(){

	r := gin.New()

	util.InitDatabase()

	r = InitRouter(r)

	app := &http.Server{
		Addr:    config.CfgApp.GetString("host")+":"+ config.CfgApp.GetString("port"),
		Handler: r,
	}
	go func() {
		// 服务连接
		if err := app.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Server Run ", config.CfgApp.GetString("host")+":"+config.CfgApp.GetString("port"))
	log.Println("Enter Control + C Shutdown Server")
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
