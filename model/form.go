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
	CreateTime     time.Time      `gorm:"autoCreateTime"`
	OrganizationID uint           `gorm:"not null"`
	Status         uint           `gorm:"not null;default:2"` // 1 pinned, 2 used, 3 unused, 4 abandoned
	Content        datatypes.JSON `gorm:"not null"`
}

// Don't forget to modify CreateFormRequest_ if you modify this
type CreateFormRequest struct {
	Name           string         `json:"Name" validate:"required"`
	Description    string         `json:"Description" validate:"required"`
	OrganizationID uint           `json:"oid" validate:"required"`
	Content        datatypes.JSON `json:"Content" validate:"required"`
}

// Don't forget to modify FormRequest_ if you modify this
type UpdateFormRequest struct {
	Name        string         `json:"Name" validate:"required"`
	Description string         `json:"Description" validate:"required"`
	Status      uint           `json:"Status" validate:"required"`
	Content     datatypes.JSON `json:"Content" validate:"required"`
}

func CreateForm(formRequest *CreateFormRequest) (Form, error) {
	dbForm := Form{}
	copier.Copy(&dbForm, formRequest)
	result := gormDb.Create(&dbForm)
	return dbForm, result.Error
}

func QueryFormById(id uint) (*Form, error) {
	var dbForm Form
	result := gormDb.First(&dbForm, id)
	return &dbForm, result.Error
}

func UpdateFormByID(formRequest *UpdateFormRequest, fid uint) error {
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

	if findOrganizationErr := gormDb.First(&Organization{}, oid).Error; findOrganizationErr != nil {
		return nil, findOrganizationErr
	}

	result := gormDb.Model(&Organization{}).Where(&Form{OrganizationID: oid}).Find(&dbForm)
	return &dbForm, result.Error
}
