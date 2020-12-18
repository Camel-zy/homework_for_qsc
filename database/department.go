package database

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Department struct {
	ID          	uint      `gorm:"not_null;auto_increment;primary_key"`
	Name        	string    `gorm:"size:40;not_null"`
	OrganizationID 	uint      `gorm:"not_null;ForeignKey:OrganizationID"`
	Description 	string    `gorm:"size:200"`
	UpdateTime  	time.Time `gorm:"not_null"`
}

func createDepartment(requestDepartment *Department) error {
	DB.Create(requestDepartment)
	return nil
}

func queryDepartment(ID uint) (*Department, error) {
	var result Department
	if err := DB.First(&result, "ID = ?", ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else {
		return &result, nil
	}
}

func updateDepartment(requestDepartment *Department) error {
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