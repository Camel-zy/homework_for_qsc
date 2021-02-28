package model

import (
	"gorm.io/datatypes"
)

// Don't forget to modify Answer_ if you modify this
type Answer struct {
	ID        uint           `gorm:"not null;autoIncrement;primaryKey"`
	FormID    uint           `gorm:"not null"`
	EventID   uint           `gorm:"not null"`
	Name      string         `gorm:"not null"`
	ZJUid     string         `gorm:"column:zju_id;size:10;not null"`
	Mobile    string         `gorm:"size:11;not null"`
	Intention datatypes.JSON `gorm:"not null"`
	Status    uint           `gorm:"not null;default:2"` // 1 abandoned, 2 used
	Content   datatypes.JSON `gorm:"not null"`
}

// Don't forget to modify AnswerRequest_ if you modify this
type AnswerRequest struct {
	Name      string         `json:"Name" validate:"required"`
	Mobile    string         `json:"Mobile" validate:"required"`
	Intention []IntentionRequest    `json:"Intention"`
	Content   datatypes.JSON `json:"Content" validate:"required"`
}

type Intention struct {
	DepartmentID uint `json:"department_id" validate:"required"`
	IntentRank   uint `json:"intent_rank"` // this feature hasn't been implemented
}

type IntentionRequest struct {
	DepartmentID string `json:"department_id"`
	IntentRank   uint `json:"intent_rank"` // this feature hasn't been implemented
}

// Don't forget to modify AnswerResponse_ if you modify this
type AnswerResponse struct {
	ID        uint
	FormID    uint
	EventID   uint
	Name      string
	ZJUid     string `gorm:"column:zju_id"`
	Mobile    string
	Intention datatypes.JSON
	Content   datatypes.JSON
}

func QueryAnswerByID(id uint) (*Answer, error) {
	var dbAnswer Answer
	result := gormDb.First(&dbAnswer, id)
	return &dbAnswer, result.Error
}

func QueryAnswer(fid uint, zjuid string, eid uint) (*AnswerResponse, error) {
	var dbAnswer AnswerResponse
	result := gormDb.Model(&Answer{FormID: fid, ZJUid: zjuid, EventID: eid}).First(&dbAnswer)
	return &dbAnswer, result.Error
}

func CreateAnswer(answerRequest *Answer, fid uint, zjuid string, eid uint) (uint, error) {
	answerRequest.FormID = fid
	answerRequest.ZJUid = zjuid
	answerRequest.EventID = eid
	result := gormDb.Create(answerRequest)
	return answerRequest.ID, result.Error
}

func UpdateAnswer(answerRequest *Answer, fid uint, zjuid string, eid uint) error {
	result := gormDb.Model(&Answer{}).
		Where(&Answer{FormID: fid, ZJUid: zjuid, EventID: eid}).
		Updates(answerRequest)
	return result.Error
}
