/*
@Time : 2020/4/21 10:57 上午
@Author : mmy83
@File : note_tag.go
@Software: GoLand
*/

package model

import "time"

type NoteTag struct {
	ID int64 `xorm:"id" jsonresponse:"id"`
	UserId int64 `xorm:"user_id" jsonresponse:"user_id"`
	NoteId int64 `xorm:"note_id" jsonresponse:"note_id"`
	TagId int64 `xorm:"tag_id" jsonresponse:"tag_id"`
	CreatedAt time.Time `xorm:"created" jsonresponse:"created_at"`
	UpdatedAt time.Time `xorm:"updated" jsonresponse:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" jsonresponse:"deleted_at"`
}

func (nt *NoteTag) TableName() string {
	return "leo_note_tags"
}
