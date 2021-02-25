package model

import (
	"github.com/satori/go.uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

// 一个志愿一条记录
type Interviewee struct {
	ID               uint           `gorm:"not null;autoIncrement;primaryKey"`
	UUID             uuid.UUID      `gorm:"not null;type:uuid"`
	EventID          uint           `gorm:"not null"`
	AnswerID         uint           `gorm:"not null"`
	DepartmentID     uint           `gorm:"not null"`           // 志愿部门
	IntentRank       uint           `gorm:"not null;default:0"` // 第几志愿
	Round            uint           `gorm:"not null;default:1"` // 公海为1，一面为2，以此类推
	SentMessage      uint           `gorm:"not null;default:1"` // 发送过选择面试场次短信的为2，没有为1
	InterviewOptions datatypes.JSON // 发送选择面试场次的短信用
	Status           uint           `gorm:"not null; default:1"` // 1 本轮接受但还没选择下轮面试时间，2 面试进行中，3 纳入组织，4 拒绝
}

type IntervieweeRequest struct {
	InterviewOptions []uint `json:"InterviewOptions"`
}

type JoinedInterview struct {
	ID            uint      `gorm:"not null;autoIncrement;primaryKey"`
	InterviewID   uint      `gorm:"not null"`
	IntervieweeID uint      `gorm:"not null"`
	Result        uint      `gorm:"default:0"`
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

func UpdateJoinedInterview(id uint, newResult uint) error {
	result := gormDb.Model(&JoinedInterview{ID: id}).Update("result", newResult)
	return result.Error
}

func QueryAllJoinedInterviewOfInterview(iid uint) (*[]JoinedInterview, error) {
	var dbJoinedInterview []JoinedInterview
	result := gormDb.Model(&JoinedInterview{}).Where(&JoinedInterview{InterviewID: iid}).Find(&dbJoinedInterview)
	return &dbJoinedInterview, result.Error
}
