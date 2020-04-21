/*
@Time : 2020/4/21 10:57 上午
@Author : mmy83
@File : note_tags.go
@Software: GoLand
*/

package model

import "time"

type NoteTag struct {
	ID int64 `xorm:"id" json:"id"`
	UserId int64 `xorm:"user_id" json:"user_id"`
	NoteId int64 `xorm:"note_id" json:"note_id"`
	TagId int64 `xorm:"tag_id" json:"tag_id"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`
}

func (nt *NoteTag) TableName() string {
	return "leo_note_tags"
}
