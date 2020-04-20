/*
@Time : 2020/4/18 5:41 下午
@Author : mmy83
@File : health.go
@Software: GoLand
*/

package check

import "github.com/gin-gonic/gin"

// 健康状况
// @Summary 健康状况 HealthCheck shows OK as the ping-pong result.
// @Description 健康状况
// @Tags check
// @Accept text/html
// @Produce text/html
// @Success 200 {string} string "OK"
// @Router /check/health [get]
// @BasePath
func HealthCheck(c *gin.Context) {
	c.JSON(200,"ok")
}


// ping检查
// @Summary ping
// @Description ping
// @Tags check
// @Success 200 {string} string "{"message": "pong"}"
// @Router /check/ping [get]
// @Security
func Ping(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}