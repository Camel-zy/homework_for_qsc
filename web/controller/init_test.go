package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/database"
	"git.zjuqsc.com/rop/rop-back-neo/test"
	"gorm.io/driver/sqlite"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	/* open a sqlite in-memory database */
	database.Connect(sqlite.Open("file::memory:?cache=shared"))
	database.CreateTables()

	test.CreateDatabaseRows()
	InitWebFramework(true)

	os.Exit(m.Run())
}
