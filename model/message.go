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
	CreateTime     time.Time `gorm:"autoCreateTime"`
}

// TODO(TO/GA): wait for model.form and logic
type MessageRequest struct {
	SenderID          uint `json:"SenderID" validate:"required"`
	ReceiverID        uint `json:"ReceiverID" validate:"required"`
	MessageTemplateID uint `json:"MessageTemplateID" validate:"required"`
	InterviewID       uint `json:"InterviewID"`
	CrossInterviewID  uint `json:"CrossInterviewID"` // TODO(TO/GA): wait for logic
	FormID            uint `json:"FormID"`
}

type MessageAPI struct {
	ID         uint
	SenderID   uint
	ReceiverID uint
	Text       string
	Reply      string
}

type MessageTemplate struct {
	ID             uint      `gorm:"not null;autoIncrement;primaryKey"`
	IDInSMSService uint      `gorm:"not null"`
	Description    string    `gorm:"size:255"`
	Text           string    `gorm:"size:1023"`
	OrganizationID uint      `gorm:"not null"`
	Status         uint      `gorm:"default:0"`
	UpdatedTime    time.Time `gorm:"autoUpdateTime"`
}

type MessageTemplateRequest struct {
	Description string `json:"Description" validate:"required"`
	Text        string `json:"Text" validate:"required"`
}

type MessageTemplateAPI struct {
	ID             uint
	Description    string
	Text           string
	OrganizationID uint
	Status         uint
}

type AllMessageTemplateAPI struct {
	ID          uint
	Description string
	Status      uint
}

func CreateMessage(requestMessage *Message) error {
	result := gormDb.Create(requestMessage)
	return result.Error
}

func QueryMessageById(id uint) (*Message, error) {
	var dbMessage Message
	result := gormDb.First(&dbMessage, id)
	return &dbMessage, result.Error
}

func QueryMessageAPIById(id uint) (*MessageAPI, error) {
	var dbMessage MessageAPI
	result := gormDb.Model(&Message{}).First(&dbMessage, id)
	return &dbMessage, result.Error
}

func UpdateMessageById(requestMessage *Message) error {
	result := gormDb.Model(&Message{ID: requestMessage.ID}).Updates(requestMessage)
	if result.RowsAffected == 0 && result.Error == nil {
		return ErrNoRowsAffected
	}
	return result.Error
}

func CreateMessageTemplate(requestMessageTemplate *MessageTemplate) error {
	result := gormDb.Create(requestMessageTemplate)
	return result.Error
}

func QueryMessageTemplateById(id uint) (*MessageTemplate, error) {
	var dbMessageTemplate MessageTemplate
	result := gormDb.First(&dbMessageTemplate, id)
	return &dbMessageTemplate, result.Error
}

func QueryMessageTemplateAPIById(id uint) (*MessageTemplateAPI, error) {
	var dbMessageTemplate MessageTemplateAPI
	result := gormDb.Model(&MessageTemplate{}).First(&dbMessageTemplate, id)
	return &dbMessageTemplate, result.Error
}

func QueryAllMessageTemplateInOrganization(oid uint) (*[]MessageTemplate, error) {
	var dbMessageTemplate []MessageTemplate
	if findOrganizationError := gormDb.First(&Organization{}, oid).Error; findOrganizationError != nil {
		return nil, findOrganizationError
	}
	result := gormDb.Where(&MessageTemplate{OrganizationID: oid}).Find(&dbMessageTemplate)
	return &dbMessageTemplate, result.Error
}

func QueryAllMessageTemplateAPIInOrganization(oid uint) (*[]AllMessageTemplateAPI, error) {
	var dbMessageTemplate []AllMessageTemplateAPI
	if findOrganizationError := gormDb.First(&Organization{}, oid).Error; findOrganizationError != nil {
		return nil, findOrganizationError
	}
	result := gormDb.Model(&MessageTemplate{}).Where(&MessageTemplate{OrganizationID: oid}).Find(&dbMessageTemplate)
	return &dbMessageTemplate, result.Error
}

func UpdateMessageTemplateById(requestMessageTemplate *MessageTemplate) error {
	result := gormDb.Model(&MessageTemplate{ID: requestMessageTemplate.ID}).Updates(requestMessageTemplate)
	if result.RowsAffected == 0 && result.Error == nil {
		return ErrNoRowsAffected
	}
	return result.Error
}
