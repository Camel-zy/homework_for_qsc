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

	api := e.Group("/api")
	api.GET("/version", getApiVersion)
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

	organization := api.Group("/organization", middleware.AuthOrganization)
	organization.GET("", getOrganization)
	organization.GET("/department", getDepartmentInOrganization)
	organization.GET("/department/all", getAllDepartmentInOrganization)
	organization.GET("/event", getEventInOrganization)
	organization.GET("/event/all", getAllEventInOrganization)
	api.GET("/organization/all", getAllOrganization)

	event := api.Group("/event", middleware.SetEventOrganization, middleware.AuthOrganization)
	event.POST("", updateEvent)
	event.GET("", getEvent)
	event.GET("/interview", getInterviewInEvent)
	event.GET("/interview/all", getAllInterviewInEvent)
	api.PUT("/event", createEvent, middleware.SetReadOrganizationIdFromForm, middleware.AuthOrganization)

	interview := api.Group("/interview") // TODO(RalXYZ): add auth middleware
	interview.PUT("", createInterview)
	interview.POST("", updateInterview)
	interview.GET("", getInterview)

	message := api.Group("/message") // TODO(TO/GA): auth middleware & test
	message.PUT("", addMessage)
	message.GET("", getMessage)
	// message.GET("/all", getAllMessage) // TODO(TO/GA): wait until we know the logic

	template := api.Group("/messageTemplate", middleware.AuthOrganization) // TODO(TO/GA): test
	template.PUT("", addMessageTemplate)
	template.POST("", setMessageTemplate, middleware.AuthMessageTemplate)
	template.GET("", getMessageTemplate, middleware.AuthMessageTemplate)
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
