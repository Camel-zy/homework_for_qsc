package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
)

func CreateEvent(requestEvent *model.Event) error {
	result := DB.Create(requestEvent)
	return result.Error
}

func QueryEventById(id uint) (*model.Event, error) {
	var dbEvent model.Event
	result := DB.First(&dbEvent, "id = ?", id)
	return &dbEvent, result.Error
}

func UpdateEventById(requestEvent *model.Event) error {
	result := DB.Model(&model.Event{ID: requestEvent.ID}).Updates(requestEvent)
	return result.Error
}
