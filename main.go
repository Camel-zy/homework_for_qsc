package main

import (
	"git.zjuqsc.com/rop/rop-back-neo/conf"
	"git.zjuqsc.com/rop/rop-back-neo/database"
	"git.zjuqsc.com/rop/rop-back-neo/web/controller"
	"gorm.io/driver/postgres"
)

func main() {
	conf.InitConf()
	database.Connect(postgres.Open(conf.GetDatabaseLoginInfo()))
	database.CreateTables()
	controller.InitWebFramework(false)
	controller.StartServer()
}
