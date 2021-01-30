package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/web/auth"
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo) {
	api := e.Group("/api")
	if !testMain {
		api.Use(auth.Middleware)
	}
	api.GET("", getApiVersion)
	api.GET("/user", getUser)
	api.GET("/user/all", getAllUser)
	api.GET("/organization", getOrganization)
	api.GET("/organization/all", getAllOrganization)
	api.GET("/organization/department", getDepartmentUnderOrganization)
	api.GET("/organization/department/all", getAllDepartmentUnderOrganization)
}
