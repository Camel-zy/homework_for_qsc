package controller

import (
	"errors"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @tags Interview
// @summary Create interview in event
// @description Create an interview in a specific event
// @router /interview [put]
// @accept json
// @param data body model.InterviewApi true "Interview Information"
// @produce json
func createInterview(c echo.Context) error {
	interviewRequest := model.InterviewApi{}
	if err := c.Bind(&interviewRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if err := c.Validate(&interviewRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}

	if err := model.CreateInterview(&interviewRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create interview fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "create interview success",
	})
}

// @tags Interview
// @summary Update interview
// @description Update an interview
// @router /interview [post]
// @param iid query uint true "Interview ID"
// @accept json
// @param data body model.InterviewApi false "Interview Information"
// @produce json
func updateInterview(c echo.Context) error {
	interviewRequest := model.InterviewApi{}
	err := c.Bind(&interviewRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	iid, err := utils.IsUnsignedInteger(c.QueryParam("iid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "iid needs to be specified correctly",
		})
	}
	interviewRequest.ID = iid

	if err := model.UpdateInterviewByID(&interviewRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "update interview fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "update interview success",
	})
}

// @tags Interview
// @summary Get interview
// @description Get information of an interview
// @router /interview [get]
// @param iid query uint true "Interview ID"
// @produce json
// @success 200 {object} model.InterviewApi
func getInterview(c echo.Context) error {
	iid, typeErr := utils.IsUnsignedInteger(c.QueryParam("iid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "iid need to be an unsigned integer",
		})
	}

	interview, itvErr := model.QueryInterviewByID(iid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "interview not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &interview,
	})
}

// @tags Interview
// @summary Get interview in event
// @description Get information of an interview in a specific event
// @router /event/interview [get]
// @param eid query uint true "Event ID"
// @param iid query uint true "Interview ID"
// @produce json
// @success 200 {object} model.InterviewApi
func getInterviewInEvent(c echo.Context) error {
	eid, errEid := utils.IsUnsignedInteger(c.QueryParam("eid"))
	iid, errIid := utils.IsUnsignedInteger(c.QueryParam("iid"))

	if errEid != nil || errIid != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid and iid need to be an unsigned integer",
		})
	}

	_, evtErr := model.QueryEventByID(eid)
	if errors.Is(evtErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "event not found",
		})
	}

	interview, itvErr := model.QueryInterviewByIDInEvent(eid, iid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "interview not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &interview,
	})
}

// @tags Interview
// @summary Get all interviews in event
// @description Get brief information of all interviews in a specific event
// @router /event/interview/all [get]
// @param eid query uint true "Event ID"
// @produce json
// @success 200 {array} model.Brief
func getAllInterviewInEvent(c echo.Context) error {
	eid, typeErr := utils.IsUnsignedInteger(c.QueryParam("eid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid need to be an unsigned integer",
		})
	}

	interviews, itvErr := model.QueryAllInterviewInEvent(eid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "event not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &interviews,
	})
}
