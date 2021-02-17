package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var e *echo.Echo
var testController = false

func InitWebFramework() {
	e = echo.New()
	e.HideBanner = true
	addRoutes(e)
	e.Validator = &utils.CustomValidator{}

	logrus.Info("Echo framework initialized")
}

func StartServer() {
	e.Logger.Fatal(e.Start(":1323"))
}
