/*
@Time : 2020/4/23 1:33 下午
@Author : mmy83
@File : admin_auth.go
@Software: GoLand
*/

package middleware

import (
	"github.com/gin-gonic/gin"
	"leonote/pkg/jwtauth"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context){
		if jwtauth.JwtAuth.IsAdmin(c) {
			c.Next()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 401,
			"msg":    "对不起，您没有管理员权限！",
		})
		c.Abort()
		return
	}
}