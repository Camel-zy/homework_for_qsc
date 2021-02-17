package controller

import (
	"errors"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func addInterview(c echo.Context) error {
	iid, typeErr := utils.IsUnsignedInteger(c.QueryParam("iid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "iid need to be an unsigned integer",
		})
	}

	_, itvErr := model.QueryInterviewByID(iid)
	if !errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "interview has existed",
		})
	}

	req := &model.InterviewCreateRequest{}
	c.Bind(req)

	interview := &model.Interview{
		ID:             iid,
		Name:           req.Name,
		Description:    req.Description,
		EventID:        req.EventID,
		Event:          req.Event,
		DepartmentID:   req.DepartmentID,
		OtherInfo:      req.OtherInfo,
		Location:       req.Location,
		MaxInterviewee: req.MaxInterviewee,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
		UpdatedTime:    req.UpdatedTime,
	}

	if CrtItvErr := model.CreateInterview(interview); CrtItvErr != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "add interview fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "add interview success",
	})
}

func setInterview(c echo.Context) error {
	iid, typeErr := utils.IsUnsignedInteger(c.QueryParam("iid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "iid need to be an unsigned integer",
		})
	}

	interview := model.Interview{}
	c.Bind(&interview)

	interview.ID = iid

	if UpdItvErr := model.UpdateInterviewByID(&interview); UpdItvErr != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "set interview fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "set interview success",
	})
}

// @tags Interview
// @summary Get interview
// @description Get information of an interview
// @router /interview/{iid} [get]
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
// @router /event/interview/{eid}{iid} [get]
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
// @router /event/interview/all/{eid} [get]
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
