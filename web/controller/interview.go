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
// @router /event/interview [put]
// @param eid query uint true "Event ID"
// @param did query uint true "Department ID"
// @param round query uint true "Round"
// @accept json
// @param data body model.InterviewRequest true "Interview Information"
// @produce json
func createInterview(c echo.Context) error {
	var eid, did, round uint
	err := echo.QueryParamsBinder(c).
		MustUint("eid", &eid).
		MustUint("did", &did).
		MustUint("round", &round).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid, did and rnd need to be an unsigned integer",
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

	if iid, err := model.CreateInterview(&interviewRequest, eid, did, round); err != nil {
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
// @success 200
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
			Data: "update interview failed",
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
		Data: interview,
	})
}

// @tags Interview
// @summary Delete interview
// @router /interview [delete]
// @param iid query uint true "Interview ID"
func deleteInterview(c echo.Context) error {
	var iid uint
	err := echo.QueryParamsBinder(c).MustUint("iid", &iid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "iid needs to be an unsigned integer",
		})
	}

	joinedInterview, err := model.QueryAllJoinedInterviewOfInterview(iid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "get interview failed",
		})
	}
	if len(*joinedInterview) != 0 {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "can't delete interview with interviewees participated",
		})
	}

	err = model.DeleteInterview(iid)
	if err != nil {
		if errors.Is(err, model.ErrNoRowsAffected) {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "NOT_FOUND",
				Data: "form not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "delete form failed",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "delete form success",
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
		Data: interview,
	})
}

// @tags Interview
// @summary Get all interviews in event
// @description Get brief information of all interviews in a specific event
// @router /event/department/interview/all [get]
// @param eid query uint true "Event ID"
// @param did query uint true "Department ID"
// @produce json
// @success 200 {array} model.InterviewResponse
func getAllInterviewOfDepartmentInEvent(c echo.Context) error {
	var eid, did uint
	err := echo.QueryParamsBinder(c).
		MustUint("eid", &eid).
		MustUint("did", &did).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid and did needs to be an unsigned integer",
		})
	}

	interviews, itvErr := model.QueryAllInterviewOfDepartmentInEvent(eid, did)
	if itvErr != nil && !errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "get interviews failed",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: interviews,
	})
}

// @tags Interview
// @summary Get all interviews of same round
// @description Get brief information of all interviews of the same round in a specific event and department
// @router /event/department/round/interview/all [get]
// @param eid query uint true "Event ID"
// @param did query uint true "Department ID"
// @param round query uint true "Round"
// @produce json
// @success 200 {array} model.InterviewResponse
func getAllInterviewOfRound(c echo.Context) error {
	var eid, did, round uint
	err := echo.QueryParamsBinder(c).
		MustUint("eid", &eid).
		MustUint("did", &did).
		MustUint("round", &round).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid, did and round need to be an unsigned integer",
		})
	}

	interviews, itvErr := model.QueryAllInterviewOfRound(eid, did, round)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "event not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: interviews,
	})
}
