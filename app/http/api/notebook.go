/*
@Time : 2020/4/20 10:00 下午
@Author : mmy83
@File : notebook.go
@Software: GoLand
*/

package api

import (
	"github.com/gin-gonic/gin"
	"leonote/app/service"
	"leonote/app/serviceimpl"
	"leonote/pkg/jwtauth"
	"leonote/pkg/response/jsonresponse"
	"log"
	"strconv"
)

var CNoteBook *NoteBook



type NoteBook struct {
	noteBookService service.NoteBook
}

func init(){
	CNoteBook = &NoteBook{
		noteBookService: serviceimpl.NewNoteBook(),
	}
}

// 获取用户
// @Summary 获取用户
// @Description 获取用户
// @Tags api
// @Success 200 {string} string {"code":200,"msg": "成功!"}
// @Router /api/notebook/:id [get]
// @Security
func (nb *NoteBook)GetNoteBook(c *gin.Context){
	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	uid := jwtauth.GetUidByLoginner(c)
	noteBook,has,err := nb.noteBookService.GetNoteBook(id,uid)

	if err!=nil{
		log.Printf("err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200706,"")
		return
	}

	if !has{

		jsonresponse.NewJsonResponse(c,200600,"")
		return
	}else {

		jsonresponse.NewJsonResponse(c,200600,noteBook)
		return
	}
}

// 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags api
// @Success 200 {string} string "{"message": "pong"}"
// @Router /api/user [get]
// @Security
func (nb *NoteBook)List(c *gin.Context)  {

	uid := jwtauth.GetUidByLoginner(c)

	noteBooks,err := nb.noteBookService.GetList(uid)
	if err != nil {
		log.Printf("list err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200706,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200600,noteBooks)
	return
}


