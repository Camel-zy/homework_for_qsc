package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
)

func CreateInterview(requestInterview *model.Interview) error {
	result := DB.Create(requestInterview)
	return result.Error
}

func QueryInterviewById(id uint) (*model.Interview, error) {
	var dbInterview model.Interview
	result := DB.First(&dbInterview, "id = ?", id)
	return &dbInterview, result.Error
}

func UpdateInterviewById(requestInterview *model.Interview) error {
	result := DB.Model(&model.Interview{ID: requestInterview.ID}).Updates(requestInterview)
	return result.Error
}
