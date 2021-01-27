package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

func GetApiVersion(c echo.Context) error {
	return c.String(http.StatusOK, viper.GetString("rop.api_version"))
}
