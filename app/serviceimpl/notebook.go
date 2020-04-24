/*
@Time : 2020/4/24 3:47 下午
@Author : mmy83
@File : notebook.go
@Software: GoLand
*/

package serviceimpl

import (
	"leonote/app/model"
	"leonote/app/service"
	"leonote/config"
	"leonote/database"
)

func NewNoteBook() service.NoteBook{
	return &noteBook{
		pageSize: config.CfgPage.GetInt("pageSize"),
	}
}

type noteBook struct {
	pageSize int
}

func (nb *noteBook) GetList(userId int64) ([]*model.NoteBook,error){
	noteBooks := make([]*model.NoteBook,0)
	err := database.Engine.Where("user_id=?",userId).Limit(nb.pageSize,0).Find(&noteBooks)
	return noteBooks,err
}

func (nb *noteBook) GetNoteBook(id int64, userId int64)(*model.NoteBook,bool,error){

	noteBook := &model.NoteBook{
		ID:     id,
		UserId: userId,
	}
	has,err := database.Engine.Get(noteBook)
	return noteBook,has, err
}