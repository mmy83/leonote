/*
@Time : 2020/4/22 8:40 下午
@Author : mmy83
@File : auth.go
@Software: GoLand
*/

package middleware

import (
	"github.com/gin-gonic/gin"
	"leonote/pkg/jwtauth"
	"net/http"
)

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context){
		if jwtauth.JwtAuth.IsLogin(c) {
			c.Next()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 401,
			"msg":    "对不起，您没有该接口访问权限，请联系管理员",
		})
		c.Abort()
		return
	}
}