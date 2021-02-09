package model

import (
	"gorm.io/gorm/clause"
	"time"
)

type Event struct {
	ID              uint          `gorm:"not null;autoIncrement;primaryKey"`
	Name            string        `gorm:"size:40;not null"`
	Description     string        `gorm:"size:200"`
	OrganizationID  uint          `gorm:"not null"`
	Organization    Organization  // FOREIGN KEY (OrganizationID) REFERENCES Organization(OrganizationID)
	Status          uint          `gorm:"default:1"`
	OtherInfo       string        `gorm:"size:200"`
	StartTime       time.Time     `gorm:"size:30;not null"`
	EndTime         time.Time     `gorm:"size:30;not null"`
	UpdatedTime     time.Time     `gorm:"not null"`
}

func CreateEvent(requestEvent *Event) error {
	result := gormDb.Create(requestEvent)
	return result.Error
}

func QueryEventById(id uint) (*Event, error) {
	var dbEvent Event
	result := gormDb.First(&dbEvent, id)
	return &dbEvent, result.Error
}

func UpdateEventById(requestEvent *Event) error {
	result := gormDb.Model(&Event{ID: requestEvent.ID}).Updates(requestEvent)
	return result.Error
}

// SELECT * FROM event;
func QueryEventByIdOfOrganization(oid uint, eid uint) (*Event, error) {
	var dbEvent Event
	result := gormDb.Preload(clause.Associations).Where(&Event{ID: eid, OrganizationID: oid}).First(&dbEvent)
	return &dbEvent, result.Error
}

func QueryAllEventOfOrganization(oid uint) (*[]Event, error) {
	var dbEvent []Event
	if findOrganizationError := gormDb.First(&Organization{}, oid).Error; findOrganizationError != nil {
		return nil, findOrganizationError
	}
	result := gormDb.Preload(clause.Associations).Where(&Event{OrganizationID: oid}).Find(&dbEvent)
	return &dbEvent, result.Error
}