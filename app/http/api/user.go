/*
@Time : 2020/4/18 9:39 下午
@Author : mmy83
@File : user.go
@Software: GoLand
*/

package api

import (
	"github.com/gin-gonic/gin"
	"leonote/app/model"
	"leonote/app/service"
	"leonote/app/serviceimpl"
	"leonote/pkg/jwtauth"
	"leonote/pkg/response/jsonresponse"
	"leonote/util"
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


// 获取登录用户信息
// @Summary 获取登录用户信息
// @Description 获取登录用户信息
// @Tags api/v1/user/info
// @Success 200 {string} string "{"code": "200000","msg":"成功！","data":[]}"
// @Router /api/v1/user/info [get]
// @Security
func (u *User)GetUserInfo(c *gin.Context){

	user_id := jwtauth.GetUidByLoginner(c)
	user,has,err := u.userService.GetUser(user_id)

	log.Printf("user %d info: %s\n",user_id,user)
	if err!=nil  {

		log.Printf("err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200506,"")
		return
	}
	if !has {

		jsonresponse.NewJsonResponse(c,200000,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200000,user)
	return
}

// 获取指定用户
// @Summary 获取指定用户
// @Description 获取指定用户
// @Tags api/v1/admin/user
// @Param id path int true "用户id"
// @Success 200 {string} string "{"code": "200000","msg":"成功！","data":[]}"
// @Router /api/v1/admin/user/:id [get]
// @Security
func (u *User)GetUser(c *gin.Context){
	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 {
		log.Printf("err: 获取数据失败" )
		jsonresponse.NewJsonResponse(c,200506,"")
		return
	}
	user,has,err := u.userService.GetUser(id)

	log.Printf("user %d info: %s\n",id,user)
	if err!=nil  {

		log.Printf("err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200506,"")
		return
	}
	if !has {

		jsonresponse.NewJsonResponse(c,200000,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200000,user)
	return
}

// 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags api/v1/admin/user
// @Success 200 {string} string "{"code":200000,"msg": "成功！","data":[]}"
// @Router /api/v1/admin/user [get]
// @Security
func (u *User)List(c *gin.Context)  {

	users,err := u.userService.GetList()

	if err != nil {

		log.Printf("list err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200506,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200000,users)
	return
}




// 创建用户
// @Summary 创建用户
// @Description 创建用户
// @Tags api/v1/admin/user
// @Param username formData string true "用户名"
// @Param nickname formData string true "昵称"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"code":200000,"msg": "成功!","data":[]}"
// @Router /api/v1/admin/user/create [post]
// @Security
func (u *User) CreateUser(c *gin.Context) {
	var user model.User
	err := c.BindQuery(&user)
	if err != nil {
		jsonresponse.NewJsonResponse(c,200500,"")
		return
	}
	user.PassWord,err = util.CreatePassword(user.PassWord)
	if err != nil {
		jsonresponse.NewJsonResponse(c,200500,"")
		return
	}
	_,err = u.userService.CreateUser(&user)
	if err != nil {
		jsonresponse.NewJsonResponse(c,200509,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200000,user)
	return
}