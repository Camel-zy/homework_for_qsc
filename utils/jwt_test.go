package utils

import (
	"git.zjuqsc.com/rop/rop-back-neo/test"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// Test if the JWT can be correctly parsed after generated
func TestJWT(t *testing.T) {
	test.MockJwtConf(600)

	rand.Seed(time.Now().Unix())
	uid := uint(rand.Intn(1e5))

	jwtString, _, _ := GenerateJwt(uid)
	jwtToken, err := ParseJwt(jwtString)
	assert.Nil(t, err)
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, strconv.Itoa(int(uid)), claims["sub"])
}