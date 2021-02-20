package model

import (
	"time"

	"github.com/jinzhu/copier"
)

type Interview struct {
	ID             uint      `gorm:"not null;autoIncrement;primaryKey"`
	Name           string    `gorm:"size:40;not null"`
	Description    string    `gorm:"size:200"`
	EventID        uint      `gorm:"not null"`
	Event          Event     // FOREIGN KEY (EventID) REFERENCES Event(EventID)
	DepartmentID   uint      `gorm:"not null"`
	OtherInfo      string    `gorm:"size:200"`
	Location       string    `gorm:"size:200"`
	MaxInterviewee uint      `gorm:"default:6"`
	StartTime      time.Time `gorm:"not null"`
	EndTime        time.Time `gorm:"not null"`
	UpdatedTime    time.Time `gorm:"autoUpdateTime"`
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
	StartTime      time.Time
	EndTime        time.Time
}

type JoinedInterview struct {
	ID          uint      `gorm:"not null;autoIncrement;primaryKey"`
	UserID      uint      `gorm:"not null"`
	InterviewID uint      `gorm:"not null"`
	Result      uint      `gorm:"default:0"`
	UpdatedTime time.Time `gorm:"not null"`
}

// TODO(OE.Heart): change this model to fit the logic
type CrossInterview struct {
	ID             uint      `gorm:"not null;autoIncrement;primaryKey"`
	OrganizationID uint      `gorm:"not null"`
	InterviewID    uint      `gorm:"not null"`
	UpdatedTime    time.Time `gorm:"not null"`
}

func CreateInterview(interviewRequest *InterviewRequest, eid uint, did uint) (uint, error) {
	dbInterview := Interview{}
	copier.Copy(&dbInterview, interviewRequest)
	dbInterview.EventID = eid
	dbInterview.DepartmentID = did
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

func UpdateJoinedInterview(id uint, newResult uint) error {
	result := gormDb.Model(&JoinedInterview{ID: id}).Update("result", newResult)
	return result.Error
}
