package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// 一个志愿一条记录
// Don't forget to modify Interviewee_ if you modify this
type Interviewee struct {
	ID               uint           `gorm:"not null;autoIncrement;primaryKey"`
	UUID             uuid.UUID      `gorm:"not null;type:uuid"`
	EventID          uint           `gorm:"not null"`
	AnswerID         uint           `gorm:"not null"`
	DepartmentID     uint           `gorm:"not null"`           // 志愿部门
	IntentRank       uint           `gorm:"not null;default:0"` // 第几志愿
	Round            uint           `gorm:"not null;default:1"` // 公海为1，一面为2，以此类推
	InterviewOptions datatypes.JSON // 发送选择面试场次的短信用
	Status           uint           `gorm:"not null; default:2"` // 1 已确认本轮面试时间/正在面试，2 本轮通过，3 已发送下轮分配短信，4 纳入组织，5 拒绝
}

type IntervieweeRequest struct {
	InterviewOptions []uint `json:"InterviewOptions"`
}

type JoinedInterview struct {
	ID            uint      `gorm:"not null;autoIncrement;primaryKey"`
	InterviewID   uint      `gorm:"not null"`
	IntervieweeID uint      `gorm:"not null"`
	UpdatedTime   time.Time `gorm:"not null"`
}

// "Create Hook" of GORM
func (i *Interviewee) BeforeCreate(tx *gorm.DB) (err error) {
	i.UUID = uuid.NewV4()
	return
}

func CreateInterviewee(interviewee *Interviewee) (uint, error) {
	result := gormDb.Create(interviewee)
	return interviewee.ID, result.Error
}

func UpdateInterviewee(interviewee *Interviewee, vid uint) error {
	result := gormDb.Model(&Interviewee{ID: vid}).Updates(interviewee)
	return result.Error
}

func QueryIntervieweeById(id uint) (*Interviewee, error) {
	var dbInterviewee Interviewee
	result := gormDb.First(&dbInterviewee, id)
	return &dbInterviewee, result.Error
}

func UpdateJoinedInterview(id uint, newResult uint) error {
	result := gormDb.Model(&JoinedInterview{ID: id}).Update("result", newResult)
	return result.Error
}

func QueryAllJoinedInterviewOfInterview(iid uint) (*[]JoinedInterview, error) {
	var dbJoinedInterview []JoinedInterview
	result := gormDb.Model(&JoinedInterview{}).Where(&JoinedInterview{InterviewID: iid}).Find(&dbJoinedInterview)
	return &dbJoinedInterview, result.Error
}

func QueryAllIntervieweeByDidAndEid(did, eid uint) (*[]Interviewee, error) {
	var dbInterviewee []Interviewee
	result := gormDb.Model(&Interviewee{}).Where(&Interviewee{DepartmentID: did, EventID: eid}).Find(&dbInterviewee)
	return &dbInterviewee, result.Error
}

func QueryAllIntervieweeByStatus(did, eid, status uint) (*[]Interviewee, error) {
	var dbInterviewee []Interviewee
	result := gormDb.Model(&Interviewee{}).Where(&Interviewee{DepartmentID: did, EventID: eid, Status: status}).Find(&dbInterviewee)
	return &dbInterviewee, result.Error
}

func QueryAllIntervieweeByRoundAndStatus(did, eid, round, status uint) (*[]Interviewee, error) {
	var dbInterviewee []Interviewee
	result := gormDb.Model(&Interviewee{}).Where(&Interviewee{DepartmentID: did, EventID: eid, Round: round, Status: status}).Find(&dbInterviewee)
	return &dbInterviewee, result.Error
}
