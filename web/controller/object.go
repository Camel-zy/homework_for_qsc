package controller

import (
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type PresignedPost struct {
	Url    string            `json:"url"`
	UUID   uuid.UUID         `json:"uuid"`
	Policy map[string]string `json:"policy"`
}

// @tags Object
// @produce json
func createObject(c echo.Context) error {
	url, formData, uuid, err := model.CreateObject(c.Request().Context(), c.QueryParam("name"))
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_OBJECT_STORE",
			Data: "internal error of object storage",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: PresignedPost{
			Url:    url.String(),
			Policy: formData,
			UUID:   uuid,
		},
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
