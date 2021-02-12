package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/web/middleware"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func addRoutes(e *echo.Echo) {
	api := e.Group("/api")

	if !testController {
		api.Use(middleware.Auth)
		if !viper.GetBool("passport.enable") {
			middleware.MockPassport()
		}
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
	organization.GET("/event/interview", getInterviewInEvent)
	organization.GET("/event/interview/all", getAllInterviewInEvent)

	/*
	CAUTIOUS: These routers are created only for demo
	This will be fixed
	(RalXYZ)
	*/
	image := api.Group("/image")
	image.POST("", setImage)
	image.GET("", getImage)
}
