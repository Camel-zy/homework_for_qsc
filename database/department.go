package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"gorm.io/gorm/clause"
)

func CreateDepartment(requestDepartment *model.Department) error {
	result := DB.Create(requestDepartment)
	return result.Error
}

func QueryDepartmentById(id uint) (*model.Department, error) {
	var dbDepartment model.Department
	result := DB.First(&dbDepartment, id)
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

func QueryDepartmentByIdUnderOrganization(oid uint, did uint) (*model.Department, error) {
	var dbDepartment model.Department
	result := DB.Preload(clause.Associations).Where(&model.Department{ID: did, OrganizationID: oid}).First(&dbDepartment)
	return &dbDepartment, result.Error
}

func QueryAllDepartmentUnderOrganization(oid uint) (*[]model.Department, error) {
	var dbDepartment []model.Department

	/* we need to tell the user whether the organization can be found */
	if findOrganizationError := DB.First(&model.Organization{}, oid).Error; findOrganizationError != nil {
		return nil, findOrganizationError
	}

	/* then, the organization exists */
	result := DB.Preload(clause.Associations).Where(&model.Department{OrganizationID: oid}).Find(&dbDepartment)
	return &dbDepartment, result.Error
}
