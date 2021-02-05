package controller

import (
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
)

func getApiVersion(c echo.Context) error {
	return c.String(http.StatusOK, viper.GetString("rop.api_version"))
}

// TODO: handle Internal Server Errors
func getUser(c echo.Context) error {
	uid, typeErr := utils.IsUnsignedInteger(c.QueryParam("uid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: "BAD_REQUEST", Data: "uid need to be an unsigned integer"})
	}

	user, usrErr := model.QueryUserById(uid);
	if  errors.Is(usrErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "USR_NOT_FOUND", Data: "user not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &user})
}

func getAllUser(c echo.Context) error {
	users, _ := model.QueryAllUser()
	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &users})
}

func getOrganization(c echo.Context) error {
	oid, typeErr := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: "BAD_REQUEST", Data: "oid need to be an unsigned integer"})
	}

	organization, orgErr := model.QueryOrganizationById(oid)
	if errors.Is(orgErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "ORG_NOT_FOUND", Data: "organization not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &organization})
}

func getAllOrganization(c echo.Context) error {
	organizations, _ := model.QueryAllOrganization()
	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &organizations})
}

func getDepartmentUnderOrganization(c echo.Context) error {
	oid, errOid := utils.IsUnsignedInteger(c.QueryParam("oid"))
	did, errDid := utils.IsUnsignedInteger(c.QueryParam("did"))

	if errOid != nil || errDid != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: "BAD_REQUEST", Data: "oid and did need to be an unsigned integer"})
	}

	_, orgErr := model.QueryOrganizationById(oid)
	if errors.Is(orgErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "ORG_NOT_FOUND", Data: "organization not found"})
	}

	department, depErr := model.QueryDepartmentByIdUnderOrganization(oid, did)
	if errors.Is(depErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "DEP_NOT_FOUND", Data: "department not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &department})
}

func getAllDepartmentUnderOrganization(c echo.Context) error {
	oid, typeErr := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: "BAD_REQUEST", Data: "oid need to be an unsigned integer"})
	}

	departments, depErr := model.QueryAllDepartmentUnderOrganization(oid)
	if errors.Is(depErr, gorm.ErrRecordNotFound){
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "ORG_NOT_FOUND", Data: "organization not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &departments})
}
