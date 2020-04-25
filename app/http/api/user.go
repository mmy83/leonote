/*
@Time : 2020/4/18 9:39 下午
@Author : mmy83
@File : user.go
@Software: GoLand
*/

package api

import (
	"github.com/gin-gonic/gin"
	"leonote/app/service"
	"leonote/app/serviceimpl"
	"leonote/pkg/response/jsonresponse"
	"log"
	"strconv"
)

var CUser *User

type User struct {
	userService service.User
}

func init(){

	CUser = &User{
		userService: serviceimpl.NewUser(),
	}
}

// 获取用户
// @Summary 获取用户
// @Description 获取用户
// @Tags api
// @Success 200 {string} string "{"message": "pong"}"
// @Router /api/user/:id [get]
// @Security
func (u *User)GetUser(c *gin.Context){
	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	user,has,err := u.userService.GetUser(id)

	if err!=nil  {

		log.Printf("err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200706,"")
		return
	}
	if !has {

		jsonresponse.NewJsonResponse(c,200601,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200601,user)
	return
}

// 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags api
// @Success 200 {string} string "{"message": "pong"}"
// @Router /api/user [get]
// @Security
func (u *User)List(c *gin.Context)  {

	users,err := u.userService.GetList()

	if err != nil {

		log.Printf("list err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200706,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200601,users)
	return
}
