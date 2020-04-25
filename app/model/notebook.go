/*
@Time : 2020/4/21 10:36 上午
@Author : mmy83
@File : notebook.go
@Software: GoLand
*/

package model

import (
	"time"
)

type NoteBook struct {
	Id int64 `xorm:"id pk autoincr" jsonresponse:"id"`
	ParentId int64 `xorm:"parent_id" jsonresponse:"parent_id" form:"parent_id"`
	UserId int64 `xorm:"user_id" jsonresponse:"user_id" form:"user_id"`
	NoteBookName string `xorm:"notebook_name" jsonresponse:"notebook_name" form:"notebook_name" binding:"required"`
	Depth int64 `xorm:"depth" jsonresponse:"depth"`
	NoteNum int64 `xorm:"note_num" jsonresponse:"note_num"`
	CreatedAt time.Time `xorm:"created" jsonresponse:"created_at"`
	UpdatedAt time.Time `xorm:"updated" jsonresponse:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" jsonresponse:"deleted_at"`
}


func (nb *NoteBook) TableName() string {
	return "leo_notebooks"
}
