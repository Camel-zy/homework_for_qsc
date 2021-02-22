package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetApiVersion (t *testing.T) {
	mockVersion := "0.1.0"
	viper.Set("rop.api_version", mockVersion)

	req := utils.CreateRequest("GET", "/api/version", nil)
	resp := utils.CreateResponse(req, e)
	bodyByteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, mockVersion, string(bodyByteArr))
}
