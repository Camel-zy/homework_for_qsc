package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
)

func CreateDepartment(requestDepartment *model.Department) error {
	result := DB.Create(requestDepartment)
	return result.Error
}

func QueryDepartmentById(oid string, did string) (*model.Department, error) {
	var dbDepartment model.Department
	result := DB.First(&dbDepartment, "id = ?", did)
	return &dbDepartment, result.Error
}

func UpdateDepartmentById(requestDepartment *model.Department) error {
	result := DB.Model(&model.Department{ID: requestDepartment.ID}).Updates(requestDepartment)
	return result.Error
}

// SELECT * FROM department;
func QueryAllDepartment() (*[]model.Department, error) {
	var dbDepartment []model.Department
	result := DB.Find(&dbDepartment)
	return &dbDepartment, result.Error
}
