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
	Name        string `json:"Name" validate:"required"`
	Description string `json:"Description" validate:"required"`
	Content     string `example:"JSON"`
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
	Name      string `validate:"required"`
	Mobile    string `validate:"required"`
	Intention []Intention
	Content   string `validate:"required" example:"JSON"`
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
	DepartmentID     uint   // 志愿部门
	IntentRank       uint   // 第几志愿
	Round            uint   // 公海为1，一面为2，以此类推
	InterviewOptions string // 发送选择面试场次的短信用
	Status           uint   // 1 已确认本轮面试时间/正在面试，2 本轮通过，3 已发送下轮分配短信，4 纳入组织，5 拒绝, 6 面试者选择没有合适的下轮面试时间
}
