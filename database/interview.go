package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"gorm.io/gorm/clause"
)

func CreateInterview(requestInterview *model.Interview) error {
	result := DB.Create(requestInterview)
	return result.Error
}

func QueryInterviewById(id uint) (*model.Interview, error) {
	var dbInterview model.Interview
	result := DB.First(&dbInterview, id)
	return &dbInterview, result.Error
}

func UpdateInterviewById(requestInterview *model.Interview) error {
	result := DB.Model(&model.Interview{ID: requestInterview.ID}).Updates(requestInterview)
	return result.Error
}

// SELECT * FROM Interview;
func QueryInterviewByIdInEvent(eid uint, iid uint) (*model.Interview, error) {
	var dbInterview model.Interview
	result := DB.Preload(clause.Associations).Where(&model.Interview{ID: iid, EventID: eid}).First(&dbInterview)
	return &dbInterview, result.Error
}

func QueryAllInterviewInEvent(eid uint) (*[]model.Interview, error) {
	var dbInterview []model.Interview
	if findEventError := DB.First(&model.Event{}, eid).Error; findEventError != nil {
		return nil, findEventError
	}
	result := DB.Preload(clause.Associations).Where(&model.Interview{EventID: eid}).Find(&dbInterview)
	return &dbInterview, result.Error
}