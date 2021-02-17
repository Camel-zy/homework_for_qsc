package model

import (
	"time"
)

type Message struct {
	ID             uint      `gorm:"not null;autoIncrement;primaryKey"`
	IDInSMSService uint      `gorm:"not null"`
	SenderID       uint      `gorm:"not null"`
	ReceiverID     uint      `gorm:"not null"`
	Text           string    `gorm:"size:1023"`
	Reply          string    `gorm:"size:1023"`
	CreatedAt      time.Time `gorm:"autoCreateTime:nano"`
}

type MessageRequest struct {
	SenderID          uint `json:"SenderID" validate:"required"`
	ReceiverID        uint `json:"ReceiverID" validate:"required"`
	MessageTemplateID uint `json:"MessageTemplateID" validate:"required"`
	JoinedInterviewID uint `json:"JoinedInterviewID"`
	FormID            uint `json:"FormID"` // TODO(TO/GA): wait for model.form
}

type MessageApi struct {
	ID         uint
	SenderID   uint
	ReceiverID uint
	Text       string
	Reply      string
}

type MessageTemplate struct {
	ID             uint   `gorm:"not null;autoIncrement;primaryKey"`
	IDInSMSService uint   `gorm:"not null"`
	Description    string `gorm:"size:255"`
	Text           string `gorm:"size:1023"`
	OrganizationID uint   `gorm:"not null"`
	Status         uint   `gorm:"default:0"`
	UpdatedAt      time.Time
}

type MessageTemplateRequest struct {
	Description    string `json:"Description" validate:"required"`
	Text           string `json:"Text" validate:"required"`
	OrganizationID uint   `json:"OrganizationID"` // not required because it might be 0
}

type MessageTemplateApi struct {
	ID             uint
	Description    string
	Text           string
	OrganizationID uint
	Status         uint
}

type AllMessageTemplateApi struct {
	ID          uint
	Description string
	Status      uint
}

func CreateMessage(requestMessage *Message) error {
	result := gormDb.Create(requestMessage)
	return result.Error
}

func QueryMessageById(id uint) (*MessageApi, error) {
	var dbMessage MessageApi
	result := gormDb.Model(&Message{}).First(&dbMessage, id)
	return &dbMessage, result.Error
}

func UpdateMessageById(requestMessage *Message) error {
	result := gormDb.Model(&Message{ID: requestMessage.ID}).Updates(requestMessage)
	return result.Error
}

func CreateMessageTemplate(requestMessageTemplate *MessageTemplate) error {
	result := gormDb.Create(requestMessageTemplate)
	return result.Error
}

func QueryMessageTemplateById(id uint) (*MessageTemplateApi, error) {
	var dbMessageTemplate MessageTemplateApi
	result := gormDb.Model(&MessageTemplate{}).First(&dbMessageTemplate, id)
	return &dbMessageTemplate, result.Error
}

func QueryAllMessageTemplateInOrganization(oid uint) (*[]AllMessageTemplateApi, error) {
	var dbMessageTemplate []AllMessageTemplateApi
	if findOrganizationError := gormDb.First(&Organization{}, oid).Error; findOrganizationError != nil {
		return nil, findOrganizationError
	}
	result := gormDb.Model(&MessageTemplate{}).Where(&MessageTemplate{OrganizationID: oid}).Find(&dbMessageTemplate)
	return &dbMessageTemplate, result.Error
}

func UpdateMessageTemplateById(requestMessageTemplate *MessageTemplate) error {
	result := gormDb.Model(&MessageTemplate{ID: requestMessageTemplate.ID}).Updates(requestMessageTemplate)
	return result.Error
}
