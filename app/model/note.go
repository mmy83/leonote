/*
@Time : 2020/4/21 10:43 上午
@Author : mmy83
@File : note.go
@Software: GoLand
*/

package model

import "time"

type Note struct {
	Id int64	`xorm:"id pk autoincr" json:"id"`
	NoteBookId int64	`xorm:"notebook_id" json:"notebook_id" form:"notebook_id" binding:"required"`
	UserId int64 `xorm:"user_id" json:"user_id"`
	Title string `xorm:"title" json:"title" form:"title" binding:"required"`
	Tags string `xorm:"tags" json:"tags" form:"tags"`
	Content string `xorm:"content" json:"content" form:"content" binding:"required"`
	IsMD  int64  `xorm:"ismd" json:"is_md" form:"ismd"`
	IsShow int64 `xorm:"isshow" json:"is_show" form:"isshow"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`
}

func (n *Note) TableName() string {
	return "leo_notes"
}