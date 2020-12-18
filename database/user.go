package database

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID          int    `gorm:"not null;autoIncrement"`
	Name        string `gorm:"size:40;not null"`
	Nickname    string `gorm:"size:40"`
	ZJUid       string `gorm:"size:10;unique;not null"`
	Mobile      string `gorm:"size:15"`
	Email       string `gorm:"size:40"`
	IP          string `gorm:"size:30"`
	IsSuperuser int    `gorm:"default:0"`
	UserAgent   string `gorm:"size:50"`
	UpdatedTime string `gorm:"size:30;not null"`
}

func createUser(requestUser *User) error {

	DB.Create(requestUser)
	return nil
}

func queryUserById(ID uint) (*User, error) {
	var result User
	if err := DB.First(&result, "ID = ?", ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else {
		return &result, nil
	}
}

func updateUserById(requestUser *User) error {
	var result User
	if err := DB.First(&result, "name = ?", requestUser.Name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	} else {
		if result := DB.Model(&result).Updates(requestUser); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}

}

func queryUserByZJUid(ZJUid uint) (*User, error) {
	var result User
	if err := DB.First(&result, "ZJUid = ?", ZJUid).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else {
		return &result, nil
	}
}

func UpdateUserByZJUid(requestUser *User) error {
	var result User
	if err := DB.First(&result, "ZJUid = ?", requestUser.ZJUid).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	} else {
		if result := DB.Model(&result).Updates(requestUser); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}
}
