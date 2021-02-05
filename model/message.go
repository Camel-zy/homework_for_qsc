package model

import (
	"time"
)

type Message struct{
	ID           uint       `gorm:"not null;autoIncrement;primaryKey"`
	SenderID     uint       `gorm:"not null"`
	ReceiverID   uint       `gorm:"not null"`
	Text         string     `gorm:"size:255"`
	Reply        string     `gorm:"size:255"`
	CreatedTime  time.Time  `gorm:"not null"`
}

func CreateMessage(requestMessage *Message)  error {
    result := gormDb.Create(requestMessage)
    return result.Error
}

func QueryMessageById (id uint) (*Message,error){
    var dbMessage Message
    result := gormDb.First(&dbMessage, id)
    return &dbMessage, result.Error
}

func UpdateMessageById(requestMessage *Message) error {
	result := gormDb.Model(&Message{ID: requestMessage.ID}).Updates(requestMessage)
	return result.Error
}
