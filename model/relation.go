package model

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

//DeleteEventHasForm required in the nearest future
