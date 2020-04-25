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
	"leonote/pkg/response/jsonresponse"
)

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context){
		if jwtauth.JwtAuth.IsLogin(c) {
			c.Next()
			return
		}
		jsonresponse.NewJsonResponse(c,200707,"")
		c.Abort()

	}
}