/*
@Time : 2020/4/21 11:07 上午
@Author : mmy83
@File : note.go
@Software: GoLand
*/

package service

import "leonote/app/model"

type Note interface {
	GetNote(id int64, userId int64)(*model.Note,bool,error)
	GetList(userId int64,notebookId int64) ([]*model.Note,error)
	CreateNote(note *model.Note) (int64,error)

}
