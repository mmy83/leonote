/*
@Time : 2020/4/26 3:43 下午
@Author : mmy83
@File : tag.go
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
)

var CTag *Tag


func init(){
	CTag = &Tag{
		tagService: serviceimpl.NewTag(),
	}
}

type Tag struct {
	tagService service.Tag
}

// 标签列表
// @Summary 标签列表
// @Description 标签列表
// @Tags api/v1/tag
// @Success 200 {string} string "{"code":200600,"msg": "成功!","data":[]}"
// @Router /api/v1/tag [get]
// @Security
func (n *Tag) List(c *gin.Context){

	uid := jwtauth.GetUidByLoginner(c)
	tags,err := n.tagService.GetList(uid)
	if err != nil {
		log.Printf("list err: %s\n", err)
		jsonresponse.NewJsonResponse(c,200706,"")
		return
	}

	jsonresponse.NewJsonResponse(c,200600,tags)
	return
}
