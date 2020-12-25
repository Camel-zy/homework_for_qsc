package main

import (
	"git.zjuqsc.com/rop/rop-back-neo/database"
	"git.zjuqsc.com/rop/rop-back-neo/database/utils"
	"git.zjuqsc.com/rop/rop-back-neo/web"
)

func main() {
	database.Connect(utils.ParseLoginInfo(utils.GetLoginInfo("conf/login.json")))
	database.Init()
	web.InitWebFramework()
	web.StartServer()
}
