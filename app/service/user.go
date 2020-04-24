/*
@Time : 2020/4/20 8:57 下午
@Author : mmy83
@File : user.go
@Software: GoLand
*/

package service

import (
	"leonote/app/model"
)

type User interface {
	GetUser(id int64) (*model.User,bool,error)
	GetList() ([]*model.User,error)
	GetUserByUserName(username string) (*model.User,bool,error)
}

