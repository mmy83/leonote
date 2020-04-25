/*
@Time : 2020/4/18 9:13 下午
@Author : mmy83
@File : user.go
@Software: GoLand
*/

package model

import (
	"time"
)

type User struct {
	Id int64 `xorm:"id pk autoincr" jsonresponse:"id"`
	UserName string `xorm:"username"  jsonresponse:"user_name" form:"username" binding:"required"`
	NickName string `xorm:"nickname"  jsonresponse:"nick_name" form:"nickname"`
	PassWord string `xorm:"password"  jsonresponse:"-" form:"password" binding:"required"`
	LastLoginTime int64 `xorm:"last_login_time" jsonresponse:"last_login_time"`
	IsAdmin int64 `xorm:"isadmin" jsonresponse:"is_admin"`
	CreatedAt time.Time `xorm:"created" jsonresponse:"created_at"`
	UpdatedAt time.Time `xorm:"updated" jsonresponse:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" jsonresponse:"deleted_at"`
}

func (u *User) TableName() string {
	return "leo_users"
}

