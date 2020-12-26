package model

import "time"

type Message struct{
	ID           uint       `gorm:"not null;autoIncrement;primaryKey"`
	SenderID     uint       `gorm:"not null"`
	ReceiverID   uint       `gorm:"not null"`
	Text         string     `gorm:"size:255"`
	Reply        string     `gorm:"size:255"`
    CreatedTime  time.Time  `gorm:"not null"`
}
