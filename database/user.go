package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
)

func CreateUser(requestUser *model.User) error {
	result := DB.Create(requestUser)
	return result.Error
}

func QueryUserById(id uint) (*model.User, error) {
	var dbUser model.User
	result := DB.First(&dbUser, "id = ?", id)
	return &dbUser, result.Error
}

func UpdateUserById(requestUser *model.User) error {
	result := DB.Model(&model.User{ID: requestUser.ID}).Updates(requestUser)
	return result.Error
}

func QueryUserByZJUid(zjuId uint) (*model.User, error) {
	var dbUser model.User
	result := DB.First(&dbUser, "zju_id = ?", zjuId)
	return &dbUser, result.Error
}

func UpdateUserByZJUid(requestUser *model.User) error {
	result := DB.Model(&model.User{ZJUid: requestUser.ZJUid}).Updates(requestUser)
	return result.Error
}
