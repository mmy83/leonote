/*
@Time : 2020/4/24 5:34 下午
@Author : mmy83
@File : password.go
@Software: GoLand
*/

package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreatePassword(pwd string) (string,error){
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("create pwd err: %s\n",err)
		return "", err
	}
	//fmt.Println(hash)
	encodePW := string(hash)  // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return encodePW,nil
}

func CompareHashAndPassword(e string, p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return  nil
}