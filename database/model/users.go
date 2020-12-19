package model

import "time"

type User struct {
	ID          int    `gorm:"not null;autoIncrement"`
	Name        string `gorm:"size:40;not null"`
	Nickname    string `gorm:"size:40"`
	ZJUid       string `gorm:"column:zju_id;size:10;unique;not null"`
	Mobile      string `gorm:"size:15"`
	Email       string `gorm:"size:40"`
	IP          string `gorm:"size:30"`
	IsSuperuser int    `gorm:"default:0"`
	UserAgent   string `gorm:"size:50"`
	UpdatedTime string `gorm:"size:30;not null"`
}

type Organization struct {
	ID          uint      `gorm:"not null;autoIncrement;primaryKey"`
	Name        string    `gorm:"size:40;not null"`
	Description string    `gorm:"size:200"`
	UpdateTime  time.Time `gorm:"not null"`
}

type Department struct {
	ID             uint      `gorm:"not null;autoIncrement;primaryKey"`
	Name           string    `gorm:"size:40;not null"`
	OrganizationID uint      `gorm:"not null;foreignKey"`
	Description    string    `gorm:"size:200"`
	UpdateTime     time.Time `gorm:"not null"`
}

type JoinedDepartment struct {
	ID         uint      `gorm:"not null;autoIncrement;primaryKey"`
	UID        uint      `gorm:"not null"`
	DID        uint      `gorm:"not null"`
	Privilege  uint      `gorm:"default:2"`
	JoinedTime time.Time `gorm:"not null"`
	UpdateTime time.Time `gorm:"not null"`
}

type Event struct {
	ID             uint   `gorm:"not null;autoIncrement;primaryKey"`
	Name           string `gorm:"size:40;not null"`
	Description    string `gorm:"size:200"`
	OrganizationID uint   `gorm:"not null;foreignKey"`
	Status         uint   `gorm:"default:1"`
	OtherInfo      string `gorm:"size:200"`
	StartTime      string `gorm:"size:30;not null"`
	EndTime        string `gorm:"size:30;not null"`
    UpdatedTime    string `gorm:"not null"`
}

type Interview struct {
	ID             uint   `gorm:"not null;autoIncrement;primaryKey"`
	Name           string `gorm:"size:40;not null"`
	Description    string `gorm:"size:200"`
	EventID        uint   `gorm:"not null;foreignKey"`
	OtherInfo      string `gorm:"size:200"`
	Location       string `gorm:"size:200"`
	MaxInterviewee uint   `gorm:"default:6"`
	StartTime      string `gorm:"not null"`
	EndTime        string `gorm:"not null"`
	UpdatedTime    string `gorm:"not null"`
}

type JoinedInterview struct {
	ID          uint   `gorm:"not null;autoIncrement;primaryKey"`
	UID         uint   `gorm:"not null"`
	IID         uint   `gorm:"not null"`
	Result      uint   `gorm:"default:0"`
	UpdatedTime string `gorm:"not null"`
}