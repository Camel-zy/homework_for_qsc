package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// TODO: put version value into config file
func GetApiVersion(c echo.Context) error {
	return c.String(http.StatusOK, "building")
}
