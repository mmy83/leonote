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


// 指定笔记本下的笔记列表
// @Summary 指定笔记本下的笔记列表
// @Description 指定笔记本下的笔记列表
// @Tags api/v1/note
// @Param notebook_id query string true "笔记本id"
// @Success 200 {string} string "{"code":200600,"msg": "成功!","data":[]}"
// @Router /api/v1/note [get]
// @Security
func (n *Note) List(c *gin.Context){

	notebookId,_ := strconv.ParseInt(c.Query("notebook_id"),10,64)
	if notebookId <= 0 {
		log.Printf("err: 获取数据失败" )
		jsonresponse.NewJsonResponse(c,200706,"")
		return
	}
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



// 获取指定笔记
// @Summary 获取指定笔记
// @Description 获取指定笔记
// @Tags api/v1/note
// @Param id path string true "笔记id"
// @Success 200 {string} string "{"code":200600,"msg": "成功!","data":[]}"
// @Router /api/v1/note/:id [get]
// @Security
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


// 创建指定笔记本下的笔记
// @Summary 创建指定笔记本下的笔记
// @Description 创建指定笔记本下的笔记
// @Tags api/v1/note
// @Param notebook_id formData string true "笔记本id"
// @Param title formData string true "笔记标题"
// @Param tags formData string true "笔记tags"
// @Param content formData string true "笔记内容"
// @Success 200 {string} string "{"code":200600,"msg": "成功!","data":[]}"
// @Router /api/v1/note/create [post]
// @Security
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