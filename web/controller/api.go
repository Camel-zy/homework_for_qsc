package controller

import (
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/database"
	"git.zjuqsc.com/rop/rop-back-neo/web/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// TODO: handle Internal Server Errors
// TODO: set proper custom error code
func GetUser(c echo.Context) error {
	uid := c.QueryParam("uid")
	if !utils.IsUnsignedInteger(uid) {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: 2, Description: "uid needs to be an unsigned integer"})
	}
	if user, err := database.QueryUserById(uid); errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: 3, Description: err.Error()})
	} else {
		return c.JSON(http.StatusOK, &user)
	}
}

func GetAllUser(c echo.Context) error {
	users, _ := database.QueryAllUser()
	return c.JSON(http.StatusOK, &users)
}

func GetOrganization(c echo.Context) error {
	oid := c.QueryParam("oid")
	if !utils.IsUnsignedInteger(oid) {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: 2, Description: "oid needs to be an unsigned integer"})
	}
	if organization, err := database.QueryOrganizationById(oid); errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: 3, Description: err.Error()})
	} else {
		return c.JSON(http.StatusOK, &organization)
	}
}

func GetAllOrganization(c echo.Context) error {
	organizations, _ := database.QueryAllOrganization()
	return c.JSON(http.StatusOK, &organizations)
}

func GetDepartment(c echo.Context) error {
	did := c.QueryParam("did")
	if !utils.IsUnsignedInteger(did) {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: 2, Description: "did need to be an unsigned integer"})
	}
	if organization, err := database.QueryDepartmentById(did); errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: 3, Description: err.Error()})
	} else {
		return c.JSON(http.StatusOK, &organization)
	}
}

func GetAllDepartment(c echo.Context) error {
	department, _ := database.QueryAllDepartment()
	return c.JSON(http.StatusOK, &department)
}
