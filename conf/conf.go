package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

/*
 */
func InitConf() {
	viper.SetConfigName("conf")  // set the config file name. Viper will automatically detect the file extension name
	viper.AddConfigPath("./conf")     // search the config file under the current directory
	// viper.AddConfigPath("foo")  // you can search this config file under multiple directories

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error while reading config file: %s \n", err))
	}

	fmt.Println("Configuration file loaded.")

	loginKeys := []string{"user", "password", "host", "port", "db_name"}
	passportKeys := []string{"is_secure_mode", "app_id", "app_secret", "api_name"}

	/* check the existence of the required values */
	getKeyErrExists := false
	getKeyErrExists = checkConfIsSet("login", loginKeys) || getKeyErrExists
	getKeyErrExists = checkConfIsSet("passport", passportKeys) || getKeyErrExists
	if getKeyErrExists {
		panic("Error occurs wile getting keys, please check your config file.")
	}

	fmt.Println("Configuration file checking succeeded. All required values are set.")
}

func checkConfIsSet(name string, keys []string) (getKeyErrExists bool) {
	getKeyErrExists = false
	for i := range keys {
		if !viper.IsSet(name + "." + keys[i]) {
			fmt.Printf("\"%s\" not set", keys[i])
			fmt.Printf("in\"" + name + "\"\n")
			getKeyErrExists = true
		}
	}
	return
}

func GetDatabaseLoginInfo() string {
	if !viper.IsSet("login") {
		panic("\"login\" not set in config file.")
	}
	loginInfo := viper.GetStringMapString("login")

	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		loginInfo["user"],
		loginInfo["password"],
		loginInfo["host"],
		loginInfo["port"],
		loginInfo["db_name"])
}
