package controller

import (
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// @tags Event
// @summary Create event in organization
// @description Create an event in a specific organization
// @router /event [put]
// @accept json
// @param data body model.EventApi true "Event Information"
// @produce json
func createEvent(c echo.Context) error {
	eventRequest := model.EventApi{}
	if err := c.Bind(&eventRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if err := c.Validate(&eventRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if eventRequest.Status > 2 {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "the value in status field is illegal",
		})
	}

	eventRequest.OrganizationID = c.Get("oid").(uint)

	if err := model.CreateEvent(&eventRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create event fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: eventRequest.ID,
	})
}

// @tags Event
// @summary Update event
// @description Update an event
// @router /event [post]
// @param eid query uint true "Event ID"
// @accept json
// @param data body model.EventApi false "Event Information"
// @produce json
func updateEvent(c echo.Context) error {
	eventRequest := model.EventApi{}
	err := c.Bind(&eventRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if eventRequest.Status > 2 {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "the value in status field is illegal",
		})
	}

	eid, err := utils.IsUnsignedInteger(c.QueryParam("eid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid needs to be specified correctly",
		})
	}
	eventRequest.ID = eid

	if err = model.UpdateEventByID(&eventRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "update event fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "update event success",
	})
}

// @tags Event
// @summary Get event
// @description Get information of an event
// @router /event [get]
// @param eid query uint true "Event ID"
// @produce json
// @success 200 {object} model.EventApi
func getEvent(c echo.Context) error {
	eid, typeErr := utils.IsUnsignedInteger(c.QueryParam("eid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid need to be an unsigned integer",
		})
	}

	event, evtErr := model.QueryEventByID(eid)
	if errors.Is(evtErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "event not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &event,
	})
}

// @tags Event
// @summary Get event in organization
// @description Get information of an event in a specific organization
// @router /organization/event [get]
// @param oid query uint true "Organization ID"
// @param eid query uint true "Event ID"
// @produce json
// @success 200 {object} model.EventApi
func getEventInOrganization(c echo.Context) error {
	oid, errOid := utils.IsUnsignedInteger(c.QueryParam("oid"))
	eid, errEid := utils.IsUnsignedInteger(c.QueryParam("eid"))

	if errOid != nil || errEid != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid and eid need to be an unsigned integer",
		})
	}

	_, orgErr := model.QueryOrganizationById(oid)
	if errors.Is(orgErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "organization not found",
		})
	}

	event, evtErr := model.QueryEventByIDInOrganization(oid, eid)
	if errors.Is(evtErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "event not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &event,
	})
}

// @tags Event
// @summary Get all events in organization
// @description Get brief information of all events in a specific organization
// @router /organization/event/all [get]
// @param oid query uint true "Organization ID"
// @produce json
// @success 200 {array} model.Brief
func getAllEventInOrganization(c echo.Context) error {
	oid, typeErr := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid need to be an unsigned integer"},
		)
	}

	events, evtErr := model.QueryAllEventInOrganization(oid)
	if errors.Is(evtErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "organization not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &events,
	})
}
