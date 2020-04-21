/*
@Time : 2020/4/21 10:36 上午
@Author : mmy83
@File : notebook.go
@Software: GoLand
*/

package model

import "time"

type NoteBook struct {
	ID int64 `xorm:"id" json:"id"`
	ParentId int64 `xorm:"parent_id" json:"parent_id"`
	UserId int64 `xorm:"user_id" json:"user_id"`
	NoteBookName string `xorm:"notebook_name" json:"notebook_name"`
	depth int64 `xorm:"depth" json:"depth"`
	NoteNum int64 `xorm:"note_num" json:"note_num"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`
}

func (nb *NoteBook) TableName() string{
	return "leo_notebooks"
}
