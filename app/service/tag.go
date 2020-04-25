/*
@Time : 2020/4/21 11:08 上午
@Author : mmy83
@File : tag.go
@Software: GoLand
*/

package service

import "leonote/app/model"

type Tag interface {
	GetList(userId int64) ([]*model.Tag,error)
}