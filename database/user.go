package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"git.zjuqsc.com/rop/rop-back-neo/database/utils"
)

func CreateUser(requestUser *model.User) error {
	return utils.Create(DB, requestUser)
}

func QueryUserById(id uint) (*model.User, error) {
	var dbUser model.User
	if result := DB.First(&dbUser, "id = ?", id); result.Error != nil {
		return nil, result.Error
	} else {
		return &dbUser, nil
	}
}

func UpdateUserById(requestUser *model.User) error {
	var dbUser model.User
	if result := DB.First(&dbUser, "name = ?", requestUser.Name); result.Error != nil {
		return result.Error
	} else {
		if result := DB.Model(&dbUser).Updates(requestUser); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}
}

func QueryUserByZJUid(zjuId uint) (*model.User, error) {
	var dbUser model.User
	if result := DB.First(&dbUser, "zju_id = ?", zjuId); result.Error != nil {
		return nil, result.Error
	} else {
		return &dbUser, nil
	}
}

func UpdateUserByZJUid(requestUser *model.User) error {
	var dbUser model.User
	if result := DB.First(&dbUser, "zju_id = ?", requestUser.ZJUid); result.Error != nil {
		return result.Error
	} else {
		if result := DB.Model(&dbUser).Updates(requestUser); result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	}
}
