package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
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
type CreateFormRequest_ struct {
	Name           string `json:"Name" validate:"required"`
	Description    string `json:"Description" validate:"required"`
	OrganizationID uint   `json:"oid" validate:"required"`
	Content        string `example:"JSON"`
}

type UpdateFormRequest_ struct {
	Name        string `json:"Name" validate:"required"`
	Description string `json:"Description" validate:"required"`
	Status      uint   `json:"Status" validate:"required"`
	Content     string `json:"Content" validate:"required"`
}

// CAUTIOUS: do NOT use this
type Answer_ struct {
	ID        uint
	FormID    uint
	EventID   uint
	Name      string
	ZJUid     string
	Mobile    string
	Intention string `example:"JSON"`
	Status    uint
	Content   string `example:"JSON"`
}

// CAUTIOUS: do NOT use this
type AnswerRequest_ struct {
	Name      string
	Mobile    string
	Intention []Intention
	Content   string `example:"JSON"`
}

// CAUTIOUS: do NOT use this
type AnswerResponse_ struct {
	ID        uint
	FormID    uint
	EventID   uint
	Name      string
	ZJUid     string
	Mobile    string
	Intention string `example:"JSON"`
	Content   string `example:"JSON"`
}

// CAUTIOUS: do NOT use this
type Interviewee_ struct {
	ID               uint
	UUID             uuid.UUID
	EventID          uint
	AnswerID         uint
	DepartmentID     uint
	IntentRank       uint
	Round            uint
	InterviewOptions string `example:"JSON"`
	Status           uint
}
