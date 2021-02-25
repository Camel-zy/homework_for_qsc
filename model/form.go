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
	Description    string         `gorm:"not null"`
	CreateTime     time.Time      `gorm:"size:30;not null"`
	OrganizationID uint           `gorm:"not null"`
	DepartmentID   uint           `gorm:"not null"`
	Status         uint           `gorm:"not null"` // 1 pinned, 2 used, 3 unused, 4 abandoned
	Content        datatypes.JSON `gorm:"not null"`
}

// Don't forget to modify FormRequest_ if you modify this
type FormRequest struct {
	Name        string         `json:"Name"`
	Description string         `json:"Description"`
	CreateTime  time.Time      `json:"CreateTime"`
	Status      uint           `json:"Status"`
	Content     datatypes.JSON `json:"Content"`
}

// Don't forget to modify FormResponse_ if you modify this
type FormResponse struct {
	ID             uint
	Name           string    `json:"Name" validate:"required"`
	Description    string    `json:"Description"`
	CreateTime     time.Time `json:"CreateTime" validate:"required"` // request string must be in RFC 3339 format
	OrganizationID uint      `json:"OrganizationID" validate:"required"`
	DepartmentID   uint
	Status         uint
	Content        datatypes.JSON
}

func CreateForm(formRequest *FormRequest, oid uint, did uint) (uint, error) {
	dbForm := Form{}
	copier.Copy(&dbForm, formRequest)
	dbForm.OrganizationID = oid
	dbForm.DepartmentID = did
	result := gormDb.Create(&dbForm)
	return dbForm.ID, result.Error
}

func QueryFormById(id uint) (*Form, error) {
	var dbForm Form
	result := gormDb.First(&dbForm, id)
	return &dbForm, result.Error
}

func UpdateFormByID(formRequest *FormRequest, fid uint) error {
	dbForm := Form{}
	copier.Copy(&dbForm, formRequest)
	result := gormDb.Model(&Form{ID: fid}).Updates(&dbForm)
	return result.Error
}

func QueryAllForm() (*[]Form, error) {
	var dbForm []Form
	result := gormDb.Find(&dbForm)
	return &dbForm, result.Error
}

func QueryAllFormByOid(oid uint) (*[]Form, error) {
	var dbForm []Form
	result := gormDb.Model(&Organization{}).Where(&Form{OrganizationID: oid}).Find(&dbForm)
	return &dbForm, result.Error
}
