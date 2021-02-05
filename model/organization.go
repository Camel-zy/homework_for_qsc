package model

import "time"

type Organization struct {
	ID           uint       `gorm:"not null;autoIncrement;primaryKey"`
	Name         string     `gorm:"size:40;not null"`
	Description  string     `gorm:"size:200"`
	UpdateTime   time.Time  `gorm:"not null"`
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

// SELECT * FROM organizations;
func QueryAllOrganization() (*[]Organization, error) {
	var dbOrganization []Organization
	result := gormDb.Find(&dbOrganization)
	return &dbOrganization, result.Error
}
