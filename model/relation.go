package model

import "gorm.io/gorm"

type EventHasForm struct {
	ID      uint `gorm:"not null;autoIncrement;primaryKey"`
	FormID  uint `gorm:"not null"`
	EventID uint `gorm:"not null"`
}

func CreateEventHasForm(fid uint, eid uint) (EventHasForm, error) {
	dbEventHasForm := EventHasForm{}
	dbEventHasForm.FormID = fid
	dbEventHasForm.EventID = eid
	result := gormDb.Create(&dbEventHasForm)
	return dbEventHasForm, result.Error
}
func QueryEventHasForm(fid uint, eid uint) (*EventHasForm, error) {
	var dbEventHasForm EventHasForm
	result := gormDb.Where(&EventHasForm{FormID: fid, EventID: eid}).First(&dbEventHasForm)
	return &dbEventHasForm, result.Error
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
