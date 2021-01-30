package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dialector gorm.Dialector) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("DB is nil")
	}
	DB = db
}

func CreateTables() {
	if DB == nil {
		panic("DB is nil")
	}
	err := DB.AutoMigrate(
		&model.User{},
		&model.Organization{},
		&model.Department{},
		&model.JoinedDepartment{},
		&model.Event{},
		&model.Interview{},
		&model.JoinedInterview{},
		&model.Message{})
	if err != nil {
		panic(err)
	}
}
