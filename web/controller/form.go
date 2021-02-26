package controller

import (
	"errors"
	"net/http"

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
	var oid, did uint
	err := echo.QueryParamsBinder(c).
		MustUint("eid", &oid).
		MustUint("did", &did).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid and did need to be an unsigned integer",
		})
	}

	formRequest := model.FormRequest{}
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

	if fid, err := model.CreateForm(&formRequest, oid, did); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create form fail",
		})
	} else {
		return c.JSON(http.StatusOK, &utils.Error{
			Code: "SUCCESS",
			Data: fid,
		})
	}
}

// @tags Form
// @summary Update a form
// @description Update a form
// @router /form [post]
// @accept  json
// @param data body model.FormApi_ true "Form information"
// @success 200
func updateForm(c echo.Context) error {
	var fid uint
	err := echo.QueryParamsBinder(c).MustUint("fid", &fid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid needs to be an unsigned integer",
		})
	}

	formRequest := model.FormRequest{}
	err = c.Bind(&formRequest)
	if err != nil {
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
