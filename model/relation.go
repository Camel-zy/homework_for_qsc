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
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

//DeleteEventHasForm required in the nearest future
