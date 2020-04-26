/*
@Time : 2020/4/26 12:35 上午
@Author : mmy83
@File : note_tag.go
@Software: GoLand
*/

package serviceimpl

import (
	"leonote/app/model"
	"leonote/app/service"
	"leonote/database"
)

func NewNoteTag() service.NoteTag{
	return &noteTag{

	}
}

type noteTag struct {
}


func (nt *noteTag)CreateNoteTag(noteTags []*model.NoteTag) (int64,error){
	return database.Engine.Insert(noteTags)
}