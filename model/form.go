package model

import (
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Don't forget to modify Form_ if you modify this
type Form struct {
	ID             uint   `gorm:"not null;autoIncrement;primaryKey"`
	Name           string `gorm:"size:40;not null"`
	Description    string
	CreateTime     time.Time      `gorm:"autoCreateTime"`
	OrganizationID uint           `gorm:"not null"`
	Status         uint           `gorm:"not null;default:2"` // 1 pinned, 2 used, 3 unused, 4 abandoned
	Content        datatypes.JSON `gorm:"not null"`
	Deleted        gorm.DeletedAt
}

// Don't forget to modify CreateFormRequest_ if you modify this
type CreateFormRequest struct {
	Name        string         `json:"Name" validate:"required"`
	Description string         `json:"Description" validate:"required"`
	Content     datatypes.JSON `json:"Content" validate:"required"`
}

// Don't forget to modify FormRequest_ if you modify this
type UpdateFormRequest struct {
	Name        string         `json:"Name"`
	Description string         `json:"Description"`
	Status      uint           `json:"Status"`
	Content     datatypes.JSON `json:"Content"`
}

func CreateForm(formRequest *Form) (Form, error) {
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

	result := gormDb.Model(&Form{}).Where(&Form{OrganizationID: oid}).Find(&dbForm)
	return &dbForm, result.Error
}

func DeleteForm(fid uint) error {
	result := gormDb.Delete(&Form{}, fid)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}
