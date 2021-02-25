// @title Recruit Open Platform API
// @version 0.1
// @description This API will be used under staging environment.
// @host rop-neo-staging.zjuqsc.com
// @BasePath /api

package main

import (
	"git.zjuqsc.com/rop/rop-back-neo/conf"
	_ "git.zjuqsc.com/rop/rop-back-neo/docs"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/rpc"
	"git.zjuqsc.com/rop/rop-back-neo/web/controller"
	"gorm.io/driver/postgres"
)

func main() {
	conf.Init()

	model.Connect(postgres.Open(conf.GetDatabaseLoginInfo()))
	model.CreateTables()

	model.ConnectObjectStorage()

	rpc.Init()

	controller.InitWebFramework()
	controller.StartServer()
}
