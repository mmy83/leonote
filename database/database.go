/*
@Time : 2020/4/18 5:19 下午
@Author : mmy83
@File : database.go
@Software: GoLand
*/

package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"leonote/config"
	"log"
)

var Engine *xorm.EngineGroup

func InitDatabase(){

	dbType := config.CfgDb.GetString("dbType")
	dbConfig := config.CfgDb.GetStringSlice("uri")

	var err error

	Engine, err = xorm.NewEngineGroup(dbType, dbConfig)
	if err != nil {
		log.Fatalf("database error: %s\n", err)
	}

}