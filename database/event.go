package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"gorm.io/gorm/clause"
)

func CreateEvent(requestEvent *model.Event) error {
	result := DB.Create(requestEvent)
	return result.Error
}

func QueryEventById(id uint) (*model.Event, error) {
	var dbEvent model.Event
	result := DB.First(&dbEvent, id)
	return &dbEvent, result.Error
}

func UpdateEventById(requestEvent *model.Event) error {
	result := DB.Model(&model.Event{ID: requestEvent.ID}).Updates(requestEvent)
	return result.Error
}

// SELECT * FROM event;
func QueryEventByIdOfOrganization(oid uint, eid uint) (*model.Event, error) {
	var dbEvent model.Event
	result := DB.Preload(clause.Associations).Where(&model.Event{ID: eid, OrganizationID: oid}).First(&dbEvent)
	return &dbEvent, result.Error
}

func QueryAllEventOfOrganization(oid uint) (*[]model.Event, error) {
	var dbEvent []model.Event
	if findOrganizationError := DB.First(&model.Organization{}, oid).Error; findOrganizationError != nil {
		return nil, findOrganizationError
	}
	result := DB.Preload(clause.Associations).Where(&model.Event{OrganizationID: oid}).Find(&dbEvent)
	return &dbEvent, result.Error
}