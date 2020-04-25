/*
@Time : 2020/4/25 1:01 下午
@Author : mmy83
@File : note.go
@Software: GoLand
*/

package serviceimpl

import (
	"leonote/app/model"
	"leonote/app/service"
	"leonote/config"
	"leonote/database"
)

func NewNote() service.Note {
	return &note{
		pageSize: config.CfgPage.GetInt("pageSize"),
	}
}

type note struct {
	pageSize int
}

func (nb *note) GetList(userId int64,notebookId int64) ([]*model.Note,error){
	notes := make([]*model.Note,0)
	err := database.Engine.Where("notebook_id=? and user_id=?",notebookId,userId).Limit(nb.pageSize,0).Find(&notes)
	return notes,err
}

func (nb *note) GetNote(id int64, userId int64)(*model.Note,bool,error){

	note := &model.Note{
		ID:     id,
		UserId: userId,
	}
	has,err := database.Engine.Get(note)
	return note,has, err
}