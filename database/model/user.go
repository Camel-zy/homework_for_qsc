package model

import "time"

type User struct {
	ID           uint       `gorm:"not null;autoIncrement"`
	Name         string     `gorm:"size:40;not null"`
	Nickname     string     `gorm:"size:40"`
	ZJUid        string     `gorm:"column:zju_id;size:10;unique;not null"`
	Mobile       string     `gorm:"size:15"`
	Email        string     `gorm:"size:40"`
	IP           string     `gorm:"size:30"`
	IsSuperuser  uint       `gorm:"default:0"`
	UserAgent    string     `gorm:"size:50"`
	UpdatedTime  time.Time  `gorm:"not null"`
}

type Organization struct {
	ID           uint       `gorm:"not null;autoIncrement;primaryKey"`
	Name         string     `gorm:"size:40;not null"`
	Description  string     `gorm:"size:200"`
	UpdateTime   time.Time  `gorm:"not null"`
}

type Department struct {
	ID              uint          `gorm:"not null;autoIncrement;primaryKey"`
	Name            string        `gorm:"size:40;not null"`
	OrganizationID  uint          `gorm:"not null"`
	Organization    Organization  // FOREIGN KEY (OrganizationID) REFERENCES Organization(OrganizationID)
	Description     string        `gorm:"size:200"`
	UpdateTime      time.Time     `gorm:"not null"`
}

type JoinedDepartment struct {
	ID            uint       `gorm:"not null;autoIncrement;primaryKey"`
	UserID        uint       `gorm:"not null"`
	DepartmentID  uint       `gorm:"not null"`
	Privilege     uint       `gorm:"default:2"`
	JoinedTime    time.Time  `gorm:"not null"`
	UpdateTime    time.Time  `gorm:"not null"`
}
