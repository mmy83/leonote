/*
@Time : 2020/4/21 10:53 上午
@Author : mmy83
@File : tag.go
@Software: GoLand
*/

package model

import "time"

type Tag struct {
	Id int64 `xorm:"id pk autoincr" jsonresponse:"id"`
	UserId int64 `xorm:"user_id" jsonresponse:"user_id"`
	TagName string `xorm:"tag_name" jsonresponse:"tag_name"`
	Count int64 `xorm:"count" jsonresponse:"count"`
	CreatedAt time.Time `xorm:"created" jsonresponse:"created_at"`
	UpdatedAt time.Time `xorm:"updated" jsonresponse:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" jsonresponse:"deleted_at"`
}

func (t *Tag) TableName() string {
	return "leo_tags"
}

