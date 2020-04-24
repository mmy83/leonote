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
	"leonote/util"
	"log"
)

func Login(c *gin.Context) {

	var loginUser model.Login
	if err:=c.BindQuery(&loginUser); err!=nil {
		log.Printf("query data err: %s\n",err)
		c.JSON(200,gin.H{
			"code": 200,
			"msg": "缺少参数！",
		})
		return
	}

	userImpl := serviceimpl.NewUser()
	user,has,err := userImpl.GetUserByUserName(loginUser.UserName)
	if err != nil{
		log.Printf("Get User Err: %s\n", err)
		c.JSON(200,gin.H{
			"code": 404,
			"msg": "用户不存在！！",
		})
		return
	}
	if !has {
		c.JSON(200,gin.H{
			"code":404,
			"msg":"user not found",
		})
		return
	}

	err = util.CompareHashAndPassword(user.PassWord,loginUser.Password)
	if err != nil{
		log.Printf("密码错误: %s\n", err)
		c.JSON(200,gin.H{
			"code": 404,
			"msg": "密码错误！",
		})
		return
	}

	log.Printf("userData : %s\n", user)

	tokenString,expire,err := jwtauth.JwtToken.CreateTokenString(*user)

	if err != nil {
		log.Printf("token create error: %s\n",err)
		c.JSON(200,gin.H{
			"code": 404,
			"msg": "获取参数失败！",
		})
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
	c.JSON(200,gin.H{
		"code": 200,
		"msg": "login ok",
	})
	return
}
