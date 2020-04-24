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
// @Success 200 {string} string "{"message": "pong"}"
// @Router /api/notebook/:id [get]
// @Security
func (nb *NoteBook)GetNoteBook(c *gin.Context){
	id,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	data, exists:= c.Get("JWT-DATA")
	if !exists {
		log.Printf("JWT-DATA NOT FOUND")
	}
	user,ok := data.(*jwtauth.JWTClaims)
	if !ok {
		log.Printf("is not ok")
	}
	noteBook,has,err := nb.noteBookService.GetNoteBook(id,user.ID)

	if err!=nil{
		log.Fatalf("err: %s\n", err)
	}

	if !has{
		c.JSON(200,gin.H{})
	}else {
		c.JSON(200,noteBook)
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

	data, exists:= c.Get("JWT-DATA")
	if !exists {
		log.Printf("JWT-DATA NOT FOUND")
	}
	user,ok := data.(*jwtauth.JWTClaims)
	if !ok {
		log.Printf("is not ok")
	}

	log.Printf("parent get user: %s\n", user)

	noteBooks,err := nb.noteBookService.GetList(user.ID)
	if err != nil {
		log.Fatalf("list err: %s\n", err)
	}
	c.JSON(200, noteBooks)
}


