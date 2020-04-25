/*
@Time : 2020/4/23 2:39 下午
@Author : mmy83
@File : login.go
@Software: GoLand
*/

package model


type Login struct {

	UserName string `form:"username" jsonresponse:"username" binding:"required"`
	Password string `form:"password" jsonresponse:"password" binding:"required"`

}