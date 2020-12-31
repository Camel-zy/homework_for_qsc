package main

import (
	"git.zjuqsc.com/rop/rop-back-neo/conf"
	"git.zjuqsc.com/rop/rop-back-neo/database"
	"git.zjuqsc.com/rop/rop-back-neo/web"
)

func main() {
	conf.InitConf()
	database.Connect(conf.GetDatabaseLoginInfo())
	database.Init()
	web.InitWebFramework()
	web.StartServer()
}
