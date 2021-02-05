package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/test"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	/* open a sqlite in-memory database */
	model.Connect(sqlite.Open("file::memory:?cache=shared"))
	model.CreateTables()

	test.CreateDatabaseRows()

	// set "rop.test" true to skip authentication
	viper.Set("rop.test", true)
	InitWebFramework()

	os.Exit(m.Run())
}
