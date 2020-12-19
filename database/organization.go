package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"git.zjuqsc.com/rop/rop-back-neo/database/proto"
)

func CreateOrganization(requestOrganization *model.Organization) error {
	return proto.Create(requestOrganization)
}

func QueryOrganization(id uint) (*model.Organization, error) {
	var dbOrganization model.Organization
	if result := DB.First(&dbOrganization, "id = ?", id); result.Error != nil {
		return nil, result.Error
	} else {
		return &dbOrganization, nil
	}
}

func UpdateOrganization(requestOrganization *model.Organization) error {
	var dbOrganization model.Organization
	if result := DB.First(&dbOrganization, "name = ?", requestOrganization.Name); result.Error != nil {
		return result.Error
	} else {
		if result := DB.Model(&dbOrganization).Updates(requestOrganization); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}
}
