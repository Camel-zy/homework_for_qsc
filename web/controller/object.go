package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

// @tags Object
// @produce json
func createObject(c echo.Context) error {
	url, err := model.CreateObject(c.Request().Context(), c.QueryParam("name"))
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_OBJECT_STORE",
			Data: "internal error of object storage",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: url.String(),
	})
}

// @tags Object
// @produce json
func sealObject(c echo.Context) error {
	uuid, err := uuid.FromString(c.QueryParam("uuid"))
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_OBJECT_STORE",
			Data: "invalid uuid",
		})
	}
	err = model.SealObject(c.Request().Context(), uuid)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_OBJECT_STORE",
			Data: "error occurs when retrieving object url from object storage",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
	})
}

// @tags Object
// @produce json
func getObject(c echo.Context) error {
	uuid, err := uuid.FromString(c.QueryParam("uuid"))
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_OBJECT_STORE",
			Data: "invalid uuid",
		})
	}
	url, err := model.GetObject(c.Request().Context(), uuid)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_OBJECT_STORE",
			Data: "error occurs when retrieving object url from object storage",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: url.String(),
	})
}
