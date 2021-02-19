package model

import (
	"git.zjuqsc.com/rop/rop-back-neo/datatype"
	"time"
)

type Form struct {
	ID             uint          `gorm:"not null;autoIncrement;primaryKey"`
	Name           string        `gorm:"size:40;not null"`
	CreateTime     time.Time     `gorm:"size:30;not null"`
	UserID         uint          `gorm:"not null"`
	User           User          // FOREIGN KEY (UserID) REFERENCES User(UserID)
	OrganizationID uint          `gorm:"not null"`
	DepartmentID   uint          `gorm:"not null"`
	Status         uint          `gorm:"not null"`
	Content        datatype.JSON `gorm:"type:datatype;not null"`
}

type FormCreateRequest struct {
	//ID             uint      `gorm:"not null;autoIncrement;primaryKey"`
	Name           string        `gorm:"size:40;not null"`
	CreateTime     time.Time     `gorm:"size:30;not null"`
	UserID         uint          `gorm:"not null"`
	User           User          // FOREIGN KEY (UserID) REFERENCES User(UserID)
	OrganizationID uint          `gorm:"not null"`
	DepartmentID   uint          `gorm:"not null"`
	Status         uint          `gorm:"not null"`
	Content        datatype.JSON `gorm:"type:datatype;not null"`
}

func CreateForm(requestForm *Form) error {
	result := gormDb.Create(requestForm)
	return result.Error
}

func QueryFormById(id uint) (*Form, error) {
	var dbForm Form
	result := gormDb.First(&dbForm, id)
	return &dbForm, result.Error
}

func UpdateFormById(requestForm *Form) error {
	result := gormDb.Model(&Form{ID: requestForm.ID}).Updates(requestForm)
	return result.Error
}

func QueryAllForm() (*[]Form, error) {
	var dbForm []Form
	result := gormDb.Find(&dbForm)
	return &dbForm, result.Error
}
