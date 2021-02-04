package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"gorm.io/gorm/clause"
)

/**************************** User ****************************/

func CreateUser(requestUser *model.User) error {
	result := DB.Create(requestUser)
	return result.Error
}

func QueryUserById(id uint) (*model.User, error) {
	var dbUser model.User
	result := DB.First(&dbUser, id)
	return &dbUser, result.Error
}

func UpdateUserById(requestUser *model.User) error {
	result := DB.Model(&model.User{ID: requestUser.ID}).Updates(requestUser)
	return result.Error
}

func QueryUserByZJUid(zjuId uint) (*model.User, error) {
	var dbUser model.User
	result := DB.First(&dbUser, "zju_id = ?", zjuId)
	return &dbUser, result.Error
}

func UpdateUserByZJUid(requestUser *model.User) error {
	result := DB.Model(&model.User{ZJUid: requestUser.ZJUid}).Updates(requestUser)
	return result.Error
}

// SELECT * FROM users;
func QueryAllUser() (*[]model.User, error) {
    var dbUser []model.User
	result := DB.Find(&dbUser)
	return &dbUser, result.Error
}

/**************************** Department ****************************/

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

/**************************** Organization ****************************/

func CreateOrganization(requestOrganization *model.Organization) error {
	result := DB.Create(requestOrganization)
	return result.Error
}

func QueryOrganizationById(id uint) (*model.Organization, error) {
	var dbOrganization model.Organization
	result := DB.First(&dbOrganization, id)
	return &dbOrganization, result.Error
}

func UpdateOrganizationById(requestOrganization *model.Organization) error {
	result := DB.Model(&model.Organization{ID: requestOrganization.ID}).Updates(requestOrganization)
	return result.Error
}

// SELECT * FROM organizations;
func QueryAllOrganization() (*[]model.Organization, error) {
	var dbOrganization []model.Organization
	result := DB.Find(&dbOrganization)
	return &dbOrganization, result.Error
}

