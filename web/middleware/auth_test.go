package middleware

import (
	"fmt"
	"git.zjuqsc.com/rop/rop-back-neo/test"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
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

const qp2glSesstokInvalid = "MockTokenInvalid"
const qp2glSesstokSecureInvalid = "MockSecureTokenInvalid"

func TestMain(m *testing.M) {
	MockPassport()

	/* initialize Viper */
	test.MockJwtConf(600)
	mockQscPassportConf()  // although this function has been called in MockPassport(), it still need to be called again

	/* generate a valid JWT string for test */
	rand.Seed(time.Now().Unix())
	uid := uint(rand.Intn(1e5))
	validJwtString, _, _ = utils.GenerateJWT(uid)

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
