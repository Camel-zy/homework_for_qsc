package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
	"git.zjuqsc.com/rop/rop-back-neo/database/proto"
)

func CreateUser(requestUser *model.User) error {
	return proto.Create(requestUser)
}

func QueryUserById(ID uint) (*model.User, error) {
	var dbUser model.User
	if result := DB.First(&dbUser, "id = ?", ID); result.Error != nil {
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

func QueryUserByZJUid(ZJUid uint) (*model.User, error) {
	var dbUser model.User
	if result := DB.First(&dbUser, "zju_id = ?", ZJUid); result.Error != nil {
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
