package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"git.zjuqsc.com/rop/rop-back-neo/database/proto"
)

func CreateInterview(requestInterview *model.Interview) error {
	return proto.Create(requestInterview)
}

func QueryInterview(ID uint) (*model.Interview, error) {
	var dbInterview model.Interview
	if result := DB.First(&dbInterview, "id = ?", ID); result.Error != nil {
		return nil, result.Error
	} else {
		return &dbInterview, nil
	}
}

func UpdateInterview(requestInterview *model.Interview) error {
	var dbInterview model.Interview
	if result := DB.First(&dbInterview, "name = ?", requestInterview.Name); result.Error != nil {
		return result.Error
	} else {
		if result := DB.Model(&dbInterview).Updates(requestInterview); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}
}
