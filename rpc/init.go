package rpc

import (
	"errors"
	"time"

	// SMSService "git.zjuqsc.com/rop/rop-sms/gRPC"
	SMSService "gRPC"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Sms server
var ErrSMSUninitialized = errors.New("sms service is not initialized")

var enable bool
var app_id uint
var app_key string
var timeout time.Duration

func Init() {
	enable = viper.GetBool("rpc.enable")
	if !enable {
		return
	}
	app_id = viper.GetUint("rpc.app_id")
	app_key = viper.GetString("rpc.app_key")
	timeout = time.Millisecond * time.Duration(viper.GetUint("rpc.timeout"))

	Sms.Init()

	// test
	_, pongErr := Sms.Ping(&SMSService.PingReq{})
	if pongErr != nil {
		panic(pongErr)
	}
	logrus.Printf("SMS gRPC server connected")
}
