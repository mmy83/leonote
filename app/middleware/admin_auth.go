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
	"leonote/pkg/response/jsonresponse"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context){
		if jwtauth.JwtAuth.IsAdmin(c) {
			c.Next()
			return
		}
		jsonresponse.NewJsonResponse(c,200708,"")
		c.Abort()
	}
}