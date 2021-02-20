package test

import (
	"bytes"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"git.zjuqsc.com/rop/rop-back-neo/web/middleware"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
)

const Qp2glSesstokValid = "MockToken"
const Qp2glSesstokSecureValid = "MockSecureToken"

var ePassport *echo.Echo

func MockPassport(mockQscPassportFunction func(c echo.Context) error) {
	/* initialize mocked QSC Passport server */
	ePassport = echo.New()
	ePassport.GET("/passport/get_member_by_token", mockQscPassportFunction)
	middleware.RequestToQscPassport = func(apiName string, params *url.Values) (resp *http.Response, err error) {
		req := utils.CreateRequest("GET", apiName + params.Encode(), nil)
		resp = utils.CreateResponse(req, ePassport)
		return
	}
	viper.Set("passport.api_name", "/passport/get_member_by_token?")
}

/* A mocked QSC Passport service for go test */
func MockQscPassportService(c echo.Context) error {
	success := &middleware.AuthResult{Err: 0, Uid: 1}
	failed := &middleware.AuthResult{Err: 1}
	if v := c.QueryParam("token"); v != "" {
		if v == Qp2glSesstokValid {
			return c.JSON(http.StatusOK, success)
		} else {
			return c.JSON(http.StatusUnauthorized, failed)
		}
	} else if v := c.QueryParam("token_secure"); v != "" {
		if v == Qp2glSesstokSecureValid {
			return c.JSON(http.StatusOK, success)
		} else {
			return c.JSON(http.StatusUnauthorized, failed)
		}
	}
	return c.JSON(http.StatusUnauthorized, failed)
}

/* configurations for mocking a QSC Passport service */
func MockQscPassportConf() {
	viper.SetConfigType("json")
	var yamlExample = []byte(`
	{
		"passport": {
			"enable": false,
			"is_secure_mode": true,
			"app_id": "NotImportant", 
			"app_secret": "StillNotImportant",
			"api_name": "/passport/get_member_by_token?"
		}
	}
	`)
	_ = viper.ReadConfig(bytes.NewBuffer(yamlExample))
}
