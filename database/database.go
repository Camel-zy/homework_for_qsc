package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"git.zjuqsc.com/rop/rop-back-neo/database/model"
)

var DB *gorm.DB

func MakeDB(dbconfig string) {
	DB, err := gorm.Open(postgres.Open(dbconfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if DB == nil {
		panic("db is nil")
	}
}

func Init() {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Organization{},
		&model.Department{},
		&model.JoinedDepartment{},
		&model.Event{},
		&model.Interview{},
		&model.JoinedInterview{}).Error
	if err != nil {
		panic(err)
	}
}
