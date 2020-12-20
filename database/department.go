package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"git.zjuqsc.com/rop/rop-back-neo/database/utils"
)

func CreateDepartment(requestDepartment *model.Department) error {
	return utils.Create(DB, requestDepartment)
}

func QueryDepartment(id uint) (*model.Department, error) {
	var dbDepartment model.Department
	if result := DB.First(&dbDepartment, "id = ?", id); result.Error != nil {
		return nil, result.Error
	} else {
		return &dbDepartment, nil
	}
}

func UpdateDepartment(requestDepartment *model.Department) error {
	var dbDepartment model.Department
	if result := DB.First(&dbDepartment, "name = ?", requestDepartment.Name); result.Error != nil {
		return result.Error
	} else {
		if result := DB.Model(&dbDepartment).Updates(requestDepartment); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}
}
