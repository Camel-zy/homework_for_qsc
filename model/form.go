package model

import (
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/datatypes"
)

// Don't forget to modify Form_ if you modify this
type Form struct {
	ID             uint           `gorm:"not null;autoIncrement;primaryKey"`
	Name           string         `gorm:"size:40;not null"`
	CreateTime     time.Time      `gorm:"size:30;not null"`
	OrganizationID uint           `gorm:"not null"`
	DepartmentID   uint           `gorm:"not null"`
	Status         uint           `gorm:"not null"` // 0 pinned, 1 used, 2 unused, 3 abandoned
	Content        datatypes.JSON `gorm:"type:datatypes;not null"`
}

// Don't forget to modify FormApi_ if you modify this
type FormApi struct {
	ID             uint
	Name           string         `json:"Name" validate:"required"`
	CreateTime     time.Time      `json:"CreateTime" validate:"required"` // request string must be in RFC 3339 format
	OrganizationID uint           `json:"OrganizationID" validate:"required"`
	DepartmentID   uint           `json:"DepartmentID" validate:"required"`
	Status         uint           `json:"Status" validate:"required"`
	Content        datatypes.JSON `json:"Content" validate:"required"`
}

// Don't forget to modify Answer_ if you modify this
type Answer struct {
	ID        uint           `gorm:"not null;autoIncrement;primaryKey"`
	FormID    uint           `gorm:"not null"`
	Name      string         `gorm:"not null"`
	ZJUid     string         `gorm:"column:zju_id;size:10;not null"`
	Mobile    string         `gorm:"size:11;not null"`
	Intention string         `gorm:"not null"`
	EventID   uint           `gorm:"not null"`
	Status    uint           `gorm:"not null"` // 0 abandoned, 1 used
	Content   datatypes.JSON `gorm:"type:datatypes;not null"`
}

// Don't forget to modify AnswerRequest_ if you modify this
type AnswerRequest struct {
	Name      string         `json:"Name"`
	ZJUid     string         `json:"ZJUid"`
	Mobile    string         `json:"Mobile"`
	Intention string         `json:"Intention"`
	Content   datatypes.JSON `json:"Content"`
}

// Don't forget to modify AnswerResponse_ if you modify this
type AnswerResponse struct {
	ID        uint
	FormID    uint
	Name      string
	ZJUid     string `gorm:"column:zju_id"`
	Mobile    string
	Intention string
	EventID   uint
	Content   datatypes.JSON
}

func CreateForm(requestForm *FormApi) error {
	dbForm := Form{}
	copier.Copy(&dbForm, requestForm)
	result := gormDb.Create(&dbForm)
	return result.Error
}

func QueryFormById(id uint) (*Form, error) {
	var dbForm Form
	result := gormDb.First(&dbForm, id)
	return &dbForm, result.Error
}

func UpdateFormByID(requestForm *FormApi) error {
	dbForm := Form{}
	copier.Copy(&dbForm, requestForm)
	result := gormDb.Model(&Form{ID: requestForm.ID}).Updates(&dbForm)
	return result.Error
}

func QueryAllForm() (*[]Form, error) {
	var dbForm []Form
	result := gormDb.Find(&dbForm)
	return &dbForm, result.Error
}

func QueryAnswerById(id uint) (*Answer, error) {
	var dbAnswer Answer
	result := gormDb.First(&dbAnswer, id)
	return &dbAnswer, result.Error
}

func QueryAnswer(fid uint, zjuid string, eid uint) (*AnswerResponse, error) {
	var dbAnswer AnswerResponse
	result := gormDb.Model(&Answer{}).First(&dbAnswer, fid, zjuid, eid)
	return &dbAnswer, result.Error
}

func CreateAnswer(answerRequest *AnswerRequest, fid uint, zjuid string, eid uint) (uint, error) {
	dbAnswer := Answer{}
	copier.Copy(&dbAnswer, answerRequest)
	dbAnswer.FormID = fid
	dbAnswer.ZJUid = zjuid
	dbAnswer.EventID = eid
	dbAnswer.Status = 1
	result := gormDb.Create(&dbAnswer)
	return dbAnswer.ID, result.Error
}

func UpdateAnswer(answerRequest *AnswerRequest, fid uint, zjuid string, eid uint) error {
	dbAnswer := Answer{}
	copier.Copy(&dbAnswer, answerRequest)
	result := gormDb.Model(&Answer{FormID: fid, ZJUid: zjuid, EventID: eid}).Updates(&dbAnswer)
	return result.Error
}
