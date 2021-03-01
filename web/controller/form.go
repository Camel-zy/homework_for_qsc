package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// @tags Form
// @summary Create a form
// @description Create a form
// @router /form [put]
// @param oid query uint true "Organization ID"
// @accept json
// @param data body model.CreateFormRequest_ true "Form information"
// @success 200 {object} model.Form_
func createForm(c echo.Context) error {
	var oid uint
	err := echo.QueryParamsBinder(c).MustUint("oid", &oid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid needs to be an unsigned integer",
		})
	}

	formRequest := model.CreateFormRequest{}
	if err := c.Bind(&formRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if err := c.Validate(&formRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}

	var form model.Form
	copier.Copy(&form, formRequest)
	form.OrganizationID = oid

	form, err = model.CreateForm(&form)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create form fail",
		})
	}
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: form,
	})
}

// @tags Form
// @summary Update a form
// @description Update a form
// @router /form [post]
// @accept json
// @param fid query uint true "Form ID"
// @param data body model.UpdateFormRequest_ true "Form information"
// @success 200 {object} model.Form_
func updateForm(c echo.Context) error {
	var fid uint
	err := echo.QueryParamsBinder(c).MustUint("fid", &fid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid needs to be an unsigned integer",
		})
	}

	formRequest := model.UpdateFormRequest{}
	err = c.Bind(&formRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if err := c.Validate(&formRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}

	if err := model.UpdateFormByID(&formRequest, fid); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "update form fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "update form success",
	})
}

// @tags Form
// @summary Get a form
// @description Get a form
// @router /form [get]
// @param fid query uint true "Form ID"
// @param eid query uint true "Event ID"
// @success 200 {object} model.Form_
func getForm(c echo.Context) error {
	var fid, eid uint
	err := echo.QueryParamsBinder(c).
		MustUint("fid", &fid).
		MustUint("eid", &eid).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid, eid needs to be an unsigned integer",
		})
	}
	_, itvErr := model.QueryEventByID(eid)  // event, itvErr := model.QueryEventByID(eid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "event doesn't exist",
		})
	}
	/*
	_, itvErr = model.QueryEventHasForm(fid, eid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "form and event aren't relative",
		})
	}
	if event.Status != 3 {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "event isn't running",
		})
	}
	 */
	form, itvErr := model.QueryFormById(fid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "form not found",
		})
	}
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &form,
	})
}

// @tags Form
// @summary Get all form relations in eid
// @router /form/all [get]
// @param eid query uint true "Event ID"
// @success 200 {array} model.EventHasFormResponse
func getFormInEvent(c echo.Context) error {
	var eid uint
	err := echo.QueryParamsBinder(c).
		MustUint("eid", &eid).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid needs to be an unsigned integer",
		})
	}

	eventHasForm, err := model.QueryEventHasFormByEid(eid)

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &eventHasForm,
	})
}

// @tags Form
// @summary Get all forms in organization
// @description Get a form
// @router /organization/form/all [get]
// @param oid query uint true "Organization ID"
// @success 200 {array} model.Form_
func getAllFormInOrganization(c echo.Context) error {
	oid, typeErr := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid need to be an unsigned integer",
		})
	}
	forms, formErr := model.QueryAllFormByOid(oid)
	if errors.Is(formErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "organization not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &forms,
	})
}

// @tags Form
// @summary Get all interview options of interviewee
// @router /form/interviewee/interview/all [get]
// @param uuid query string true "Interviewee's UUID"
// @success 200 {array} model.InterviewResponse
func getAllInterviewOfInterviewee(c echo.Context) error {
	uuid, err := uuid.FromString(c.QueryParam("uuid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "can't parse uuid",
		})
	}
	interviewee, err := model.QueryIntervieweeByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "NOT_FOUND",
				Data: "interviewee not found",
			})
		}
	}

	interview := make([]model.InterviewResponse, 0)
	var interviewOptions []interface{}
	err = json.Unmarshal(interviewee.InterviewOptions, &interviewOptions)
	for _, v := range interviewOptions {
		iid := uint(v.(float64))
		temp, err := model.QueryInterviewByID(iid)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "get interviews fail",
			})
		}
		NowInterviewee, err := model.QueryNumberOfIntervieweesInInterviewByInterviewID(iid)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "get interviews fail",
			})
		}
		if NowInterviewee < int64(temp.MaxInterviewee) {
			var tempInterview model.InterviewResponse
			copier.Copy(&tempInterview, &temp)
			interview = append(interview, tempInterview)
		}
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &interview,
	})
}
