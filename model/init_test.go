package model

import (
	"os"
	"testing"

	"git.zjuqsc.com/rop/rop-back-neo/conf"
	"gorm.io/driver/sqlite"
)

func TestMain(m *testing.M) {
	/* open a sqlite in-memory database */
	conf.Init()
	Connect(sqlite.Open("file::memory:?cache=shared"))
	// Connect(postgres.Open(conf.GetDatabaseLoginInfo()))
	CreateTables()
	ConnectObjectStorage()

	// set testController true to skip authentication fully
	os.Exit(m.Run())

}
