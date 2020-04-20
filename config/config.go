/*
@Time : 2020/4/18 4:50 下午
@Author : mmy83
@File : config.go
@Software: GoLand
*/

package config

import (
	"github.com/spf13/viper"
	"log"
)

var CfgDb *viper.Viper
var CfgApp *viper.Viper
var CfgJwt *viper.Viper
var CfgPage *viper.Viper

func init(){

	viper.SetConfigName("settings")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	CfgDb = viper.Sub("settings.database")
	if CfgDb == nil {
		panic("config not found settings.database")
	}

	CfgApp = viper.Sub("settings.application")
	if CfgApp == nil {
		panic("config not found settings.application")
	}

	CfgJwt = viper.Sub("settings.jwt")
	if CfgJwt == nil {
		panic("config not found settings.jwt")
	}

	CfgPage = viper.Sub("settings.page")
	if CfgPage == nil {
		panic("config not found settings.page")
	}

}
