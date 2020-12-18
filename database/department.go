package database

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Department struct {
	ID             uint      `gorm:"not null;autoIncrement;primaryKey"`
	Name           string    `gorm:"size:40;not null"`
	OrganizationID uint      `gorm:"not null;ForeignKey:OrganizationID"`
	Description    string    `gorm:"size:200"`
	UpdateTime     time.Time `gorm:"not null"`
}

func CreateDepartment(requestDepartment *Department) error {
	DB.Create(requestDepartment)
	return nil
}

func QueryDepartment(ID uint) (*Department, error) {
	var result Department
	if err := DB.First(&result, "ID = ?", ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else {
		return &result, nil
	}
}

func UpdateDepartment(requestDepartment *Department) error {
	var result Department
	if err := DB.First(&result, "name = ?", requestDepartment.Name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	} else {
		if result := DB.Model(&result).Updates(requestDepartment); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}
}
