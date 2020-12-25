package model

import "time"

type Event struct {
	ID              uint          `gorm:"not null;autoIncrement;primaryKey"`
	Name            string        `gorm:"size:40;not null"`
	Description     string        `gorm:"size:200"`
	OrganizationID  uint          `gorm:"not null"`
	Organization    Organization  // FOREIGN KEY (OrganizationID) REFERENCES Organization(OrganizationID)
	Status          uint          `gorm:"default:1"`
	OtherInfo       string        `gorm:"size:200"`
	StartTime       time.Time     `gorm:"size:30;not null"`
	EndTime         time.Time     `gorm:"size:30;not null"`
	UpdatedTime     time.Time     `gorm:"not null"`
}

type Interview struct {
	ID              uint       `gorm:"not null;autoIncrement;primaryKey"`
	Name            string     `gorm:"size:40;not null"`
	Description     string     `gorm:"size:200"`
	EventID         uint       `gorm:"not null"`
	Event           Event      // FOREIGN KEY (EventID) REFERENCES Event(EventID)
	OtherInfo       string     `gorm:"size:200"`
	Location        string     `gorm:"size:200"`
	MaxInterviewee  uint       `gorm:"default:6"`
	StartTime       time.Time  `gorm:"not null"`
	EndTime         time.Time  `gorm:"not null"`
	UpdatedTime     time.Time  `gorm:"not null"`
}

type JoinedInterview struct {
	ID           uint       `gorm:"not null;autoIncrement;primaryKey"`
	UserID       uint       `gorm:"not null"`
	InterviewID  uint       `gorm:"not null"`
	Result       uint       `gorm:"default:0"`
	UpdatedTime  time.Time  `gorm:"not null"`
}
