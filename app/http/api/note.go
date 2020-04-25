/*
@Time : 2020/4/25 11:42 下午
@Author : mmy83
@File : note.gp
@Software: GoLand
*/

package api

import (
	"github.com/gin-gonic/gin"
	"leonote/app/model"
	"leonote/app/service"
	"leonote/app/serviceimpl"
	"leonote/pkg/jwtauth"
	"leonote/pkg/response/jsonresponse"
	"log"
	"strconv"
)

var CNote *Note


func init(){
	CNote = &Note{
		noteService: serviceimpl.NewNote(),
	}
}

type Note struct {
	noteService service.Note
}


func (n *Note) List(c *gin.Context){

	notebookId,_ := strconv.ParseInt(c.Query("notebook_id"),10,64)
	uid := jwtauth.GetUidByLoginner(c)
	notes,err := n.noteService.GetList(uid,notebookId)
	if err != nil {
		log.Printf("list err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200706,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200600,notes)
	return
}

func (n *Note) GetNote(c *gin.Context){
	noteId,_ := strconv.ParseInt(c.Param("id"),10,64)
	uid := jwtauth.GetUidByLoginner(c)

	note, has, err := n.noteService.GetNote(noteId,uid)

	if err!=nil  {

		log.Printf("err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200706,"")
		return
	}
	if !has {

		jsonresponse.NewJsonResponse(c,200600,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200600,note)
	return
}

func (n *Note) CreateNote(c *gin.Context) {

	uid := jwtauth.GetUidByLoginner(c)
	var note model.Note
	err := c.BindQuery(&note)
	note.UserId = uid

	if err != nil {
		jsonresponse.NewJsonResponse(c,200700,"")
		return
	}

	_,err = n.noteService.CreateNote(&note)
	if err != nil {
		jsonresponse.NewJsonResponse(c,200709,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200600,note)
	return
}