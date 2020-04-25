/*
@Time : 2020/4/17 9:34 下午
@Author : mmy83
@File : api.go
@Software: GoLand
*/

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"leonote/app/http/api"
	"leonote/app/middleware"
	_ "leonote/docs"
)

func InitApiRouter(r *gin.Engine) *gin.Engine {



	auth := r.Group("/api/v1")

	auth.POST("/login", api.Login)

	auth.Use(middleware.LoginAuth())
	{

		auth.GET("/check/ping", api.Ping)
		auth.GET("/check/health", api.HealthCheck)

		auth.GET("/user/:id", api.CUser.GetUser)
		auth.GET("/user", api.CUser.List)


		auth.GET("/notebook/:id", api.CNoteBook.GetNoteBook)
		auth.GET("/notebook", api.CNoteBook.List)

	}

	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.IsAdminAuth())
	{
		admin.GET("/user/:id", api.CUser.GetUser)
		admin.GET("/user", api.CUser.List)
	}






	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
