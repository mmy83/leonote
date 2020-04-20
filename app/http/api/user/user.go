/*
@Time : 2020/4/18 9:39 下午
@Author : mmy83
@File : user.go
@Software: GoLand
*/

package user

import (
	"github.com/gin-gonic/gin"
	"leonote/app/model"
	"log"
	"strconv"
)

// 获取用户
// @Summary 获取用户
// @Description 获取用户
// @Tags api
// @Success 200 {string} string "{"message": "pong"}"
// @Router /api/user/:id [get]
// @Security
func GetUser(c *gin.Context){
	user := &model.User{}
	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	has,err := user.Get(id)

	if err!=nil{
		log.Fatalf("err: %s\n", err)
	}

	if !has{
		c.JSON(200,gin.H{})
	}else {
		c.JSON(200,user)
	}
}

// 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags api
// @Success 200 {string} string "{"message": "pong"}"
// @Router /api/user [get]
// @Security
func List(c *gin.Context)  {
	user := &model.User{}
	users,err := user.List()
	if err != nil {
		log.Fatalf("list err: %s\n", err)
	}
	c.JSON(200, users)
}
