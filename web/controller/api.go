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
	if uid, err := utils.IsUnsignedInteger(c.QueryParam("uid")); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: 2, Description: "uid needs to be an unsigned integer"})
	} else if user, err := database.QueryUserById(uid); errors.Is(err, gorm.ErrRecordNotFound) {
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
	if oid, err := utils.IsUnsignedInteger(c.QueryParam("oid")); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: 2, Description: "oid needs to be an unsigned integer"})
	} else if organization, err := database.QueryOrganizationById(oid); errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: 3, Description: err.Error()})
	} else {
		return c.JSON(http.StatusOK, &organization)
	}
}

func GetAllOrganization(c echo.Context) error {
	organizations, _ := database.QueryAllOrganization()
	return c.JSON(http.StatusOK, &organizations)
}

func GetDepartmentUnderOrganization(c echo.Context) error {
	oid, errOid := utils.IsUnsignedInteger(c.QueryParam("oid"))
	did, errDid := utils.IsUnsignedInteger(c.QueryParam("did"))
	if errOid != nil || errDid != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: 2, Description: "oid and did need to be an unsigned integer"})
	}
	if department, err := database.QueryDepartmentByIdUnderOrganization(oid, did); errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: 3, Description: err.Error()})
	} else {
		return c.JSON(http.StatusOK, &department)
	}
}

func GetAllDepartmentUnderOrganization(c echo.Context) error {
	if oid, err := utils.IsUnsignedInteger(c.QueryParam("oid")); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: 2, Description: "oid needs to be an unsigned integer"})
	} else if departments, err := database.QueryAllDepartmentUnderOrganization(oid); errors.Is(err, gorm.ErrRecordNotFound){
		return c.JSON(http.StatusNotFound, &utils.Error{Code: 3, Description: "organization " + err.Error()})
	} else {
		return c.JSON(http.StatusOK, &departments)
	}
}
