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
	"leonote/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func NewApp() *application{
	return &application{
		Env: config.CfgApp.GetString("env"),
		Name: config.CfgApp.GetString("name"),
		Host: config.CfgApp.GetString("host"),
		Port: config.CfgApp.GetString("port"),
		LogPath: config.CfgApp.GetString("logpath"),
		//ReadTimeout: config.CfgApp.GetDuration("readtimeout"),
		//WriteTimeout: config.CfgApp.GetDuration("writertimeout"),
	}
}

type application struct {

	Env               string
	Name              string
	Host              string
	Port              string

	//ErrorLog          *log.Logger
	LogPath           string

	ReadTimeout       time.Duration
	WriteTimeout      time.Duration

}

func (a *application)Run(){

	gin.SetMode(a.Env)

	r := gin.New()

	database.InitDatabase()

	r = InitRouter(r)

	app := &http.Server{
		Addr:    a.Host+":"+ a.Port,
		Handler: r,
		//ReadTimeout: a.ReadTimeout,
		//WriteTimeout: a.WriteTimeout,
	}
	go func() {
		// 服务连接
		if err := app.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Server Run ", a.Host+":"+a.Port)
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
