package auth

import (
	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// Test if the JWT can be correctly parsed after generated
func TestJWT(t *testing.T) {
	/* init configuration */
	viper.SetConfigType("json")
	var yamlExample = []byte(`
	{
		"jwt": {
			"issuer": "rop", 
			"max_age": 600, 
			"secret_key": "AllYourBase"
		}
	}
	`)
	_ = viper.ReadConfig(bytes.NewBuffer(yamlExample))

	rand.Seed(time.Now().Unix())
	uid := uint(rand.Intn(1e5))

	jwtString, _ := generateJWT(uid)
	jwtToken, err := parseJWT(jwtString)
	assert.Nil(t, err)
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, strconv.Itoa(int(uid)), claims["sub"])
}
