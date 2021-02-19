package controller

import (
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// @tags Organization
// @summary Get information of organization
// @description Get information of a specific organization
// @router /organization{oid} [get]
// @param oid query uint true "Organization ID"
// @produce json
// @success 200 {object} model.OrganizationApi
func getOrganization(c echo.Context) error {
	oid, typeErr := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "oid need to be an unsigned integer"},
			)
	}

	organization, orgErr := model.QueryOrganizationById(oid)
	if errors.Is(orgErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "organization not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &organization,
	})
}

// @tags Organization
// @summary Get all organizations
// @description Get brief information of all organizations
// @router /organization/all [get]
// @produce json
// @success 200 {array} model.Brief
func getAllOrganization(c echo.Context) error {
	organizations, _ := model.QueryAllOrganization(c.Get("uid").(uint))
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &organizations,
	})
}
