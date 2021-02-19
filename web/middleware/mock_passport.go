package middleware

import (
	"bytes"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
)

const qp2glSesstokValid = "MockToken"
const qp2glSesstokSecureValid = "MockSecureToken"

var ePassport *echo.Echo

func MockPassport(mockQscPassportFunction func(c echo.Context) error) {
	/* initialize mocked QSC Passport server */
	ePassport = echo.New()
	ePassport.GET("/passport/get_member_by_token", mockQscPassportFunction)
	requestToQscPassport = func(apiName string, params *url.Values) (resp *http.Response, err error) {
		req := utils.CreateRequest("GET", apiName + params.Encode(), nil)
		resp = utils.CreateResponse(req, ePassport)
		return
	}
	viper.Set("passport.api_name", "/passport/get_member_by_token?")
}

/* A mocked QSC Passport service for go test */
func mockQscPassportService(c echo.Context) error {
	success := &auth{Err: 0, Uid: 1}
	failed := &auth{Err: 1}
	if v := c.QueryParam("token"); v != "" {
		if v == qp2glSesstokValid {
			return c.JSON(http.StatusOK, success)
		} else {
			return c.JSON(http.StatusUnauthorized, failed)
		}
	} else if v := c.QueryParam("token_secure"); v != "" {
		if v == qp2glSesstokSecureValid {
			return c.JSON(http.StatusOK, success)
		} else {
			return c.JSON(http.StatusUnauthorized, failed)
		}
	}
	return c.JSON(http.StatusUnauthorized, failed)
}

/*
The authentication will 100% pass
*/
func MockQscPassportServiceWillPass(c echo.Context) error {
	return c.JSON(http.StatusOK, &auth{Err: 0, Uid: 1})
}

/* configurations for mocking a QSC Passport service */
func mockQscPassportConf() {
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
