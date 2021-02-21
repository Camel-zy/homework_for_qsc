package middleware

import (
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

func CheckMinioStatus(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !viper.GetBool("minio.enable") {
			return c.JSON(http.StatusServiceUnavailable,  &utils.Error{
				Code: "SERVICE_UNAVAILABLE",
				Data: "The connection to MinIO service has been closed by the administrator",
			})
		}
		return next(c)
	}
}
