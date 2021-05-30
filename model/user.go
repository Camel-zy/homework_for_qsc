package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID           uint       `gorm:"not null;autoIncrement"`
	Name         string     `gorm:"size:40;not null"`
	Nickname     string     `gorm:"size:40"`
	ZJUid        string     `gorm:"column:zju_id;size:10;unique;not null"`
	PassportId   uint       `gorm:"unique"`
	Password 	 []byte 	`gorm:"column:password;not null"`
	Mobile       string     `gorm:"size:15"`
	Email        string     `gorm:"size:40"`
	IP           string     `gorm:"size:30"`
	IsSuperuser  uint       `gorm:"default:0"`
	UserAgent    string     `gorm:"size:50"`
	UpdatedTime  time.Time  `gorm:"autoUpdateTime"`
}

func CreateUser(requestUser *User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword(requestUser.Password, bcrypt.DefaultCost)
	requestUser.Password = hashedPassword
	result := gormDb.Create(requestUser)
	return result.Error
}

func QueryUserById(id uint) (*User, error) {
	var dbUser User
	result := gormDb.First(&dbUser, id)
	return &dbUser, result.Error
}

func QueryUser(requestUser *User) (*User, error) {
	var dbUser User
	result := gormDb.Where(requestUser).First(&dbUser)
	return &dbUser, result.Error
}

func UpdateUserById(requestUser *User) error {
	result := gormDb.Model(&User{ID: requestUser.ID}).Updates(requestUser)
	return result.Error
}

func QueryUserByZJUid(zjuId string) (*User, error) {
	var dbUser User
	result := gormDb.First(&dbUser, "zju_id = ?", zjuId)
	return &dbUser, result.Error
}

func UpdateUserByZJUid(requestUser *User) error {
	result := gormDb.Model(&User{}).Where(&User{ZJUid: requestUser.ZJUid}).Updates(requestUser)
	return result.Error
}

// SELECT * FROM users;
func QueryAllUser() (*[]User, error) {
    var dbUser []User
	result := gormDb.Find(&dbUser)
	return &dbUser, result.Error
}
