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
	event.GET("/department/round/interview/all", getAllInterviewOfRound)
	event.PUT("/interview", createInterview, middleware.CheckDepartmentInOrganization)
	event.GET("/department/round", getRoundNumOfJoindEvent)
	event.POST("/department/round", updateRoundNumOfJoinedEvent)

	interview := api.Group("/interview", middleware.SetInterviewOrganization, middleware.AuthOrganization)
	interview.POST("", updateInterview)
	interview.GET("", getInterview)

	interviewee := api.Group("/interviewee", middleware.AuthInterviewee)
	interviewee.POST("/options", updateIntervieweeOptions)
	interviewee.POST("/admit", admitInterviewee)
	interviewee.POST("/next", nextInterviewee)
	interviewee.POST("/reject", rejectInterviewee)

	form := api.Group("/form") //
	form.PUT("", createForm)
	form.POST("", updateForm)
	form.GET("", getForm)

	answer := api.Group("/answer")
	// answer.GET("", getAnswer)
	answer.PUT("", updateAnswer)

	// TODO(TO/GA): Delete it
	message := api.Group("/message", middleware.GetOrganizationIdFromParam, middleware.AuthOrganization)
	message.GET("/cost", getMessageCost)
	// message.PUT("/form", sendFormConfirmMessage)
	// message.PUT("/interview/select", sendInterviewSelectMessage)
	// message.PUT("/interview/confirm", sendInterviewConfirmMessage)
	// message.PUT("/reject", sendRejectMessage)

	object := api.Group("/object")
	object.POST("/create", createObject)
	object.POST("/seal", sealObject)
	object.GET("/get", getObject)
}
