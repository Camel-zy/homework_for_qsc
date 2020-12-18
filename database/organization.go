package database

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Organization struct {
	ID          uint      `gorm:"not null;autoIncrement;primaryKey"`
	Name        string    `gorm:"size:40;not null"`
	Description string    `gorm:"size:200"`
	UpdateTime  time.Time `gorm:"not null"`
}

func CreateOrganization(requestOrganization *Organization) error {
	/*
		if err := DB.First(&Organization{}, "name = ?", requestOrganization.Name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			DB.Create(requestOrganization)
			return nil
		} else {
			return errors.New("user exists")
		}
	*/
	DB.Create(requestOrganization)
	return nil
}

func QueryOrganization(ID uint) (*Organization, error) {
	var result Organization
	if err := DB.First(&result, "ID = ?", ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else {
		return &result, nil
	}
}

func UpdateOrganization(requestOrganization *Organization) error {
	var result Organization
	if err := DB.First(&result, "name = ?", requestOrganization.Name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	} else {
		if result := DB.Model(&result).Updates(requestOrganization); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}
}
