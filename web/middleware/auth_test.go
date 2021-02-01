package middleware

import (
	"bytes"
	"fmt"
	"git.zjuqsc.com/rop/rop-back-neo/test"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

var ePassport *echo.Echo

var validJwtString string
var testCases []struct{
	name          string
	cookie        string
	isSecureMode  bool
	expectSuccess bool
}

const qp2glSesstokName = "qp2gl_sesstok"
const qp2glSesstokSecureName = "qp2gl_sesstok_secure"

const qp2glSesstok = "MockToken"
const qp2glSesstokSecure = "MockSecureToken"
const qp2glSesstokInvalid = "MockTokenInvalid"
const qp2glSesstokSecureInvalid = "MockSecureTokenInvalid"

func TestMain(m *testing.M) {
	/* initialize mocked QSC Passport server */
	ePassport = echo.New()
	ePassport.GET("/passport/get_member_by_token", mockQscPassport)
	requestToQscPassport = func(apiName string, params *url.Values) (resp *http.Response, err error) {
		req := test.CreateRequest("GET", apiName + params.Encode(), nil)
		resp = test.CreateResponse(req, ePassport)
		return
	}

	/* initialize Viper */
	test.MockJwtConf(600)
	mockQscPassportConf()

	/* generate a valid JWT string for test */
	rand.Seed(time.Now().Unix())
	uid := uint(rand.Intn(1e5))
	validJwtString, _ = utils.GenerateJWT(uid)

	/* this constructor needs to be called after everything has been initialized */
	constructTestCases()

	os.Exit(m.Run())
}

/* A mocked QSC Passport service */
func mockQscPassport(c echo.Context) error {
	success := &auth{Err: 0, Uid: 2333}
	failed := &auth{Err: 1}
	if v := c.QueryParam("token"); v != "" {
		if v == qp2glSesstok {
			return c.JSON(http.StatusOK, success)
		} else {
			return c.JSON(http.StatusUnauthorized, failed)
		}
	} else if v := c.QueryParam("token_secure"); v != "" {
		if v == qp2glSesstokSecure {
			return c.JSON(http.StatusOK, success)
		} else {
			return c.JSON(http.StatusUnauthorized, failed)
		}
	}
	return c.JSON(http.StatusUnauthorized, failed)
}

/* configurations for mocking a QSC Passport service */
func mockQscPassportConf() {
	viper.SetConfigType("json")
	var yamlExample = []byte(`
	{
		"passport": {
			"app_id": "NotImportant", 
			"app_secret": "StillNotImportant",
			"api_name": "/passport/get_member_by_token?"
		}
	}
	`)
	_ = viper.ReadConfig(bytes.NewBuffer(yamlExample))
}

func TestMiddleware(t *testing.T) {
	e := echo.New()
	e.Use(Auth)
	e.GET("/test_authentication", end)

	// t.Parallel()
	for _, v := range testCases {
		// v := v
		t.Run(v.name, func(t *testing.T) {
			// t.Parallel()
			viper.Set("passport.is_secure_mode", v.isSecureMode)  // change secure mode based on testing case
			req := test.CreateRequest("GET", "/test_authentication", nil)
			req.Header.Set("Cookie", v.cookie)

			resp := test.CreateResponse(req, e)

			if v.expectSuccess {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			} else {
				assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
			}
		})
	}
}

/*
This constructor needs to be called
after everything has been initialized.
It initializes the test cases.
*/
func constructTestCases() {
	testCases = []struct{
		name          string
		cookie        string
		isSecureMode  bool
		expectSuccess bool
	} {
		{
			name: "NoCookieIsSet",
			cookie: "",
			expectSuccess: false,
		},
		{
			name: "RopJwtValid",
			cookie: fmt.Sprintf(jwtName + "=" + validJwtString),
			expectSuccess: true,
		},
		{
			name: "RopJwtInvalid",
			cookie: fmt.Sprintf(jwtName + "=" + "InvalidJwtString"),
			expectSuccess: false,
		},
		{
			name: "PassportCookieValid",
			cookie: fmt.Sprintf(qp2glSesstokName + "=" + qp2glSesstok),
			expectSuccess: true,
			isSecureMode: false,
		},
		{
			name: "PassportCookieInvalid",
			cookie: fmt.Sprintf(qp2glSesstokName + "=" + qp2glSesstokInvalid),
			expectSuccess: false,
			isSecureMode: false,
		},
		{
			name: "PassportSecureCookieValid",
			cookie: fmt.Sprintf(qp2glSesstokSecureName + "=" + qp2glSesstokSecure),
			expectSuccess: true,
			isSecureMode: true,
		},
		{
			name: "PassportSecureCookieInvalid",
			cookie: fmt.Sprintf(qp2glSesstokSecureName + "=" + qp2glSesstokSecureInvalid),
			expectSuccess: false,
			isSecureMode: true,
		},
		{
			name: "PassportSecureModeError",
			cookie: fmt.Sprintf(qp2glSesstokName + "=" + qp2glSesstok),
			expectSuccess: false,
			isSecureMode: true,
		},
		{
			name: "PassportSecureModeError",
			cookie: fmt.Sprintf(qp2glSesstokSecureName + "=" + qp2glSesstokSecure),
			expectSuccess: false,
			isSecureMode: false,
		},
	}
}

/*
A simple endpoint for mocked controller
 */
func end(c echo.Context) error {
	return nil
}
