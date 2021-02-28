package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/jinzhu/copier"
)

type Interview struct {
	ID                 uint   `gorm:"not null;autoIncrement;primaryKey"`
	Name               string `gorm:"size:40;not null"`
	Description        string `gorm:"size:200"`
	EventID            uint   `gorm:"not null"`
	Event              Event  // FOREIGN KEY (EventID) REFERENCES Event(EventID)
	DepartmentID       uint   `gorm:"not null"`
	OtherInfo          string `gorm:"size:200"`
	Location           string `gorm:"size:200"`
	MaxInterviewee     uint   `gorm:"default:6"`
	Round              uint   `gorm:"not null;default:1"` // 一面为2，二面为3，以此类推
	CrossTag           uint   `gorm:"not null;default:1"` // 2 represent CrossInterview, default 1
	InvolvedDepartment datatypes.JSON
	StartTime          time.Time `gorm:"not null"`
	EndTime            time.Time `gorm:"not null"`
	UpdatedTime        time.Time `gorm:"autoUpdateTime"`
}

type InterviewRequest struct {
	Name           string    `json:"Name" validate:"required"`
	Description    string    `json:"Description"`
	OtherInfo      string    `json:"OtherInfo"`
	Location       string    `json:"Location"`
	MaxInterviewee uint      `json:"MaxInterviewee"`                // default 6
	StartTime      time.Time `json:"StartTime" validate:"required"` // request string must be in RFC 3339 format
	EndTime        time.Time `json:"EndTime" validate:"required"`   // request string must be in RFC 3339 format
}

type InterviewResponse struct {
	ID             uint
	Name           string
	Description    string
	EventID        uint
	DepartmentID   uint
	OtherInfo      string
	Location       string
	MaxInterviewee uint // default 6
	Round          uint // 一面为2，二面为3，以此类推
	StartTime      time.Time
	EndTime        time.Time
}

func CreateInterview(interviewRequest *InterviewRequest, eid uint, did uint, rnd uint) (uint, error) {
	dbInterview := Interview{}
	copier.Copy(&dbInterview, interviewRequest)
	dbInterview.EventID = eid
	dbInterview.DepartmentID = did
	dbInterview.Round = rnd
	result := gormDb.Create(&dbInterview)
	return dbInterview.ID, result.Error
}

func UpdateInterviewByID(requestInterview *InterviewRequest, iid uint) error {
	dbInterview := Interview{}
	copier.Copy(&dbInterview, requestInterview)
	result := gormDb.Model(&Interview{ID: iid}).Updates(&dbInterview)
	return result.Error
}

func QueryInterviewByID(id uint) (*InterviewResponse, error) {
	var dbInterview InterviewResponse
	result := gormDb.Model(&Interview{}).First(&dbInterview, id)
	return &dbInterview, result.Error
}

func QueryInterviewByIDWithPreload(id uint) (*Interview, error) {
	var dbInterview Interview
	result := gormDb.Preload(clause.Associations).Model(&Interview{}).First(&dbInterview, id)
	return &dbInterview, result.Error
}

func QueryInterviewByIDInEvent(eid uint, iid uint) (*InterviewResponse, error) {
	var dbInterview InterviewResponse
	result := gormDb.Model(&Interview{}).Where(&Interview{ID: iid, EventID: eid}).First(&dbInterview)
	return &dbInterview, result.Error
}

func QueryAllInterviewInEvent(eid uint) (*[]Brief, error) {
	var dbInterview []Brief
	if findEventError := gormDb.First(&Event{}, eid).Error; findEventError != nil {
		return nil, findEventError
	}
	result := gormDb.Model(&Interview{}).Where(&Interview{EventID: eid}).Find(&dbInterview)
	return &dbInterview, result.Error
}

func QueryAllInterviewOfDepartmentInEvent(eid, did uint) (*[]InterviewResponse, error) {
	var dbInterview []InterviewResponse
	result := gormDb.Model(&Interview{}).Where(&Interview{EventID: eid, DepartmentID: did}).Find(&dbInterview)
	return &dbInterview, result.Error
}

func QueryAllInterviewOfRound(eid uint, did uint, rnd uint) (*[]Brief, error) {
	var dbInterview []Brief
	if findEventError := gormDb.First(&Event{}, eid).Error; findEventError != nil {
		return nil, findEventError
	}
	result := gormDb.Model(&Interview{}).Where(&Interview{EventID: eid, DepartmentID: did, Round: rnd}).Find(&dbInterview)
	return &dbInterview, result.Error
}

func QueryInterviewByIntervieweeAndRound(vid uint, rnd uint) (*Interview, error) {
	var dbInterview Interview
	if gormDb.Joins(`INNER JOIN joined_interviews
	ON interviews.id = joined_interviews.interview_id
	AND joined_interviews.interviewee_id = ?
	AND interviews.round = ?`, vid, rnd).Find(&dbInterview).RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &dbInterview, nil
}
