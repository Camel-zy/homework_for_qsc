package auth

import (
	"encoding/json"
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

// TODO: add error message and error code
// TODO: do not request for authentication every time
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isSecureMode := viper.GetBool("passport.is_secure_mode")
		appID := viper.GetString("passport.app_id")
		appSecret := viper.GetString("passport.app_secret")
		apiName := viper.GetString("passport.api_name")

		/* get cookies */
		var cookieName , tokenUrlParamName string
		if isSecureMode {
			cookieName = "qp2gl_sesstok_secure"
			tokenUrlParamName = "token_secure"
		} else {
			cookieName = "qp2gl_sesstok"
			tokenUrlParamName = "token"
		}
		cookie, getCookieErr := c.Cookie(cookieName)
		if getCookieErr != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

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
			panic(getErr)
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
			return c.JSON(http.StatusUnauthorized, nil)
		}

		return next(c)
	}
}
