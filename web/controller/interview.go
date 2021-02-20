package controller

import (
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// @tags Interview
// @summary Create interview in event
// @description Create an interview in a specific event
// @router /event/interview [put]
// @param eid query uint true "Event ID"
// @param did query uint true "Department ID"
// @accept json
// @param data body model.InterviewRequest true "Interview Information"
// @produce json
func createInterview(c echo.Context) error {
	var eid, did uint
	err := echo.QueryParamsBinder(c).
		MustUint("eid", &eid).
		MustUint("did", &did).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid and did need to be an unsigned integer",
		})
	}

	interviewRequest := model.InterviewRequest{}
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

	if iid, err := model.CreateInterview(&interviewRequest, eid, did); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create interview fail",
		})
	} else {
		return c.JSON(http.StatusOK, &utils.Error{
			Code: "SUCCESS",
			Data: iid,
		})
	}
}

// @tags Interview
// @summary Update interview
// @description Update an interview
// @router /interview [post]
// @param iid query uint true "Interview ID"
// @accept json
// @param data body model.InterviewRequest false "Interview Information"
// @produce json
func updateInterview(c echo.Context) error {
	var iid uint
	err := echo.QueryParamsBinder(c).MustUint("iid", &iid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "iid needs to be an unsigned integer",
		})
	}

	interviewRequest := model.InterviewRequest{}
	err = c.Bind(&interviewRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}

	if err := model.UpdateInterviewByID(&interviewRequest, iid); err != nil {
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
// @success 200 {object} model.InterviewResponse
func getInterview(c echo.Context) error {
	var iid uint
	err := echo.QueryParamsBinder(c).MustUint("iid", &iid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "iid needs to be an unsigned integer",
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
// @success 200 {object} model.InterviewResponse
func getInterviewInEvent(c echo.Context) error {
	var eid, iid uint
	err := echo.QueryParamsBinder(c).
		MustUint("eid", &eid).
		MustUint("iid", &iid).
		BindError()
	if err != nil {
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
	var eid uint
	err := echo.QueryParamsBinder(c).MustUint("eid", &eid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid needs to be an unsigned integer",
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
