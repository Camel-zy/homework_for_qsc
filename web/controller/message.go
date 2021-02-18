package controller

import (
	"errors"
	"fmt"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// @tags Message
// @summary Send a message
// @description send a message
// @router /message [put]
// @accept json
// @param data body model.MessageRequest true "Message Information"
// @produce json
// @success 200 {object} model.MessageAPI
func addMessage(c echo.Context) error {
	var messageRequest model.MessageRequest
	if bindErr := c.Bind(&messageRequest); bindErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: bindErr.Error(),
		})
	}

	if validateErr := c.Validate(&messageRequest); validateErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: validateErr.Error(),
		})
	}

	mapping, mapErr := utils.GenerateMap(messageRequest)
	if mapErr != nil {
		if errors.Is(mapErr, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "fill placeholders fail due to the lack of information",
			})
		}
		logrus.Errorln("fill placeholders fail", mapErr.Error())
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "fill placeholders fail",
		})
	}

	fmt.Println(mapping)
	// TODO(jy): call rop-sms
	// TODO(TO/GA): insert into db
	return nil
}

// @tags Message
// @summary Get a message
// @description Get information of a specific message
// @router /message/{mid} [get]
// @param mid query uint true "Message ID"
// @produce json
// @success 200 {object} model.MessageAPI
func getMessage(c echo.Context) error {
	mid, typeErr := utils.IsUnsignedInteger(c.QueryParam("mid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "mid need to be an unsigned integer"},
		)
	}

	message, msgErr := model.QueryMessageById(mid)
	if msgErr != nil {
		if errors.Is(msgErr, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "NOT_FOUND",
				Data: "message not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "get message fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &message,
	})
}

func modifyMessageTemplate(modifyDB func(model.MessageTemplateRequest) (*model.MessageTemplate, error)) func(c echo.Context) error {
	return func(c echo.Context) error {
		var messageTemplateRequest model.MessageTemplateRequest
		if bindErr := c.Bind(&messageTemplateRequest); bindErr != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: bindErr.Error(),
			})
		}

		if validateErr := c.Validate(&messageTemplateRequest); validateErr != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: validateErr.Error(),
			})
		}

		messageTemplate, msgTplErr := modifyDB(messageTemplateRequest)
		if msgTplErr != nil {
			if msgTplErr == model.ErrInternalError {
				return c.JSON(http.StatusInternalServerError, &utils.Error{
					Code: "INTERNAL_SERVER_ERR",
					Data: "modify message template fail",
				})
			}
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: msgTplErr.Error(),
			})
		}

		// TODO(jy): call rop-sms

		var MessageTemplateAPI model.MessageTemplateAPI
		copier.Copy(&MessageTemplateAPI, &messageTemplate)
		return c.JSON(http.StatusOK, &utils.Error{
			Code: "SUCCESS",
			Data: MessageTemplateAPI,
		})
	}
}

// @tags MessageTemplate
// @summary Add a message template
// @description Add a message template
// @router /message/template [put]
// @accept  json
// @param data body model.MessageTemplateRequest true "Message Template Infomation"
// @produce json
// @success 200 {object} model.MessageTemplateAPI
func addMessageTemplate(c echo.Context) error {
	return modifyMessageTemplate(func(messageTemplateRequest model.MessageTemplateRequest) (*model.MessageTemplate, error) {
		messageTemplate := &model.MessageTemplate{
			Description:    messageTemplateRequest.Description,
			Text:           messageTemplateRequest.Text,
			OrganizationID: messageTemplateRequest.OrganizationID,
		}
		model.CreateMessageTemplate(messageTemplate)
		return messageTemplate, nil
	})(c)
}

// @tags MessageTemplate
// @summary Update a message template
// @description Update a message template
// @router /message/template/{tid} [post]
// @param tid query uint true "Message Template ID"
// @accept  json
// @param data body model.MessageTemplateRequest true "Message Template Infomation"
// @produce json
// @success 200 {object} model.MessageTemplateAPI
func setMessageTemplate(c echo.Context) error {
	return modifyMessageTemplate(func(messageTemplateRequest model.MessageTemplateRequest) (*model.MessageTemplate, error) {
		tid, typeErr := utils.IsUnsignedInteger(c.QueryParam("tid"))
		if typeErr != nil {
			return nil, errors.New("tid need to be an unsigned integer")
		}

		messageTemplate := &model.MessageTemplate{
			ID:             tid,
			Description:    messageTemplateRequest.Description,
			Text:           messageTemplateRequest.Text,
			OrganizationID: messageTemplateRequest.OrganizationID,
			Status:         0,
		}
		updateErr := model.UpdateMessageTemplateById(messageTemplate)
		if updateErr != nil {
			if updateErr == model.ErrNoRowsAffected {
				return nil, errors.New("message template not found")
			}
			return nil, model.ErrInternalError
		}
		return messageTemplate, nil
	})(c)
}

// @tags MessageTemplate
// @summary Get a message template
// @description Get information of a specific message template
// @router /message/template/{tid} [get]
// @param tid query uint true "Message Template ID"
// @produce json
// @success 200 {object} model.MessageTemplateAPI
func getMessageTemplate(c echo.Context) error {
	tid, typeErr := utils.IsUnsignedInteger(c.QueryParam("tid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "tid need to be an unsigned integer",
		})
	}

	messageTemplate, msgTplErr := model.QueryMessageTemplateById(tid)
	if msgTplErr != nil {
		if errors.Is(msgTplErr, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "NOT_FOUND",
				Data: "message template not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "get message template fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &messageTemplate,
	})
}

// @tags MessageTemplate
// @summary Get all message templates
// @description Get information of all message templates of a specific organization
// @router /message/template/all/{oid} [get]
// @param oid query uint true "Organization ID"
// @produce json
// @success 200 {object} []model.AllMessageTemplateAPI
func getAllMessageTemplate(c echo.Context) error {
	oid, typeErr := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid need to be an unsigned integer",
		})
	}

	messageTemplates, msgTplsErr := model.QueryAllMessageTemplateInOrganization(oid)
	if msgTplsErr != nil {
		if errors.Is(msgTplsErr, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "NOT_FOUND",
				Data: "organization not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "get all message templates fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &messageTemplates,
	})
}
