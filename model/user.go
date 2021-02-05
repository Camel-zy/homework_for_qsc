package model

import (
	"time"
)

type User struct {
	ID           uint       `gorm:"not null;autoIncrement"`
	Name         string     `gorm:"size:40;not null"`
	Nickname     string     `gorm:"size:40"`
	ZJUid        string     `gorm:"column:zju_id;size:10;unique;not null"`
	Mobile       string     `gorm:"size:15"`
	Email        string     `gorm:"size:40"`
	IP           string     `gorm:"size:30"`
	IsSuperuser  uint       `gorm:"default:0"`
	UserAgent    string     `gorm:"size:50"`
	UpdatedTime  time.Time  `gorm:"not null"`
}

func CreateUser(requestUser *User) error {
	result := gormDb.Create(requestUser)
	return result.Error
}

func QueryUserById(id uint) (*User, error) {
	var dbUser User
	result := gormDb.First(&dbUser, id)
	return &dbUser, result.Error
}

func UpdateUserById(requestUser *User) error {
	result := gormDb.Model(&User{ID: requestUser.ID}).Updates(requestUser)
	return result.Error
}

func QueryUserByZJUid(zjuId uint) (*User, error) {
	var dbUser User
	result := gormDb.First(&dbUser, "zju_id = ?", zjuId)
	return &dbUser, result.Error
}

func UpdateUserByZJUid(requestUser *User) error {
	result := gormDb.Model(&User{ZJUid: requestUser.ZJUid}).Updates(requestUser)
	return result.Error
}

// SELECT * FROM users;
func QueryAllUser() (*[]User, error) {
    var dbUser []User
	result := gormDb.Find(&dbUser)
	return &dbUser, result.Error
}
