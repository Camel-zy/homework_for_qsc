package controller

import (
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/database"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// TODO: handle Internal Server Errors
func GetUser(c echo.Context) error {
	uid, typeErr := utils.IsUnsignedInteger(c.QueryParam("uid"));
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: "BAD_REQUEST", Data: "uid needs to be an unsigned integer"})
	}

	user, usrErr := database.QueryUserById(uid);
	if  errors.Is(usrErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "USR_NOT_FOUND", Data: "user not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &user})
}

func GetAllUser(c echo.Context) error {
	users, _ := database.QueryAllUser()
	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &users})
}

func GetOrganization(c echo.Context) error {
	oid, typeErr := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: "BAD_REQUEST", Data: "oid needs to be an unsigned integer"})
	}

	organization, orgErr := database.QueryOrganizationById(oid);
	if errors.Is(orgErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "ORG_NOT_FOUND", Data: "organization not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &organization})
}

func GetAllOrganization(c echo.Context) error {
	organizations, _ := database.QueryAllOrganization()
	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &organizations})
}

func GetDepartmentUnderOrganization(c echo.Context) error {
	oid, errOid := utils.IsUnsignedInteger(c.QueryParam("oid"))
	did, errDid := utils.IsUnsignedInteger(c.QueryParam("did"))

	if errOid != nil || errDid != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: "BAD_REQUEST", Data: "oid and did need to be an unsigned integer"})
	}

	_, orgErr := database.QueryOrganizationById(oid)
	if errors.Is(orgErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "ORG_NOT_FOUND", Data: "organization not found"})
	}

	department, depErr := database.QueryDepartmentByIdUnderOrganization(oid, did)
	if errors.Is(depErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "DEP_NOT_FOUND", Data: "department not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &department})
}

func GetAllDepartmentUnderOrganization(c echo.Context) error {
	oid, typeErr := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: "BAD_REQUEST", Data: "oid needs to be an unsigned integer"})
	}

	departments, depErr := database.QueryAllDepartmentUnderOrganization(oid)
	if errors.Is(depErr, gorm.ErrRecordNotFound){
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "ORG_NOT_FOUND", Data: "organization not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &departments})
}
