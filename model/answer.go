package model

import (
	"encoding/json"
	"github.com/jinzhu/copier"
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
	Intention []Intention    `json:"Intention"`
	Content   datatypes.JSON `json:"Content" validate:"required"`
}

type Intention struct {
	DepartmentID uint `json:"department_id" validate:"required"`
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

func QueryAnswerById(id uint) (*Answer, error) {
	var dbAnswer Answer
	result := gormDb.First(&dbAnswer, id)
	return &dbAnswer, result.Error
}

func QueryAnswer(fid uint, zjuid string, eid uint) (*AnswerResponse, error) {
	var dbAnswer AnswerResponse
	result := gormDb.Model(&Answer{FormID: fid, ZJUid: zjuid, EventID: eid}).First(&dbAnswer)
	return &dbAnswer, result.Error
}

func CreateAnswer(answerRequest *AnswerRequest, fid uint, zjuid string, eid uint) (uint, error) {
	dbAnswer := Answer{}
	copier.Copy(&dbAnswer, answerRequest)
	dbAnswer.FormID = fid
	dbAnswer.ZJUid = zjuid
	dbAnswer.EventID = eid
	dbIntention, err := json.Marshal(answerRequest.Intention)
	if err != nil {
		return 0, err
	}
	dbAnswer.Intention = dbIntention
	result := gormDb.Create(&dbAnswer)
	return dbAnswer.ID, result.Error
}

func UpdateAnswer(answerRequest *AnswerRequest, fid uint, zjuid string, eid uint) error {
	dbAnswer := Answer{}
	copier.Copy(&dbAnswer, answerRequest)
	result := gormDb.Model(&Answer{FormID: fid, ZJUid: zjuid, EventID: eid}).Updates(&dbAnswer)
	return result.Error
}
