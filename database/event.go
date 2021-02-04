package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"gorm.io/gorm/clause"
)

/**************************** Event ****************************/

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

/**************************** Interview ****************************/

func CreateInterview(requestInterview *model.Interview) error {
	result := DB.Create(requestInterview)
	return result.Error
}

func QueryInterviewById(id uint) (*model.Interview, error) {
	var dbInterview model.Interview
	result := DB.First(&dbInterview, id)
	return &dbInterview, result.Error
}

func UpdateInterviewById(requestInterview *model.Interview) error {
	result := DB.Model(&model.Interview{ID: requestInterview.ID}).Updates(requestInterview)
	return result.Error
}

// SELECT * FROM Interview;
func QueryInterviewByIdInEvent(eid uint, iid uint) (*model.Interview, error) {
	var dbInterview model.Interview
	result := DB.Preload(clause.Associations).Where(&model.Interview{ID: iid, EventID: eid}).First(&dbInterview)
	return &dbInterview, result.Error
}

func QueryAllInterviewInEvent(eid uint) (*[]model.Interview, error) {
	var dbInterview []model.Interview
	if findEventError := DB.First(&model.Event{}, eid).Error; findEventError != nil {
		return nil, findEventError
	}
	result := DB.Preload(clause.Associations).Where(&model.Interview{EventID: eid}).Find(&dbInterview)
	return &dbInterview, result.Error
}

/**************************** Joined Interview ****************************/

func UpdateJoinedInterview(id uint, newResult uint) error {
	result := DB.Model(&model.JoinedInterview{ID: id}).Update("result", newResult)
	return result.Error
}
