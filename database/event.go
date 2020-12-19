package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"git.zjuqsc.com/rop/rop-back-neo/database/proto"
)

func CreateEvent(requestEvent *model.Event) error {
	return proto.Create(requestEvent)
}

func QueryEvent(id uint) (*model.Event, error) {
	var dbEvent model.Event
	if result := DB.First(&dbEvent, "id = ?", id); result.Error != nil {
		return nil, result.Error
	} else {
		return &dbEvent, nil
	}
}

func UpdateEvent(requestEvent *model.Event) error {
	var dbEvent model.Event
	if result := DB.First(&dbEvent, "name = ?", requestEvent.Name); result.Error != nil {
		return result.Error
	} else {
		if result := DB.Model(&dbEvent).Updates(requestEvent); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}
}
