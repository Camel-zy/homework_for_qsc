package model

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Organization struct {
	ID          uint      `gorm:"not null;autoIncrement;primaryKey"`
	Name        string    `gorm:"size:40;not null"`
	Description string    `gorm:"size:200"`
	MessageCost float64   `gorm:"not null;default:0"`
	UpdateTime  time.Time `gorm:"autoUpdateTime"`
}

type OrganizationApi struct {
	ID          uint
	Name        string
	Description string
	MessageCost float64
}

/*
This model is still under discussion
*/
type OrganizationHasUser struct {
	ID             uint
	Role           uint `gorm:"default:1"` // 1 department admin, 2 organization admin
	PassportId     uint
	OrganizationId uint
	DepartmentId   uint
	UpdateTime     time.Time `gorm:"autoUpdateTime"`
}

func CreateOrganization(requestOrganization *Organization) error {
	result := gormDb.Create(requestOrganization)
	return result.Error
}

func QueryOrganizationById(id uint) (*OrganizationApi, error) {
	var dbOrganization OrganizationApi
	result := gormDb.Model(&Organization{}).First(&dbOrganization, id)
	return &dbOrganization, result.Error
}

func UpdateOrganizationById(requestOrganization *Organization) error {
	result := gormDb.Model(&Organization{ID: requestOrganization.ID}).Updates(requestOrganization)
	return result.Error
}

func QueryAllOrganization(uid uint) (*[]Brief, error) {
	var dbOrganizationIds []OrganizationHasUser
	gormDb.Select("organization_id").Where(&OrganizationHasUser{PassportId: uid}).Find(&dbOrganizationIds)

	var organizationIds []uint
	organizationIdsHelperMap := make(map[uint]bool)

	for _, v := range dbOrganizationIds {
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
	err := gormDb.First(&OrganizationHasUser{}, OrganizationHasUser{PassportId: uid, OrganizationId: oid}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else if err != nil {
		logrus.WithField("uid", uid).WithField("oid", oid).Warn(err)
		return false
	}
	return true
}
