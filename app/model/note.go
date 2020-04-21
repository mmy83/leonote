/*
@Time : 2020/4/21 10:43 上午
@Author : mmy83
@File : note.go
@Software: GoLand
*/

package model

import "time"

type Note struct {
	ID int64	`xorm:"id" json:"id"`
	NoteBookId int64	`xorm:"notebook_id" json:"notebook_id"`
	UserId int64 `xorm:"user_id" json:"user_id"`
	Title string `xorm:"title" json:"title"`
	Tags string `xorm:"tags" json:"tags"`
	Content string `xorm:"content" json:"content"`
	IsMD  int64  `xorm:"ismd" json:"is_md"`
	IsShow int64 `xorm:"isshow" json:"is_show"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`
}

func (n *Note) TableName() string {
	return "leo_notes"
}