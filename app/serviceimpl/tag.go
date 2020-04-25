/*
@Time : 2020/4/25 1:18 下午
@Author : mmy83
@File : tag.go
@Software: GoLand
*/

package serviceimpl

import (
	"leonote/app/model"
	"leonote/database"
)

type tag struct {
	pageSize int
}

func (t *tag) GetList(userId int64) ([]*model.Tag,error){

	tags := make([]*model.Tag,0)
	err := database.Engine.Where("user_id=?",userId).Limit(t.pageSize,0).Find(&tags)
	return tags,err

}