package model

type Message struct {
	ID             uint    `gorm:"not null;autoIncrement;primaryKey"`
	IDInSMSService uint    `gorm:"not null"`
	DepartmentID   uint    `gorm:"not null"`
	ReceiverID     uint    `gorm:"not null"`
	Cost           float32 `gorm:"not null"`
}

type MessageRequest struct {
	DepartmentID      uint `json:"DepartmentID" validate:"required"`
	AnswerID          uint `json:"AnswerID" validate:"required"`
	MessageTemplateID uint `json:"MessageTemplateID" validate:"required"`
	InterviewID       uint `json:"InterviewID"`
}

type MessageTemplate struct {
	ID             uint   `gorm:"not null;autoIncrement;primaryKey"`
	Title          string `gorm:"not null;size:255"`
	IDInSMSService uint   `gorm:"not null"`
	OrganizationID uint   `gorm:"not null"`
}

type MessageTemplateRequest struct {
	Title string `json:"Title" validate:"required"`
	Text  string `json:"Text" validate:"required"`
}

type MessageTemplateAPI struct {
	ID             uint
	Title          string
	Text           string
	OrganizationID uint
	Status         uint
}

type AllMessageTemplateAPI struct {
	ID     uint
	Title  string
	Status uint
}

func CreateMessage(requestMessage *Message) error {
	result := gormDb.Create(requestMessage)
	return result.Error
}

func CreateMessageTemplate(requestMessageTemplate *MessageTemplate) error {
	result := gormDb.Create(requestMessageTemplate)
	return result.Error
}

func UpdateMessageTemplateById(requestMessageTemplate *MessageTemplate) error {
	result := gormDb.Model(&MessageTemplate{ID: requestMessageTemplate.ID}).Updates(requestMessageTemplate)
	if result.RowsAffected == 0 && result.Error == nil {
		return ErrNoRowsAffected
	}
	return result.Error
}

func QueryMessageTemplateById(id uint) (*MessageTemplate, error) {
	var dbMessageTemplate MessageTemplate
	result := gormDb.First(&dbMessageTemplate, id)
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
