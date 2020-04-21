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
	_ "leonote/docs"
)

func InitApiRouter(r *gin.Engine) *gin.Engine {

	r.GET("/check/ping", api.Ping)
	r.GET("/check/health", api.HealthCheck)



	r.GET("/api/user/:id", api.CUser.GetUser)
	r.GET("/api/user", api.CUser.List)






	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
