package database

import (
	"time"
)

type JoinedDepartment struct {
	ID         uint      `gorm:"not null;autoIncrement;primaryKey"`
	UID        uint      `gorm:"not null"`
	DID        uint      `gorm:"not null"`
	Privilege  uint      `gorm:"default:'2'"`
	JoinedTime time.Time `gorm:"not null"`
	UpdateTime time.Time `gorm:"not null"`
}
