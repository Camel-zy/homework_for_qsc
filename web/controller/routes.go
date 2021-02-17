package controller

import (
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/web/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	my := api.Group("/my") // mainly for frontend rendering shortcut
	my.GET("/calendar", getMyCalendar)

	api.GET("/organization/all", getAllOrganization)
	organization := api.Group("/organization", middleware.AuthOrganization)
	organization.GET("", getOrganization)
	organization.GET("/department", getDepartmentInOrganization)
	organization.GET("/department/all", getAllDepartmentInOrganization)
	organization.GET("/event", getEventInOrganization)
	organization.GET("/event/all", getAllEventInOrganization)

	event := api.Group("/event", middleware.SetEventOrganization, middleware.AuthOrganization)
	event.PUT("", addEvent)
	event.POST("", setEvent)
	event.GET("", getEvent)
	event.GET("/interview", getInterviewInEvent)
	event.GET("/interview/all", getAllInterviewInEvent)

	interview := api.Group("/interview")
	interview.PUT("", addInterview)
	interview.POST("", setInterview)
	interview.GET("", getInterview)

	message := api.Group("/message") // TODO(TO/GA): auth middleware
	message.PUT("", addMessage)
	message.GET("", getMessage)
	// message.GET("/all", getAllMessage) TODO(TO/GA): wait until we know the logic
	template := message.Group("/template")
	template.PUT("", addMessageTemplate)
	template.POST("", setMessageTemplate)
	template.GET("", getMessageTemplate)
	template.GET("/all", getAllMessageTemplate)

	/*
		CAUTIOUS: These routers are created only for demo
		This will be fixed
		(RalXYZ)
	*/
	image := api.Group("/image")
	image.POST("", setImage)
	image.GET("", getImage)
}
