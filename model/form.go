package model

import (
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/datatypes"
)

// Don't forget to modify Form_ if you modify this
type Form struct {
	ID             uint           `gorm:"not null;autoIncrement;primaryKey"`
	Name           string         `gorm:"size:40;not null"`
	CreateTime     time.Time      `gorm:"size:30;not null"`
	OrganizationID uint           `gorm:"not null"`
	Status         uint           `gorm:"not null"` // 1 pinned, 2 used, 3 unused, 4 abandoned
	Content        datatypes.JSON `gorm:"not null"`
}

// Don't forget to modify FormApi_ if you modify this
type FormApi struct {
	ID             uint
	Name           string         `json:"Name" validate:"required"`
	CreateTime     time.Time      `json:"CreateTime" validate:"required"` // request string must be in RFC 3339 format
	OrganizationID uint           `json:"OrganizationID" validate:"required"`
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
