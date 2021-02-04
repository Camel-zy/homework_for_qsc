package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/web/middleware"
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo) {
	api := e.Group("/api")

	if !testMain {
		api.Use(middleware.Auth)
	}

	api.GET("", getApiVersion)

	user := api.Group("/user")
	user.GET("", getUser)
	user.GET("/all", getAllUser)

	organization := api.Group("/organization")
	organization.GET("", getOrganization)
	organization.GET("/all", getAllOrganization)
	organization.GET("/department", getDepartmentUnderOrganization)
	organization.GET("/department/all", getAllDepartmentUnderOrganization)
	organization.GET("/event", getEventOfOrganization)
	organization.GET("/event/all", getAllEventOfOrganization)

	event := api.Group("/event")
	event.GET("interview", getInterviewInEvent)
	event.GET("interview/all", getAllInterviewInEvent)

	/*
	CAUTIOUS: These routers are created only for demo
	This will be fixed on Feb 5th
	(RalXYZ)
	*/
	image := api.Group("/image")
	image.POST("", setImage)
	image.GET("", getImage)
}
