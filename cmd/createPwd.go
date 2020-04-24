/*
@Time : 2020/4/24 4:52 下午
@Author : mmy83
@File : createPwd.go
@Software: GoLand
*/

package main


import (
	"fmt"

	"leonote/util"
)

func main()  {
	passwordOK := "111111"
	passwordERR := "adminxx"

	encodePW,err := util.CreatePassword(passwordOK)
	if err != nil {
		fmt.Printf("create pwd err: %s\n",err)
	}

	fmt.Printf("pwd: %s\n",encodePW)



	// 正确密码验证
	err = util.CompareHashAndPassword(encodePW,passwordOK)
	if err != nil {
		fmt.Printf("pw wrong\n")
	} else {
		fmt.Printf("pw ok\n")
	}

	// 错误密码验证
	err = util.CompareHashAndPassword(encodePW,passwordERR)
	if err != nil {
		fmt.Printf("pw wrong\n")
	} else {
		fmt.Printf("pw ok\n")
	}
}
