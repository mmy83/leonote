/*
@Time : 2020/4/22 4:10 下午
@Author : mmy83
@File : auth_test.go
@Software: GoLand
*/

package test

import (
	"leonote/pkg/jwtauth"
	"log"
	"testing"
)

func TestAuth(t *testing.T) {
	user:= jwtauth.UserData{}
	user["id"] = "1"
	user["username"] ="mmy83"
	user["nickname"] ="mmy83"

	jwtauth := jwtauth.NewJwtAuth()
	str,_,_:=jwtauth.CreateToken(user)
	log.Print(str)


	claims,_ := jwtauth.GetClaimsFromToken(str)

	log.Print(claims)
}