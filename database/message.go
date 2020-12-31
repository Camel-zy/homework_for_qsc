package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
)

func CreateMessage(requestMessage *model.Message)  error {
    result := DB.Create(requestMessage)
    return result.Error
}

func QueryMessageById (id uint) (*model.Message,error){
    var dbMessage model.Message
    result := DB.First(&dbMessage, id)
    return &dbMessage, result.Error
}

func UpdateMessageById(requestMessage *model.Message) error {
	result := DB.Model(&model.Message{ID: requestMessage.ID}).Updates(requestMessage)
	return result.Error
}