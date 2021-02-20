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
// @router /organization/event [put]
// @param oid query uint true "Organization ID"
// @accept json
// @param data body model.EventRequest true "Event Information"
// @produce json
func createEvent(c echo.Context) error {
	var oid uint
	err := echo.QueryParamsBinder(c).MustUint("oid", &oid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid needs to be an unsigned integer",
		})
	}

	eventRequest := model.EventRequest{}
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

	if eid, err := model.CreateEvent(&eventRequest, oid); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create event fail",
		})
	} else {
		return c.JSON(http.StatusOK, &utils.Error{
			Code: "SUCCESS",
			Data: eid,
		})
	}
}

// @tags Event
// @summary Update event
// @description Update an event
// @router /event [post]
// @param eid query uint true "Event ID"
// @accept json
// @param data body model.EventRequest false "Event Information"
// @produce json
func updateEvent(c echo.Context) error {
	var eid uint
	err := echo.QueryParamsBinder(c).MustUint("eid", &eid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid needs to be an unsigned integer",
		})
	}

	eventRequest := model.EventRequest{}
	err = c.Bind(&eventRequest)
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

	if err = model.UpdateEventByID(&eventRequest, eid); err != nil {
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
// @success 200 {object} model.EventResponse
func getEvent(c echo.Context) error {
	var eid uint
	err := echo.QueryParamsBinder(c).MustUint("eid", &eid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid needs to be an unsigned integer",
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
// @success 200 {object} model.EventResponse
func getEventInOrganization(c echo.Context) error {
	var eid, oid uint
	err := echo.QueryParamsBinder(c).
		MustUint("eid", &eid).
		MustUint("oid", &oid).
		BindError()
	if err != nil {
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
	var oid uint
	err := echo.QueryParamsBinder(c).MustUint("oid", &oid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid needs to be an unsigned integer",
		})
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
