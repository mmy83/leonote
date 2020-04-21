/*
@Time : 2020/4/21 10:53 上午
@Author : mmy83
@File : tag.go
@Software: GoLand
*/

package model

import "time"

type Tag struct {
	ID int64 `xorm:"id" json:"id"`
	UserId int64 `xorm:"user_id" json:"user_id"`
	TagName string `xorm:"tag_name" json:"tag_name"`
	Count int64 `xorm:"count" json:"count"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`
}

func (t *Tag) TableName() string {
	return "leo_tags"
}

