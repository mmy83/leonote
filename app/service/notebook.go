/*
@Time : 2020/4/20 10:49 下午
@Author : mmy83
@File : notebook.go
@Software: GoLand
*/

package service

import "leonote/app/model"

type NoteBook interface {
	GetNoteBook(id int64, userId int64)(*model.NoteBook,bool,error)
	GetList(userId int64) ([]*model.NoteBook,error)
}


