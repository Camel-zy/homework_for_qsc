package controller

import (
	"fmt"
	"git.zjuqsc.com/rop/rop-back-neo/database"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"io"
	"net/http"
)

var fileKey = "file"


/*
CAUTIOUS: This function is only a demo
This will be fixed on Feb 5th
(RalXYZ)
*/
func setImage(c echo.Context) error {
	/* get file from HTTP request */
	fileHeader, err := c.FormFile(fileKey)
	if err != nil {
		panic(err)
	}

	/* open file */
	file, err := fileHeader.Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mimeType, _ := fileType(file)
	if mimeType[:5] != "image" {
		return c.JSON(http.StatusUnsupportedMediaType, &utils.Error{
			Code: "INVALID_FILE_TYPE",
			Data: "The type of uploaded file is not a valid MIME image type",
		})
	}

	// FIXME: set file name and bucket name properly
	err = database.CreateFile(c.Request().Context(), fileHeader.Filename, mimeType, file, fileHeader.Size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "ERROR_STORE_FILE",
			Data: "Error comes from object storage server",
		})
	}

	return c.NoContent(http.StatusOK)
}

/*
CAUTIOUS: This function is only a demo
It can only retrieve a file named "test.jpg" currently.
This will be fixed on Feb 5th
(RalXYZ)
 */
func getImage(c echo.Context) error {
	fileName := "test.jpg"  // FIXME: get this from psql
	file, err := database.GetFile(c.Request().Context(), fileName)
	if err != nil {
		panic(err)
	}

	/* check if the image has not been found */
	mimeType, err := fileType(file)
	if t, ok := err.(minio.ErrorResponse); ok {
		if t.StatusCode == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "FILE_NOT_FOUND",
				Data: "File not found",
			})
		} else {
			panic(t)
		}
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	c.Response().Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, fileName))
	c.Response().Header().Set("Cache-Control", "public; max-age=259200")  // 3 months

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
