package controller

import (
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func getUser(c echo.Context) error {
	uid, typeErr := utils.IsUnsignedInteger(c.QueryParam("uid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "uid need to be an unsigned integer",
		})
	}

	user, usrErr := model.QueryUserById(uid);
	if  errors.Is(usrErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "user not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &user,
	})
}

func getAllUser(c echo.Context) error {
	users, _ := model.QueryAllUser()
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &users,
	})
}
