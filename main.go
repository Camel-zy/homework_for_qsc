package main

import (
	"git.zjuqsc.com/rop/rop-back-neo/conf"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/web/controller"
	"gorm.io/driver/postgres"
)

func main() {
	conf.Init()
	model.Connect(postgres.Open(conf.GetDatabaseLoginInfo()))
	model.CreateTables()

	model.ConnectObjectStorage()

	controller.InitWebFramework()
	controller.StartServer()
}
