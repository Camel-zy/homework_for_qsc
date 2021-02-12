package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var e *echo.Echo
var testController = false

func InitWebFramework() {
	e = echo.New()
	e.HideBanner = true
	addRoutes(e)

	logrus.Info("Echo framework initialized")
}

func StartServer() {
	e.Logger.Fatal(e.Start(":1323"))
}
