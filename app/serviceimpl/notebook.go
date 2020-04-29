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
	"log"
)

func NewNoteBook() service.NoteBook{
	return &noteBook{
		pageSize: config.CfgPage.GetInt("pageSize"),
	}
}

type noteBook struct {
	pageSize int
}


func (nb *noteBook) GetList(userId int64) ([]*model.NoteBookTree,error){
	noteBooks := make([]*model.NoteBook,0)
	err := database.Engine.Where("user_id=?",userId).Limit(nb.pageSize,0).Find(&noteBooks)
	log.Printf("noteBooks: %s\n",noteBooks)
	return nb.createNoteBookTree(0,noteBooks),err
}

func (nb *noteBook) createNoteBookTree(parent_id int64,noteBooks []*model.NoteBook)  []*model.NoteBookTree {

	noteBookTrees := make([]*model.NoteBookTree,0)
	for _,noteBook := range noteBooks {
		if parent_id == noteBook.ParentId {
			noteBookTrees = append(noteBookTrees,&model.NoteBookTree{
				NoteBook:  noteBook,
				Childrens: nb.createNoteBookTree(noteBook.Id,noteBooks),
			})

		}
	}
	return noteBookTrees

}


func (nb *noteBook) GetNoteBook(id int64, userId int64)(*model.NoteBook,bool,error){

	noteBook := &model.NoteBook{
		Id:     id,
		UserId: userId,
	}
	has,err := database.Engine.Get(noteBook)
	return noteBook,has, err
}

func (nb *noteBook) CreateNoteBook(book *model.NoteBook) (int64,error){

	log.Printf("notebook: %v\n", book)
	affected, err := database.Engine.Insert(book)

	log.Printf("notebook after: %v\n",book)
	return affected,err

}