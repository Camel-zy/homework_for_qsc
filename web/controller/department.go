package controller

import (
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// @tags Department
// @summary Get department in organization
// @description Get information of a department in a specific organization
// @router /organization/department [get]
// @param oid query uint true "Organization ID"
// @param did query uint true "Department ID"
// @produce json
// @success 200 {object} model.DepartmentApi
func getDepartmentInOrganization(c echo.Context) error {
	var oid, did uint
	err := echo.QueryParamsBinder(c).
		MustUint("oid", &oid).
		MustUint("did", &did).
		BindError()

	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid and did need to be an unsigned integer",
		})
	}

	_, err = model.QueryOrganizationById(oid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "organization not found",
		})
	}

	department, err := model.QueryDepartmentByIdUnderOrganization(oid, did)
	if err != nil {
		return utils.GetApiReturnNotFoundOrInternalError(c, "department", err)
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: department,
	})
}

// @tags Department
// @summary Get all departments in organization
// @description Get brief information of all departments in a specific organization
// @router /organization/department/all [get]
// @param oid query uint true "Organization ID"
// @produce json
// @success 200 {array} model.Brief
func getAllDepartmentInOrganization(c echo.Context) error {
	var oid uint
	err := echo.QueryParamsBinder(c).
		MustUint("oid", &oid).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid need to be an unsigned integer",
		})
	}

	departments, err := model.QueryAllDepartmentInOrganization(oid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "NOT_FOUND",
				Data: "organization not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "get department fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: departments,
	})
}

// @tags Department
// @summary Get all departments in event
// @router /event/department/all [get]
// @param eid query uint true "Event ID"
// @produce json
// @success 200 {array} model.Brief
func getAllDepartmentInEvent(c echo.Context) error {
	var eid uint
	err := echo.QueryParamsBinder(c).
		MustUint("eid", &eid).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid need to be an unsigned integer",
		})
	}

	departments, err := model.QueryAllDepartmentByEid(eid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "NOT_FOUND",
				Data: "event not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "get department fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: departments,
	})

}
