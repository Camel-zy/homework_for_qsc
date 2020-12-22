package database

import "git.zjuqsc.com/rop/rop-back-neo/database/model"

func UpdateJoinedInterview(id uint, newResult uint) error {
	result := DB.Model(&model.JoinedInterview{ID: id}).Update("result", newResult)
	return result.Error
}
