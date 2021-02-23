package model

import (
	"time"
)

type Department struct {
	ID             uint         `gorm:"not null;autoIncrement;primaryKey"`
	Name           string       `gorm:"size:40;not null"`
	OrganizationID uint         `gorm:"not null"`
	Organization   Organization // FOREIGN KEY (OrganizationID) REFERENCES Organization(OrganizationID)
	SerialNumInOrg uint         `gorm:"not null"`
	Description    string       `gorm:"size:200"`
	MessageCost    float32      `gorm:"not null;default:0"`
	UpdateTime     time.Time    `gorm:"autoUpdateTime"`
}

type DepartmentApi struct {
	ID             uint
	Name           string
	OrganizationID uint
	Description    string
	MessageCost    float32
}

type JoinedDepartment struct {
	ID           uint      `gorm:"not null;autoIncrement;primaryKey"`
	UserID       uint      `gorm:"not null"`
	DepartmentID uint      `gorm:"not null"`
	Privilege    uint      `gorm:"default:2"`
	JoinedTime   time.Time `gorm:"not null"`
	UpdateTime   time.Time `gorm:"autoUpdateTime"`
}

func CreateDepartment(requestDepartment *Department) error {
	result := gormDb.Create(requestDepartment)
	return result.Error
}

func QueryDepartmentById(id uint) (*Department, error) {
	var dbDepartment Department
	result := gormDb.First(&dbDepartment, id)
	return &dbDepartment, result.Error
}

func UpdateDepartmentById(requestDepartment *Department) error {
	result := gormDb.Model(&Department{ID: requestDepartment.ID}).Updates(requestDepartment)
	return result.Error
}

// SELECT * FROM department;
func QueryAllDepartment() (*[]Department, error) {
	var dbDepartment []Department
	result := gormDb.Find(&dbDepartment)
	return &dbDepartment, result.Error
}

func QueryDepartmentByIdUnderOrganization(oid uint, did uint) (*DepartmentApi, error) {
	var dbDepartment DepartmentApi
	result := gormDb.Model(&Department{}).Where(&Department{ID: did, OrganizationID: oid}).First(&dbDepartment)
	return &dbDepartment, result.Error
}

func QueryAllDepartmentInOrganization(oid uint) (*[]Brief, error) {
	var dbDepartment []Brief

	/* we need to tell the user whether the organization can be found */
	if findOrganizationError := gormDb.First(&Organization{}, oid).Error; findOrganizationError != nil {
		return nil, findOrganizationError
	}

	/* then, the organization exists */
	result := gormDb.Model(&Department{}).Where(&Department{OrganizationID: oid}).Find(&dbDepartment)
	return &dbDepartment, result.Error
}
