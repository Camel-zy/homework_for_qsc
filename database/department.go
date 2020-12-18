package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"git.zjuqsc.com/rop/rop-back-neo/database/proto"
)

func CreateDepartment(requestDepartment *model.Department) error {
	return proto.Create(requestDepartment)
}

func QueryDepartment(ID uint) (*model.Department, error) {
	var dbDepartment model.Department
	if result := DB.First(&dbDepartment, "id = ?", ID); result.Error != nil {
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
