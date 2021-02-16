package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/web/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"github.com/swaggo/echo-swagger"
	"net/http"
)

func addRoutes(e *echo.Echo) {
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: viper.GetStringSlice("rop.allow_origins"),
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/api", getApiVersion)

	api := e.Group("/api")
	api.GET("/doc/*", echoSwagger.WrapHandler)

	if !testController {
		api.Use(middleware.Auth)
		if !viper.GetBool("passport.enable") {
			middleware.MockPassport(middleware.MockQscPassportServiceWillPass)
		}
	}

	user := api.Group("/user")
	user.GET("", getUser)
	user.GET("/all", getAllUser)

	my := api.Group("/my")  // mainly for frontend rendering shortcut
	my.GET("/calendar", getMyCalendar)

	organization := api.Group("/organization")
	organization.GET("", getOrganization)
	organization.GET("/all", getAllOrganization)
	organization.GET("/department", getDepartmentInOrganization)
	organization.GET("/department/all", getAllDepartmentInOrganization)
	organization.GET("/event", getEventInOrganization)
	organization.GET("/event/all", getAllEventInOrganization)

	event := api.Group("/event")
	event.GET("/interview", getInterviewInEvent)
	event.GET("/interview/all", getAllInterviewInEvent)

	/*
	CAUTIOUS: These routers are created only for demo
	This will be fixed
	(RalXYZ)
	*/
	image := api.Group("/image")
	image.POST("", setImage)
	image.GET("", getImage)
}
