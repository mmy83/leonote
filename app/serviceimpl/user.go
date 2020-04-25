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
	pageSize int
}

func NewUser() service.User{
	return &user{
		pageSize: config.CfgPage.GetInt("pageSize"),
	}
}

func (u *user) GetUser(id int64) (*model.User,bool,error)  {
	user:= &model.User{}
	user.Id = id
	has, err := database.Engine.Get(user)
	return user,has, err
}

func (u *user) GetList() ([]*model.User,error)  {
	users := make([]*model.User,0)
	err := database.Engine.Limit(u.pageSize,0).Find(&users)
	return users,err
}


func (u *user)GetUserByUserName(username string) (*model.User,bool,error){
	user:= &model.User{}
	user.UserName = username
	has, err := database.Engine.Get(user)
	return user,has, err
}

func (u *user) CreateUser(user *model.User) (int64,error){

	affected, err := database.Engine.Insert(user)

	return affected,err
}

