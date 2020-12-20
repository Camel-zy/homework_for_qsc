package utils

import (
	"gorm.io/gorm"
)

func Create(db *gorm.DB, value interface{}) error {
	if result := db.Create(value); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
