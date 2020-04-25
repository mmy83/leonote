/*
@Time : 2020/4/22 9:17 下午
@Author : mmy83
@File : login.go
@Software: GoLand
*/

package api

import (
	"github.com/gin-gonic/gin"
	"leonote/app/model"
	"leonote/app/serviceimpl"
	"leonote/config"
	"leonote/pkg/jwtauth"
	"leonote/pkg/response/jsonresponse"
	"leonote/util"
	"log"
)

func Login(c *gin.Context) {

	var loginUser model.Login
	if err:=c.BindQuery(&loginUser); err!=nil {
		log.Printf("query data err: %s\n",err)
		jsonresponse.NewJsonResponse(c,200705,"")
		return
	}

	userImpl := serviceimpl.NewUser()
	user,has,err := userImpl.GetUserByUserName(loginUser.UserName)
	if err != nil || !has{
		log.Printf("Get User Err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200702,"")
		return
	}

	err = util.CompareHashAndPassword(user.PassWord,loginUser.Password)
	if err != nil{
		log.Printf("密码错误: %s\n", err)
		jsonresponse.NewJsonResponse(c,200704,"")
		return
	}

	log.Printf("userData : %s\n", user)

	tokenString,expire,err := jwtauth.JwtToken.CreateTokenString(*user)

	if err != nil {
		log.Printf("token create error: %s\n",err)
		jsonresponse.NewJsonResponse(c,200706,"")
		return
	}

	cookieName := config.CfgJwt.GetString("cookie")
	domain := config.CfgApp.GetString("domain")
	c.SetCookie(
		cookieName,
		tokenString,
		int(expire.Unix()),
		"/",
		domain,
		false,
		true,
	)
	data := make(map[string]interface{})

	data["token"] = tokenString

	jsonresponse.NewJsonResponse(c,200601,data)
	return

}
