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
	ID int64 `xorm:"id"  json:"id"`
	UserName string `xorm:"username"  json:"user_name"`
	NickName string `xorm:"nickname"  json:"nick_name"`
	PassWord string `xorm:"password"  json:"-"`
	LastLoginTime int64 `xorm:"last_login_time" json:"last_login_time"`
	IsAdmin int64 `xorm:"isadmin" json:"is_admin"`
	CreatedAt time.Time `xor,:"created" json:"created_at"`
	UpdatedAt time.Time `xor,:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`
}

func (u *User) TableName() string {
	return "leo_users"
}
