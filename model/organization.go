package model

import "time"

type Organization struct {
	ID          uint      `gorm:"not null;autoIncrement;primaryKey"`
	Name        string    `gorm:"size:40;not null"`
	Description string    `gorm:"size:200"`
	UpdateTime  time.Time `gorm:"not null"`
}

/*
This model is still under discussion
*/
type OrganizationHasUser struct {
	ID             uint
	Role           uint // 0 department admin, 1 organization admin
	UserId         uint
	OrganizationId uint
	DepartmentId   uint
	UpdateTime     time.Time
}

func CreateOrganization(requestOrganization *Organization) error {
	result := gormDb.Create(requestOrganization)
	return result.Error
}

func QueryOrganizationById(id uint) (*Organization, error) {
	var dbOrganization Organization
	result := gormDb.First(&dbOrganization, id)
	return &dbOrganization, result.Error
}

func UpdateOrganizationById(requestOrganization *Organization) error {
	result := gormDb.Model(&Organization{ID: requestOrganization.ID}).Updates(requestOrganization)
	return result.Error
}

func QueryAllOrganization(uid uint) (*[]Organization, error) {
	var dbOrganizationIds []OrganizationHasUser
	gormDb.Select("organization_id").Where(&OrganizationHasUser{UserId: uid}).Find(&dbOrganizationIds)

	var organizationIds []uint
	organizationIdsHelperMap := make(map[uint] bool)

	for _ , v := range dbOrganizationIds {
		if _, ok := organizationIdsHelperMap[v.OrganizationId]; !ok {
			organizationIdsHelperMap[v.OrganizationId] = true
			organizationIds = append(organizationIds, v.OrganizationId)
		}
	}

	var dbOrganization []Organization
	result := gormDb.Find(&dbOrganization, organizationIds)
	return &dbOrganization, result.Error
}
