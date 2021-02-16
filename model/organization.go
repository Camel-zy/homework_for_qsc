package model

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

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
	result := gormDb.Model(&Organization{}).First(&dbOrganization, id)
	return &dbOrganization, result.Error
}

func UpdateOrganizationById(requestOrganization *Organization) error {
	result := gormDb.Model(&Organization{ID: requestOrganization.ID}).Updates(requestOrganization)
	return result.Error
}


func QueryAllOrganization(uid uint) (*[]Brief, error) {
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

	var dbOrganization []Brief
	result := gormDb.Model(&Organization{}).Find(&dbOrganization, organizationIds)
	return &dbOrganization, result.Error
}

func UserIsInOrganization(uid uint, oid uint) bool {
	err := gormDb.First(&OrganizationHasUser{}, OrganizationHasUser{UserId: uid, OrganizationId: oid}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else if err != nil {
		logrus.WithField("uid", uid).WithField("oid", oid).Warn(err)
		return false
	}
	return true
}
