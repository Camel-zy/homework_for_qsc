package auth

import (
	"encoding/json"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
)

type auth struct {
	Err  int   `json:"err"`
	Uid  uint  `json:"uid"`
}

// TODO: do not request for authentication every time
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isSecureMode := viper.GetBool("passport.is_secure_mode")

		/* get cookies */
		var cookieName, tokenUrlParamName string
		if isSecureMode {
			cookieName = "qp2gl_sesstok_secure"
			tokenUrlParamName = "token_secure"
		} else {
			cookieName = "qp2gl_sesstok"
			tokenUrlParamName = "token"
		}
		cookie, getCookieErr := c.Cookie(cookieName)
		if getCookieErr != nil {
			return c.JSON(http.StatusUnauthorized, &utils.Error{Code: "COOKIE_NOT_FOUND", Data: "qsc passport cookie is required"})
		}

		authQSCPassportError := authByQSCPassport(c, cookie, tokenUrlParamName)
		if authQSCPassportError != nil {
			return authQSCPassportError
		}

		return next(c)
	}
}

/*
This function checks the validity of QSC Passport cookie.
It sends a request to QSC Passport authentication server,
and check the response. If the user is authorized, return nil.
Else, return an echo JSON response.
 */
func authByQSCPassport(c echo.Context, cookie *http.Cookie, tokenUrlParamName string) error {
	appID := viper.GetString("passport.app_id")
	appSecret := viper.GetString("passport.app_secret")
	apiName := viper.GetString("passport.api_name")

	/* generate url parameter string */
	params := url.Values{}
	params.Add("appid", appID)
	params.Add("appsecret", appSecret)
	params.Add(tokenUrlParamName, cookie.Value)

	/* create a request */
	req, _ := http.NewRequest("GET", apiName + params.Encode(), nil)

	/* send a request and get the response */
	client := &http.Client{}
	resp, getErr := client.Do(req)
	if getErr != nil{
		return c.JSON(http.StatusServiceUnavailable, &utils.Error{Code: "AUTH_SERVICE_ERROR", Data: "error occurs when sending request to auth service"})
	}
	defer resp.Body.Close()

	/* read the body of the response */
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		panic(readErr)
	}

	/* get the value of key "err" from the JSON response */
	authResult := auth{}
	jsonErr := json.Unmarshal(body, &authResult)
	if jsonErr != nil {
		panic(jsonErr)
	}

	/* the request can be authorized IF AND ONLY IF error code is 0 */
	if authResult.Err != 0 {
		return c.JSON(http.StatusUnauthorized, &utils.Error{Code: "AUTH_FAILED", Data: "auth failed according to the response of QSC Passport auth service"})
	}

	return nil
}
