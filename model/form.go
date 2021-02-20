package model

import (
	"github.com/jinzhu/copier"
	"gorm.io/datatypes"
	"time"
)

type Form struct {
	ID             uint           `gorm:"not null;autoIncrement;primaryKey"`
	Name           string         `gorm:"size:40;not null"`
	CreateTime     time.Time      `gorm:"size:30;not null"`
	UserID         uint           `gorm:"not null"`
	User           User           // FOREIGN KEY (UserID) REFERENCES User(UserID)
	OrganizationID uint           `gorm:"not null"`
	DepartmentID   uint           `gorm:"not null"`
	Status         uint           `gorm:"not null"` // 0 pinned, 1 used, 2 unused, 3 abandoned
	Content        datatypes.JSON `gorm:"type:datatypes;not null"`
}

type FormApi struct {
	ID             uint
	Name           string         `json:"Name" validate:"required"`
	CreateTime     time.Time      `json:"CreateTime" validate:"required"` // request string must be in RFC 3339 format
	UserID         uint           `json:"UserID" validate:"required"`
	OrganizationID uint           `json:"OrganizationID" validate:"required"`
	DepartmentID   uint           `json:"DepartmentID" validate:"required"`
	Status         uint           `json:"Status" validate:"required"`
	Content        datatypes.JSON `json:"Content" validate:"required"`
}

func CreateForm(requestForm *FormApi) error {
	dbForm := Form{}
	copier.Copy(&dbForm, requestForm)
	result := gormDb.Create(&dbForm)
	return result.Error
}

func QueryFormById(id uint) (*Form, error) {
	var dbForm Form
	result := gormDb.First(&dbForm, id)
	return &dbForm, result.Error
}

func UpdateFormByID(requestForm *FormApi) error {
	dbForm := Form{}
	copier.Copy(&dbForm, requestForm)
	result := gormDb.Model(&Form{ID: requestForm.ID}).Updates(&dbForm)
	return result.Error
}

func QueryAllForm() (*[]Form, error) {
	var dbForm []Form
	result := gormDb.Find(&dbForm)
	return &dbForm, result.Error
}
