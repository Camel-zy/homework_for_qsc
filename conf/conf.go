package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("conf")  // set the config file name. Viper will automatically detect the file extension name
	viper.AddConfigPath("/etc/rop-neo")  // search `/etc`
	viper.AddConfigPath("./")     // search the config file under the current directory

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	logrus.Info("Configuration file loaded")

	var confItems = map[string][]string {
		"rop": {"api_version", "allow_origins"},
		"sql": {"user", "password", "host", "port", "db_name"},
		"minio": {"enable", "endpoint", "id", "secret", "secure", "bucket_name"},
		"passport": {"enable", "is_secure_mode", "app_id", "app_secret", "api_name"},
		"jwt": {"issuer", "max_age", "secret_key"},
	}

	for k, v := range confItems {
		checkConfIsSet(k, v)
	}

	logrus.Info("All required values in configuration file are set")
}

func checkConfIsSet(name string, keys []string) {
	for i := range keys {
		wholeKey := name + "." + keys[i]
		if !viper.IsSet(wholeKey) {
			logrus.WithField(wholeKey, nil).
				Fatal("The following item of your configuration file hasn't been set properly: ")
		}
	}
	return
}

func GetDatabaseLoginInfo() string {
	loginInfo := viper.GetStringMapString("sql")

	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		loginInfo["user"],
		loginInfo["password"],
		loginInfo["host"],
		loginInfo["port"],
		loginInfo["db_name"])
}
