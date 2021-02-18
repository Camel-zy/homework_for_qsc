package model

import (
	"github.com/jinzhu/copier"
	"time"
)

type Event struct {
	ID             uint         `gorm:"not null;autoIncrement;primaryKey"`
	Name           string       `gorm:"size:40;not null"`
	Description    string       `gorm:"size:200"`
	OrganizationID uint         `gorm:"not null"`
	Organization   Organization // FOREIGN KEY (OrganizationID) REFERENCES Organization(OrganizationID)
	Status         uint         `gorm:"default:0"` // 0 disabled, 1 testing, 2 running
	OtherInfo      string       `gorm:"size:200"`
	StartTime      time.Time    `gorm:"size:30;not null"`
	EndTime        time.Time    `gorm:"size:30;not null"`
	UpdatedTime    time.Time    `gorm:"autoUpdateTime"`
}

type EventApi struct {
	ID             uint
	Name           string    `json:"Name" validate:"required"`
	Description    string    `json:"Description"`
	OrganizationID uint      `json:"OrganizationID" validate:"required"`
	Status         uint      `json:"Status"` // 0 disabled (default), 1 testing, 2 running
	OtherInfo      string    `json:"OtherInfo"`
	StartTime      time.Time `json:"StartTime" validate:"required"` // request string must be in RFC 3339 format
	EndTime        time.Time `json:"EndTime" validate:"required"` // request string must be in RFC 3339 format
}

func CreateEvent(requestEvent *EventApi) error {
	dbEvent := Event{}
	copier.Copy(&dbEvent, requestEvent)
	result := gormDb.Create(&dbEvent)
	return result.Error
}

func UpdateEventByID(requestEvent *EventApi) error {
	dbEvent := Event{}
	copier.Copy(&dbEvent, requestEvent)
	result := gormDb.Model(&Event{ID: requestEvent.ID}).Updates(&dbEvent)
	return result.Error
}

func QueryEventByID(id uint) (*EventApi, error) {
	var dbEvent EventApi
	result := gormDb.Model(&Event{}).First(&dbEvent, id)
	return &dbEvent, result.Error
}

func QueryEventByIDInOrganization(oid uint, eid uint) (*EventApi, error) {
	var dbEvent EventApi
	result := gormDb.Model(&Event{}).Where(&Event{ID: eid, OrganizationID: oid}).First(&dbEvent)
	return &dbEvent, result.Error
}

func QueryAllEventInOrganization(oid uint) (*[]Brief, error) {
	var dbEvent []Brief
	if findOrganizationError := gormDb.First(&Organization{}, oid).Error; findOrganizationError != nil {
		return nil, findOrganizationError
	}
	result := gormDb.Model(&Event{}).Where(&Event{OrganizationID: oid}).Find(&dbEvent)
	return &dbEvent, result.Error
}
