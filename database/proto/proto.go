package proto

import "git.zjuqsc.com/rop/rop-back-neo/database"

func Create(value interface{}) error {
	if result := database.DB.Create(value); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
