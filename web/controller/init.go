package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var e *echo.Echo
var testMain = false

func InitWebFramework() {
	testMain = viper.GetBool("rop.test")
	e = echo.New()
	e.HideBanner = true
	addRoutes(e)
}

func StartServer() {
	e.Logger.Fatal(e.Start(":1323"))
}
