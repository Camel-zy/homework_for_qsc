package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func MakeDB(dbconfig string) {
	db, err := gorm.Open(postgres.Open(dbconfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
