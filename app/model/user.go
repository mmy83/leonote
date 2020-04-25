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
	ID int64 `xorm:"id"  jsonresponse:"id"`
	UserName string `xorm:"username"  jsonresponse:"user_name"`
	NickName string `xorm:"nickname"  jsonresponse:"nick_name"`
	PassWord string `xorm:"password"  jsonresponse:"-"`
	LastLoginTime int64 `xorm:"last_login_time" jsonresponse:"last_login_time"`
	IsAdmin int64 `xorm:"isadmin" jsonresponse:"is_admin"`
	CreatedAt time.Time `xor,:"created" jsonresponse:"created_at"`
	UpdatedAt time.Time `xor,:"updated" jsonresponse:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" jsonresponse:"deleted_at"`
}

func (u *User) TableName() string {
	return "leo_users"
}

