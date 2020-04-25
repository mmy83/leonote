/*
@Time : 2020/4/21 12:58 下午
@Author : mmy83
@File : jwtauth.go
@Software: GoLand
*/

package jwtauth

import (
	"github.com/gin-gonic/gin"
	"log"
)
var JwtAuth *jwtAuth

type jwtAuth struct {}

func init()  {
	JwtAuth = &jwtAuth{}
}

func (a *jwtAuth)IsLogin(c *gin.Context) bool{
	tokenStr,err := JwtExtractor.ExtractToken(c)
	if err!=nil {
		log.Printf("TOKEN GET ERROR: %v\n",err)
		return false
	}

	userData , err := JwtToken.ParseToken(tokenStr)

	if err!=nil {
		log.Printf("TOKEN PARSE ERROR: %v\n",err)
		return false
	}

	if userData == nil {
		return false
	}
	log.Printf("userData set: %s\n",userData)
	c.Set("JWT-DATA",userData)
	return true

}

func (a *jwtAuth) IsAdmin(c *gin.Context) bool {

	if ! a.IsLogin(c) {
		return false
	}
	userMap,ok :=GetLoginner(c)
	if !ok {
		return false
	}
	log.Printf("isadmin: %s\n",userMap.IsAdmin)
	if userMap.IsAdmin == 1  {
		return true
	}
	log.Println("is not admin")
	return false
}


func GetLoginner(c *gin.Context) (*JWTClaims,bool){
	user, exists:= c.Get("JWT-DATA")
	log.Printf("JWT-DATA get : %s\n", user)
	if !exists {
		return &JWTClaims{},false
	}
	userMap,ok := user.(*JWTClaims)
	if !ok {

		log.Printf("is ok?")
		return &JWTClaims{},false
	}
	log.Printf("userMap data: %s\n",userMap)

	return userMap,true
}

func GetUidByLoginner(c *gin.Context) int64 {
	user ,ok:= GetLoginner(c)
	if !ok {
		log.Printf("is ok?")
		return 0
	}
	return user.ID
}