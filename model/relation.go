package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EventHasForm struct {
	ID      uint `gorm:"not null;autoIncrement;primaryKey"`
	FormID  uint `gorm:"not null"`
	Form    Form
	EventID uint `gorm:"not null"`
	Deleted gorm.DeletedAt
}

type EventHasFormResponse struct {
	ID      uint
	FormID  uint
	Form    Form_
	EventID uint
}

func CreateEventHasForm(fid uint, eid uint) (*EventHasForm, error) {
	dbEventHasForm := EventHasForm{}
	dbEventHasForm.FormID = fid
	dbEventHasForm.EventID = eid
	result := gormDb.Create(&dbEventHasForm)
	return &dbEventHasForm, result.Error
}
func QueryEventHasForm(fid uint, eid uint) (*EventHasForm, error) {
	var dbEventHasForm EventHasForm
	result := gormDb.Where(&EventHasForm{FormID: fid, EventID: eid}).First(&dbEventHasForm)
	return &dbEventHasForm, result.Error
}

func QueryEventHasFormByEid(eid uint) (*[]EventHasForm, error) {
	var dbEventHasForm []EventHasForm
	result := gormDb.Preload(clause.Associations).Where(&EventHasForm{EventID: eid}).Find(&dbEventHasForm)
	return &dbEventHasForm, result.Error
}

func DeleteEventHasForm(fid, eid uint) error {
	result := gormDb.
		Where(&EventHasForm{FormID: fid, EventID: eid}).
		Delete(&EventHasForm{})
	return result.Error
}

func QueryAllFormByEventID(eid uint) (*[]Form, error) {
	var dbEventHasForm []EventHasForm
	result := gormDb.Model(&EventHasForm{}).Where(&EventHasForm{EventID: eid}).Find(&dbEventHasForm)
	form := make([]Form, 0)
	for _, v := range dbEventHasForm {
		tmp, err := QueryFormById(v.FormID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, ErrInternalError
			}
			return nil, err
		}
		form = append(form, *tmp)
	}
	return &form, result.Error
}

//DeleteEventHasForm required in the nearest future
