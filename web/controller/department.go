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
// @router /organization/department/{oid}{did} [get]
// @param oid query uint true "Organization ID"
// @param did query uint true "Department ID"
// @produce json
// @success 200 {object} model.DepartmentApi
func getDepartmentInOrganization(c echo.Context) error {
	oid, errOid := utils.IsUnsignedInteger(c.QueryParam("oid"))
	did, errDid := utils.IsUnsignedInteger(c.QueryParam("did"))

	if errOid != nil || errDid != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid and did need to be an unsigned integer",
		})
	}

	_, orgErr := model.QueryOrganizationById(oid)
	if errors.Is(orgErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "organization not found",
		})
	}

	department, depErr := model.QueryDepartmentByIdUnderOrganization(oid, did)
	if errors.Is(depErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "department not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &department,
	})
}

// @tags Department
// @summary Get all departments in organization
// @description Get brief information of all departments in a specific organization
// @router /organization/department/all/{oid} [get]
// @param oid query uint true "Organization ID"
// @produce json
// @success 200 {array} model.Brief
func getAllDepartmentInOrganization(c echo.Context) error {
	oid, typeErr := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid need to be an unsigned integer",
		})
	}

	departments, depErr := model.QueryAllDepartmentInOrganization(oid)
	if errors.Is(depErr, gorm.ErrRecordNotFound){
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "organization not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &departments,
	})
}
