package web

import (
	"git.zjuqsc.com/rop/rop-back-neo/web/controller"
	"git.zjuqsc.com/rop/rop-back-neo/web/middleware"
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo) {
	e.Use(middleware.Auth)
	api := e.Group("/api")
	api.POST("/user/", controller.GetUser)
	api.GET("/user/all", controller.GetAllUser)
}
