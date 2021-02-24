package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// @tags Avatar
// @summary Get the information of uploading avatar
// @description You will receive multiple fields in "Data" field if the request status is 200 OK
// @description One of these fields is "url", you are supposed to send a request to this URL
// @description Other fields need to be set into the multipart/form-data of the request mentioned above
// @description Meanwhile, you are supposed to set "Content-Type" field according to the MIME type of the avatar
// @description Finally, set "file" field with the avatar you are going to upload
// @router /avatar [post]
// @produce json
func setAvatar(c echo.Context) error {
	url, formData, err := model.CreateFile(c.Request().Context(),
		`avatar/` + strconv.FormatUint(uint64(c.Get("uid").(uint)), 10),
		nil)

	formData["url"] = url.String()

	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_STORE_IMAGE",
			Data: "internal error comes from object storage server",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: formData,
	})
}

// @tags Avatar
// @summary Get a URL from which you can get the avatar
// @description If everything is fine, the Data field will be a URL, from which you can get the avatar
// @router /avatar [get]
// @produce json
func getAvatar(c echo.Context) error {
	objectName := `avatar/` + strconv.FormatUint(uint64(c.Get("uid").(uint)), 10)

	url, err := model.GetFile(c.Request().Context(), objectName)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_GET_FILE",
			Data: "error occurs when retrieving data from object storage",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: url.String(),
	})
}
