/*
@Time : 2020/4/21 10:36 上午
@Author : mmy83
@File : notebook.go
@Software: GoLand
*/

package model

import "time"

type NoteBook struct {
	ID int64 `xorm:"id" jsonresponse:"id"`
	ParentId int64 `xorm:"parent_id" jsonresponse:"parent_id"`
	UserId int64 `xorm:"user_id" jsonresponse:"user_id"`
	NoteBookName string `xorm:"notebook_name" jsonresponse:"notebook_name"`
	depth int64 `xorm:"depth" jsonresponse:"depth"`
	NoteNum int64 `xorm:"note_num" jsonresponse:"note_num"`
	CreatedAt time.Time `xorm:"created" jsonresponse:"created_at"`
	UpdatedAt time.Time `xorm:"updated" jsonresponse:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" jsonresponse:"deleted_at"`
}

func (nb *NoteBook) TableName() string{
	return "leo_notebooks"
}
