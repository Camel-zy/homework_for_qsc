package middleware

import (
	"bytes"
	"fmt"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/test"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

var validJwtString string
var testCases []struct{
	name          string
	cookie        string
	isSecureMode  bool
	expectSuccess bool
}

const qp2glSesstokValid = "MockToken"
const qp2glSesstokSecureValid = "MockSecureToken"

const qp2glSesstokInvalid = "MockTokenInvalid"
const qp2glSesstokSecureInvalid = "MockSecureTokenInvalid"

func TestMain(m *testing.M) {
	/* open a sqlite in-memory database */
	model.Connect(sqlite.Open("file::memory:?cache=shared"))
	model.CreateTables()
	test.CreateDatabaseRows()

	viper.Set("passport.enable", true)
	mockPassport(mockQscPassportService)

	/* initialize Viper */
	test.MockJwtConf(600)
	mockQscPassportConf() // although this function has been called in mockPassport(), it still need to be called again

	/* generate a valid JWT string for test */
	rand.Seed(time.Now().Unix())
	uid := uint(rand.Intn(1e5))
	validJwtString, _, _ = utils.GenerateJwt(uid)

	/* this constructor needs to be called after everything has been initialized */
	constructTestCases()

	os.Exit(m.Run())
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
			req := utils.CreateRequest("GET", "/test_authentication", nil)
			req.Header.Set("Cookie", v.cookie)

			resp := utils.CreateResponse(req, e)

			if v.expectSuccess {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			} else {
				assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
			}
		})
	}
}

var ePassport *echo.Echo

func mockPassport(mockQscPassportFunction func(c echo.Context) error) {
	/* initialize mocked QSC Passport server */
	ePassport = echo.New()
	ePassport.GET("/passport/get_member_by_token", mockQscPassportFunction)
	RequestToQscPassport = func(apiName string, params *url.Values) (resp *http.Response, err error) {
		req := utils.CreateRequest("GET", apiName+params.Encode(), nil)
		resp = utils.CreateResponse(req, ePassport)
		return
	}
	viper.Set("passport.api_name", "/passport/get_member_by_token?")
}

/* A mocked QSC Passport service for go test */
func mockQscPassportService(c echo.Context) error {
	success := &PassportAuthResult{Err: 0, Uid: 1}
	failed := &PassportAuthResult{Err: 1}
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
			cookie: fmt.Sprintf(qp2glSesstokName + "=" + qp2glSesstokValid),
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
			cookie: fmt.Sprintf(qp2glSesstokSecureName + "=" + qp2glSesstokSecureValid),
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
			cookie: fmt.Sprintf(qp2glSesstokName + "=" + qp2glSesstokValid),
			expectSuccess: false,
			isSecureMode: true,
		},
		{
			name: "PassportSecureModeError",
			cookie: fmt.Sprintf(qp2glSesstokSecureName + "=" + qp2glSesstokSecureValid),
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
