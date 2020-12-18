package database

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type JoinedDepartment struct {
	ID          uint      `gorm:"not_null;auto_increment;primary_key"`
	UID         uint      `gorm:"not_null"`
	DID         uint      `gorm:"not_null"`
	Privilege   uint      `gorm:"default:'2'"`
	JoinedTime  time.Time `gorm:"not_null"`
	UpdateTime  time.Time `gorm:"not_null"`
}

func createJoinedDepartment(requestJoinedDepartment *JoinedDepartment) error {
	DB.Create(requestJoinedDepartment)
	return nil
}

func queryJoinedDepartment(ID uint) (*JoinedDepartment, error) {
	var result JoinedDepartment
	if err := DB.First(&result, "ID = ?", ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else {
		return &result, nil
	}
}

func updateJoinedDepartment(requestJoinedDepartment *JoinedDepartment) error {
	var result JoinedDepartment
	if err := DB.First(&result, "name = ?", requestJoinedDepartment.Name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	} else {
		if result := DB.Model(&result).Updates(requestJoinedDepartment); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}
}