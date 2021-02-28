package model

import (
	"time"

	"gorm.io/gorm/clause"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	IntervieweeStart           = 0 // 不要使用
	IntervieweeTimeChecked     = 1 // 已确认本轮面试时间或正在面试
	IntervieweeRoundPassed     = 2 // 本轮通过
	IntervieweeMessageSent     = 3 // 已发送下轮分配短信，面试者未进行选择
	IntervieweeOrgAdmitted     = 4 // 纳入组织
	IntervieweeOrgRejected     = 5 // 拒绝
	IntervieweeNextRoundNoTime = 6 // 已发送下轮分配短信，但面试者选择没有合适的下轮面试时间
	IntervieweeEnd             = 7 // 不要使用
)

// 一个志愿一条记录
// Don't forget to modify Interviewee_ if you modify this
type Interviewee struct {
	ID               uint      `gorm:"not null;autoIncrement;primaryKey"`
	UUID             uuid.UUID `gorm:"not null;type:uuid"`
	EventID          uint      `gorm:"not null"`
	AnswerID         uint      `gorm:"not null"`
	Answer           Answer
	DepartmentID     uint           `gorm:"not null"`           // 志愿部门
	IntentRank       uint           `gorm:"not null;default:0"` // 第几志愿
	Round            uint           `gorm:"not null;default:1"` // 公海为1，一面为2，以此类推
	InterviewOptions datatypes.JSON // 发送选择面试场次的短信用
	Status           uint           `gorm:"not null; default:2"` // 1 已确认本轮面试时间/正在面试，2 本轮通过，3 已发送下轮分配短信，4 纳入组织，5 拒绝, 6 面试者选择没有合适的下轮面试时间
}

type IntervieweeRequest struct {
	InterviewOptions []uint `json:"InterviewOptions"`
}

type JoinedInterview struct {
	ID            uint      `gorm:"not null;autoIncrement;primaryKey"`
	InterviewID   uint      `gorm:"not null"`
	IntervieweeID uint      `gorm:"not null"`
	UpdatedTime   time.Time `gorm:"autoUpdateTime"`
	Deleted       gorm.DeletedAt
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

func UpdateIntervieweeByUuid(interviewee *Interviewee, uuid uuid.UUID) error {
	result := gormDb.Model(&Interviewee{}).
		Where(&Interviewee{UUID: uuid}).Updates(interviewee)
	return result.Error
}

func QueryIntervieweeById(id uint) (*Interviewee, error) {
	var dbInterviewee Interviewee
	result := gormDb.First(&dbInterviewee, id)
	return &dbInterviewee, result.Error
}

func QueryIntervieweeByUUID(uuid uuid.UUID) (*Interviewee, error) {
	var dbInterviewee Interviewee
	result := gormDb.Where(&Interviewee{UUID: uuid}).First(&dbInterviewee)
	return &dbInterviewee, result.Error
}

func UpdateJoinedInterview(id uint, newResult uint) error {
	result := gormDb.Model(&JoinedInterview{ID: id}).Update("result", newResult)
	return result.Error
}

func CreateJoinedInterview(iid, vid uint) error {
	result := gormDb.Create(&JoinedInterview{InterviewID: iid, IntervieweeID: vid})
	return result.Error
}

func DeleteJoinedInterviewByIidAndVid(iid, vid uint) error {
	var dbJoinedInterview JoinedInterview
	result := gormDb.Model(&JoinedInterview{}).Where(&JoinedInterview{InterviewID: iid, IntervieweeID: vid}).Find(&dbJoinedInterview)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNoRowsAffected
	}
	result = gormDb.Delete(&dbJoinedInterview)
	return result.Error
}

func QueryAllJoinedInterviewOfInterview(iid uint) (*[]JoinedInterview, error) {
	var dbJoinedInterview []JoinedInterview
	result := gormDb.Model(&JoinedInterview{}).Where(&JoinedInterview{InterviewID: iid}).Find(&dbJoinedInterview)
	return &dbJoinedInterview, result.Error
}

func QueryAllIntervieweeByDidAndEid(did, eid uint) (*[]Interviewee, error) {
	var dbInterviewee []Interviewee
	result := gormDb.Preload(clause.Associations).Model(&Interviewee{}).Where(&Interviewee{DepartmentID: did, EventID: eid}).Find(&dbInterviewee)
	return &dbInterviewee, result.Error
}

func QueryAllIntervieweeByStatus(did, eid, status uint) (*[]Interviewee, error) {
	var dbInterviewee []Interviewee
	result := gormDb.Preload(clause.Associations).Model(&Interviewee{}).Where(&Interviewee{DepartmentID: did, EventID: eid, Status: status}).Find(&dbInterviewee)
	return &dbInterviewee, result.Error
}

func QueryAllIntervieweeByRoundAndStatus(did, eid, round, status uint) (*[]Interviewee, error) {
	var dbInterviewee []Interviewee
	result := gormDb.Preload(clause.Associations).Model(&Interviewee{}).Where(&Interviewee{DepartmentID: did, EventID: eid, Round: round, Status: status}).Find(&dbInterviewee)
	return &dbInterviewee, result.Error
}

func QueryNumberOfIntervieweesInInterviewByInterviewID(vid uint) (int64, error) {
	var number int64
	result := gormDb.Model(&JoinedInterview{}).Where(&JoinedInterview{InterviewID: vid}).Count(&number)
	return number, result.Error
}
