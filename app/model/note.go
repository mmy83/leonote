/*
@Time : 2020/4/21 10:43 上午
@Author : mmy83
@File : note.go
@Software: GoLand
*/

package model

import "time"

type Note struct {
	ID int64	`xorm:"id" jsonresponse:"id"`
	NoteBookId int64	`xorm:"notebook_id" jsonresponse:"notebook_id"`
	UserId int64 `xorm:"user_id" jsonresponse:"user_id"`
	Title string `xorm:"title" jsonresponse:"title"`
	Tags string `xorm:"tags" jsonresponse:"tags"`
	Content string `xorm:"content" jsonresponse:"content"`
	IsMD  int64  `xorm:"ismd" jsonresponse:"is_md"`
	IsShow int64 `xorm:"isshow" jsonresponse:"is_show"`
	CreatedAt time.Time `xorm:"created" jsonresponse:"created_at"`
	UpdatedAt time.Time `xorm:"updated" jsonresponse:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" jsonresponse:"deleted_at"`
}

func (n *Note) TableName() string {
	return "leo_notes"
}