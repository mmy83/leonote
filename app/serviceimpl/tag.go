/*
@Time : 2020/4/25 1:18 下午
@Author : mmy83
@File : tag.go
@Software: GoLand
*/

package serviceimpl

import (
	"leonote/app/model"
	"leonote/app/service"
	"leonote/database"
)

func NewTag() service.Tag{
	return &tag{}
}

type tag struct {

}

func (t *tag) GetList(userId int64) ([]*model.Tag,error){

	tags := make([]*model.Tag,0)
	err := database.Engine.Where("user_id=?",userId).Find(&tags)
	return tags,err

}

func (t *tag) CreateTags(tags []*model.Tag) (int64,error) {

	return database.Engine.Insert(tags)

}