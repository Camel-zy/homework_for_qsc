package controller

import (
	"errors"
	"fmt"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"net/http"
)

var fileKey = "file"

/*
CAUTIOUS: This function is only a demo
This will be fixed
(RalXYZ)
*/
func setImage(c echo.Context) error {
	/* get file from HTTP request */
	fileHeader, err := c.FormFile(fileKey)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "The image hasn't been set in the valid field",
		})
	}

	/* open file */
	file, err := fileHeader.Open()
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "ERROR_READ_FILE",
			Data: "Error occurs while opening the uploaded file",
		})
	}
	defer file.Close()

	mimeType, _ := fileType(file)
	if mimeType[:5] != "image" {
		return c.JSON(http.StatusUnsupportedMediaType, &utils.Error{
			Code: "INVALID_FILE_TYPE",
			Data: "The type of uploaded file is not a valid MIME image type",
		})
	}

	/*
	We cannot control the filename of the uploaded file,
	since it is submitted by the user.
	When we try to store it into the object storage,
	we have to keep the filename unique.
	This is why we need to generate a UUID v4 string
	and set it as the file's name
	before storing it into the storage.

	Meanwhile, we need to know who owns the avatar,
	so we need to store the user's ID.
	Also, when the front-end tries to retrieve the avatar,
	it's original filename needs to be recovered,
	and this is why we need to store OriginalName.
	This constructs a relation between filenames and ID,
	so these texts will be stored into SQL database.
	 */
	uuidFileName := uuid.NewV4().String() // create a UUID v4 string (RFC 4122)
	image := model.Image{
		OriginalName: fileHeader.Filename,
		CurrentName:  uuidFileName,
		UserID:       c.Get("uid").(uint),
	}
	err = image.Save()  // save the name relation into SQL database
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_STORE_IMAGE",
			Data: "Internal error comes from SQL server",
		})
	}

	err = model.CreateFile(c.Request().Context(), uuidFileName, mimeType, file, fileHeader.Size)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_STORE_MAGE",
			Data: "Internal error comes from object storage server",
		})
	}

	return c.NoContent(http.StatusOK)
}

/*
CAUTIOUS: This function is only a demo
This will be fixed
(RalXYZ)
 */
func getImage(c echo.Context) error {
	image := model.Image{UserID: c.Get("uid").(uint)}
	err := image.GetByUid()  // get data from SQL
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "IMAGE_NOT_FOUND",
			Data: "Image not found",
		})
	} else if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_GET_FILE",
			Data: "Error occurs when retrieving data from SQL",
		})
	}

	file, err := model.GetFile(c.Request().Context(), image.CurrentName)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_GET_FILE",
			Data: "Error occurs when retrieving data from object storage",
		})
	}

	/* check if the image has not been found */
	mimeType, err := fileType(file)
	if t, ok := err.(minio.ErrorResponse); ok {
		if t.StatusCode == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "IMAGE_NOT_FOUND",
				Data: "Image not found",
			})
		} else {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "ERROR_GET_FILE",
				Data: "Error occurs when retrieving data from object storage",
			})
		}
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_READ_FILE",
			Data: "Error occurs when trying to read the retrieved file",
		})
	}

	c.Response().Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, image.OriginalName))
	c.Response().Header().Set("Cache-Control", "public; max-age=259200")  // 3 months
	// the browser need to cache this for a while, because this is the avatar

	return c.Stream(http.StatusOK, mimeType, file)
}

/*
check data type of the uploaded file
http.DetectContentType() will only process
the first 512 bytes of the []byte parameter
so we only need to read 512 bytes
*/
func fileType(file io.Reader) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	return http.DetectContentType(buffer), err
}
