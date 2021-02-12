package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/test"
	"gorm.io/driver/sqlite"
	"os"
	"testing"
)

type req struct {
	urlPath    string
	urlQuery   string
}
type resp struct {
	statusCode int
	jsonStruct interface{}  // TODO: maybe we need to change the type of this
}

type testCase struct {
	name string
	req  req
	resp resp
}

func TestMain(m *testing.M) {
	/* open a sqlite in-memory database */
	model.Connect(sqlite.Open("file::memory:?cache=shared"))
	model.CreateTables()

	test.CreateDatabaseRows()

	// set testController true to skip authentication fully
	testController = true
	InitWebFramework()

	os.Exit(m.Run())
}
