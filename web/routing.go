package web

import (
	"git.zjuqsc.com/rop/rop-back-neo/web/auth"
	"git.zjuqsc.com/rop/rop-back-neo/web/controller"
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.Use(auth.Middleware)
	api.GET("", controller.GetApiVersion)
	api.GET("/user", controller.GetUser)
	api.GET("/user/all", controller.GetAllUser)
	api.GET("/organization", controller.GetOrganization)
	api.GET("/organization/all", controller.GetAllOrganization)
	api.GET("/organization/department", controller.GetDepartmentUnderOrganization)
	api.GET("/organization/department/all", controller.GetAllDepartmentUnderOrganization)
}
