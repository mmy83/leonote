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
	Id int64 `xorm:"id pk autoincr" json:"id"`
	UserName string `xorm:"username"  json:"user_name" form:"username" binding:"required"`
	NickName string `xorm:"nickname"  json:"nick_name" form:"nickname"`
	//Avatar string `xorm:"avatar"  json:"avatar" form:"avatar"`
	PassWord string `xorm:"password"  json:"-" form:"password" binding:"required"`
	LastLoginTime int64 `xorm:"last_login_time" json:"last_login_time"`
	IsAdmin int64 `xorm:"isadmin" json:"is_admin"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`
}

func (u *User) TableName() string {
	return "leo_users"
}

