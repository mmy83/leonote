/*
@Time : 2020/4/24 10:26 下午
@Author : mmy83
@File : json_response.go
@Software: GoLand
*/

package jsonresponse

import (
	"github.com/gin-gonic/gin"

)

func NewJsonResponse(c *gin.Context,code int,data interface{}){

	msg := getMsg(code)
	httpCode := code/1000

	c.JSON(httpCode,gin.H{
		"code": msg.Code,
		"msg": msg.Msg,
		"data": data,
	})

}






