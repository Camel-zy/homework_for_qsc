package model

import (
	"time"
)

// CAUTIOUS: do NOT use this
type Form_ struct {
	ID             uint
	Name           string
	CreateTime     time.Time
	OrganizationID uint
	DepartmentID   uint
	Status         uint
	Content        string `example:"JSON"`
}

// CAUTIOUS: do NOT use this
type FormApi_ struct {
	ID             uint
	Name           string
	CreateTime     time.Time
	OrganizationID uint
	DepartmentID   uint
	Status         uint
	Content        string `example:"JSON"`
}

// CAUTIOUS: do NOT use this
type Answer_ struct {
	ID        uint
	FormID    uint
	Name      string
	ZJUid     string
	Mobile    string
	Intention string
	EventID   uint
	Status    uint
	Content   string `example:"JSON"`
}

// CAUTIOUS: do NOT use this
type AnswerRequest_ struct {
	Name      string
	ZJUid     string
	Mobile    string
	Intention string
	Content   string `example:"JSON"`
}

// CAUTIOUS: do NOT use this
type AnswerResponse_ struct {
	ID        uint
	FormID    uint
	Name      string
	ZJUid     string
	Mobile    string
	Intention string
	EventID   uint
	Content   string `example:"JSON"`
}
