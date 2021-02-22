package controller

import (
	"net/http"

	echoSwagger "github.com/swaggo/echo-swagger"

	"git.zjuqsc.com/rop/rop-back-neo/web/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func addRoutes(e *echo.Echo) {
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: viper.GetStringSlice("rop.allow_origins"),
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/api/version", getApiVersion)
	e.GET("/api/doc/*", echoSwagger.WrapHandler)

	api := e.Group("/api", middleware.Auth)

	user := api.Group("/user")
	user.GET("", getUser)
	user.GET("/all", getAllUser)

	my := api.Group("/my") // mainly for frontend rendering shortcut
	my.GET("/calendar", getMyCalendar)

	api.GET("/organization/all", getAllOrganization)

	organization := api.Group("/organization", middleware.GetOrganizationIdFromParam, middleware.AuthOrganization)
	organization.GET("", getOrganization)
	organization.GET("/department", getDepartmentInOrganization)
	organization.GET("/department/all", getAllDepartmentInOrganization)
	organization.GET("/event", getEventInOrganization)
	organization.GET("/event/all", getAllEventInOrganization)
	organization.PUT("/event", createEvent)

	event := api.Group("/event", middleware.SetEventOrganization, middleware.AuthOrganization)
	event.POST("", updateEvent)
	event.GET("", getEvent)
	event.GET("/interview", getInterviewInEvent)
	event.GET("/interview/all", getAllInterviewInEvent)
	event.PUT("/interview", createInterview, middleware.CheckDepartmentInOrganization)

	interview := api.Group("/interview", middleware.SetInterviewOrganization, middleware.AuthOrganization)
	interview.POST("", updateInterview)
	interview.GET("", getInterview)

	form := api.Group("/form") //
	form.PUT("", createForm)
	form.POST("", updateForm)
	form.GET("", getForm)

	answer := api.Group("/answer")
	answer.GET("", getAnswer)
	answer.POST("", updateAnswer)

	message := api.Group("/message") // TODO(TO/GA): auth middleware & test
	message.PUT("", addMessage)
	// message.GET("/all", getAllMessage) // TODO(TO/GA): wait until we know the logic

	template := api.Group("/messageTemplate", middleware.GetOrganizationIdFromParam, middleware.AuthOrganization) // TODO(TO/GA): test
	template.PUT("", addMessageTemplate)
	template.POST("", setMessageTemplate, middleware.AuthMessageTemplate)
	template.GET("", getMessageTemplate, middleware.AuthMessageTemplate)
	template.GET("/all", getAllMessageTemplate)

	avatar := api.Group("/avatar", middleware.CheckMinioStatus)
	avatar.POST("", setAvatar)
	avatar.GET("", getAvatar)
}
