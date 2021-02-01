package test

import (
	"bytes"
	"github.com/spf13/viper"
)

func MockJwtConf(maxAge int)  {
	viper.SetConfigType("json")
	var yamlExample = []byte(`
	{
		"jwt": {
			"issuer": "rop", 
			"secret_key": "MockSecretKey"
		}
	}
	`)
	_ = viper.ReadConfig(bytes.NewBuffer(yamlExample))
	viper.Set("jwt.max_age", maxAge)
}