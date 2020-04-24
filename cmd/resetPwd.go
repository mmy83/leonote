/*
@Time : 2020/4/24 5:05 下午
@Author : mmy83
@File : resetPwd.go
@Software: GoLand
*/

package main

import (
	"leonote/app/model"
	"leonote/database"
	"leonote/util"
	"log"
	"os"
)

func main(){
	args := os.Args[1:]
	argCount := len(args)

	if argCount !=2 {
		log.Fatalf("参数错误，参数个数: %d\n", argCount)
	}
	for key,value := range args{
		log.Printf("参数 %d: %s\n", key+1,value)
	}

	encodePW,err :=util.CreatePassword(args[1])

	if err != nil {
		log.Fatalf("create pwd err: %s\n",err)
	}

	database.InitDatabase()
	 user := model.User{
	 	UserName: args[0],
	 }

	has,err := database.Engine.Get(&user)
	if err != nil {
		log.Fatalf("获取用户错误: %s\n", err)
	}
	if !has {
		log.Fatalf("用户不存在！")
	}



	u := new(model.User)
	u.UserName = args[0]
	u.PassWord = encodePW

	affected, err := database.Engine.Where("ID = ?",user.ID).Update(u)
	if err != nil {
		log.Fatalf("修改错误: %s\n", err)
	}
	if affected > 0 {
		log.Printf("修改用户数量：%d\n\n", affected)
		log.Printf("重置用户密码成功，用户名：%s，新密码：%s\n",args[0],args[1])
	}
}
