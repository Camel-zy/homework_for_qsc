package database

import "time" // TODO: move this another place (RalXYZ)

type Organization struct {
	ID          uint      `gorm:"not_null;auto_increment;primary_key"`
	Name        string    `gorm:"size:40;not_null"`
	Description string    `gorm:"size:200"`
	UpdateTime  time.Time `gorm:"not_null"`
}
