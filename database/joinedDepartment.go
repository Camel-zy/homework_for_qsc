package database

import (
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