package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("conf")  // set the config file name. Viper will automatically detect the file extension name
	viper.AddConfigPath("./conf")     // search the config file under the current directory
	// viper.AddConfigPath("foo")  // you can search this config file under multiple directories

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error while reading config file: %s \n", err))
	}

	fmt.Println("Configuration file loaded.")

	var confItems = map[string][]string {
		"rop": {"api_version"},
		"sql": {"user", "password", "host", "port", "db_name"},
		"passport": {"is_secure_mode", "app_id", "app_secret", "api_name"},
		"jwt": {"issuer", "max_age", "secret_key"},
	}

	for k, v := range confItems {
		err := checkConfIsSet(k, v)
		if err {
			panic(fmt.Sprintf("\"%s\" item of your config file hasn't been set properly. \nPlease check your config file.", k))
		}
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
	loginInfo := viper.GetStringMapString("sql")

	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		loginInfo["user"],
		loginInfo["password"],
		loginInfo["host"],
		loginInfo["port"],
		loginInfo["db_name"])
}
