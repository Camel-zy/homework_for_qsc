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
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: 2, Description: "uid needs to be an uint"})
	}
	if user, err := database.QueryUserById(uid); errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: 3, Description: err.Error()})
	} else {
		return c.JSON(http.StatusOK, &user)
	}
}
