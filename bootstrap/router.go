/*
@Time : 2020/4/17 9:54 下午
@Author : mmy83
@File : router.go
@Software: GoLand
*/

package bootstrap

import (
	"github.com/gin-gonic/gin"
	"leonote/router"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r = router.InitApiRouter(r)
	return r
}
