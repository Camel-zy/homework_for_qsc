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
	UUID   string            `json:"uuid"`
	Policy map[string]string `json:"policy"`
}

// @tags Object
// @router /object/create [post]
// @description Get a URL and necessary policies for object uploading
// @description Send a POST request to the URL given by the response of this API,
// @description while you need to set the given policies into the request headers respectively
// @param name query string true "Object name"
// @produce json
// @success 200 {object} PresignedPost
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
			UUID:   uuid.String(),
		},
	})
}

// @tags Object
// @router /object/seal [post]
// @description Mark an object as successfully uploaded, ready for downloading
// @description You need to send a request to this API after you've successfully uploaded an object
// @param uuid query string true "Object UUID"
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
// @router /object/get [get]
// @description Get a url for object downloading
// @param uuid query string true "Object UUID"
// @produce json
// @success 200 {string} string URL
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
