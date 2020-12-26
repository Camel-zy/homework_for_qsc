package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
)

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
