package controller

import (
	"github.com/labstack/echo/v4"
)

var e *echo.Echo
var testMain = false

func InitWebFramework(test bool) {
	testMain = test
	e = echo.New()
	e.HideBanner = true
	addRoutes(e)
}

func StartServer() {
	e.Logger.Fatal(e.Start(":1323"))
}
