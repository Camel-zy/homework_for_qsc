package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @tags Form
// @summary Create a form
// @description Create a form
// @router /form [put]
// @accept  json
// @param data body model.FormApi_ true "Form information"
// @success 200
func createForm(c echo.Context) error {
	FormRequest := model.FormApi{}
	if err := c.Bind(&FormRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if err := c.Validate(&FormRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if err := json.Valid(FormRequest.Content); err != true {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "wrong format, JSON required",
		})
	}
	if err := model.CreateForm(&FormRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create form fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "create form success",
	})
}

// @tags Form
// @summary Update a form
// @description Update a form
// @router /form [post]
// @accept  json
// @param data body model.FormApi_ true "Form information"
// @success 200
func updateForm(c echo.Context) error {
	FormRequest := model.FormApi{}
	err := c.Bind(&FormRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	fid, err := utils.IsUnsignedInteger(c.QueryParam("fid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid needs to be specified correctly",
		})
	}
	FormRequest.ID = fid
	if err := json.Valid(FormRequest.Content); err != true {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "wrong format, JSON required",
		})
	}
	if err := model.UpdateFormByID(&FormRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "update Form fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "update Form success",
	})
}

// @tags Form
// @summary Get a form
// @description Get a form
// @router /form [get]
// @param fid query uint true "Form ID"
// @success 200 {object} model.Form_
func getForm(c echo.Context) error {
	fid, typeErr := utils.IsUnsignedInteger(c.QueryParam("fid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid need to be an unsigned integer",
		})
	}

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

// @tags Answer
// @summary Update an answer
// @description Update an answer
// @router /answer [post]
// @accept  json
// @param fid query uint true "Form ID"
// @param zjuid query uint true "ZJUID"
// @param eid query uint true "Event ID"
// @param data body model.AnswerRequest_ true "Answer information"
// @success 200
func updateAnswer(c echo.Context) error {
	fid, typeErr := utils.IsUnsignedInteger(c.QueryParam("fid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid need to be an unsigned integer",
		})
	}
	zjuid, typeErr := utils.IsUnsignedInteger(c.QueryParam("zjuid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "zjuid need to be an unsigned integer",
		})
	}
	eid, typeErr := utils.IsUnsignedInteger(c.QueryParam("eid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid need to be an unsigned integer",
		})
	}
	_, itvErr := model.QueryAnswer(fid, strconv.FormatUint(uint64(zjuid), 10), eid)
	answerRequest := model.AnswerRequest{}
	err := c.Bind(&answerRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		if aid, err := model.CreateAnswer(&answerRequest, fid, strconv.FormatUint(uint64(zjuid), 10), eid); err != nil {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "create answer fail",
			})
		} else {
			return c.JSON(http.StatusOK, &utils.Error{
				Code: "SUCCESS",
				Data: aid,
			})
		}
	}
	if err := model.UpdateAnswer(&answerRequest, fid, strconv.FormatUint(uint64(zjuid), 10), eid); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "update Answer fail",
		})
	}
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "update answer success",
	})
}

// @tags Answer
// @summary Get an answer
// @description Get an answer
// @router /answer [get]
// @param fid query uint true "Form ID"
// @param zjuid query uint true "ZJUID"
// @param eid query uint true "Event ID"
// @success 200 {object} model.AnswerResponse_
func getAnswer(c echo.Context) error {
	fid, typeErr := utils.IsUnsignedInteger(c.QueryParam("fid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid need to be an unsigned integer",
		})
	}
	zjuid, typeErr := utils.IsUnsignedInteger(c.QueryParam("zjuid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "zjuid need to be an unsigned integer",
		})
	}
	eid, typeErr := utils.IsUnsignedInteger(c.QueryParam("eid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid need to be an unsigned integer",
		})
	}
	answer, itvErr := model.QueryAnswer(fid, strconv.FormatUint(uint64(zjuid), 10), eid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "answer not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &answer,
	})
}
