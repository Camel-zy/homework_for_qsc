package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getMyCalendar(c echo.Context) error {
	organizations, _ := model.QueryAllOrganization(c.Get("uid").(uint))
	var allEvents []model.Brief
	for _, org := range *organizations {
		events, _ := model.QueryAllEventInOrganization(org.ID)
		allEvents = append(allEvents, *events...)
	}
	var allInterviews []model.Brief
	for _, event := range allEvents {
		interviews, _ := model.QueryAllInterviewInEvent(event.ID)
		allInterviews = append(allInterviews, *interviews...)
	}
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &allInterviews,
	})
}
