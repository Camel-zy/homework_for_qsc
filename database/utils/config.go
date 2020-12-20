package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Login struct {
	DbUser   string `json:"user"`
	DbPasswd string `json:"passwd"`
	DbHost   string `json:"host"`
	DbPort   string `json:"port"`
	DbName   string `json:"dbName"`
}

func GetLoginInfo(file string) *Login {
	var login Login
	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(byteValue), &login)
	if err != nil {
		panic(err)
	}

	return &login
}

func ParseLoginInfo(config *Login) string {
	dbConfig := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DbUser,
		config.DbPasswd,
		config.DbHost,
		config.DbPort,
		config.DbName)
	return dbConfig
}
