/*
@Time : 2020/4/20 9:02 下午
@Author : mmy83
@File : user.go
@Software: GoLand
*/

package serviceimpl

import (
	"leonote/app/model"
	"leonote/app/service"
	"leonote/config"
	"leonote/database"
)


type user struct {

}

func New() service.User{
	return &user{}
}

func (u *user) GetUser(id int64) (*model.User,bool,error)  {
	user:= &model.User{}
	user.ID = id
	has, err := database.Engine.Get(user)
	return user,has, err
}

func (u *user) GetList() ([]*model.User,error)  {
	users := make([]*model.User,0)
	err := database.Engine.Limit(config.CfgPage.GetInt("pageSize"),0).Find(&users)
	return users,err
}