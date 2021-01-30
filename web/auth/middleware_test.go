package auth

import (
	"fmt"
	"git.zjuqsc.com/rop/rop-back-neo/web/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func end(c echo.Context) error {
	return nil
}

// TODO: test QSC Passport authentication
func TestMiddleware(t *testing.T) {
	e := echo.New()
	e.GET("/test_authentication", end)

	rand.Seed(time.Now().Unix())
	uid := uint(rand.Intn(1e5))
	jwtString, _ := generateJWT(uid)

	req := utils.CreateRequest("GET", "/test_authentication", nil)
	req.Header.Set("Cookie", fmt.Sprintf(jwtName + "=" + jwtString))

	resp := utils.CreateResponse(req, e)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
