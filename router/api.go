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
	"leonote/app/http/api/check"
	"leonote/app/http/api/user"
	_ "leonote/docs"
)

func InitApiRouter(r *gin.Engine) *gin.Engine {

	r.GET("/check/ping", check.Ping)
	r.GET("/check/health", check.HealthCheck)



	r.GET("/api/user/:id", user.GetUser)
	r.GET("/api/user", user.List)






	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
