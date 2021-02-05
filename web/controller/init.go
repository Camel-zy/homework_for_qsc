package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var e *echo.Echo
var testMain = false

func InitWebFramework() {
	testMain = viper.GetBool("rop.test")
	e = echo.New()
	e.HideBanner = true
	addRoutes(e)

	logrus.Info("Echo framework initialized")
}

func StartServer() {
	e.Logger.Fatal(e.Start(":1323"))
}
