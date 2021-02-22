package controller

import (
	"errors"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
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

	sendErr := utils.SendMessage(messageRequest)
	if sendErr != nil {
		if errors.Is(sendErr, model.ErrInternalError) {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "send message fail",
			})
		}
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: sendErr.Error(),
		})
	}
	return nil
}

func modifyMessageTemplate(callSMSAndModifyDB func(model.MessageTemplateRequest) (*model.MessageTemplate, error)) func(c echo.Context) error {
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

		messageTemplate, msgTplErr := callSMSAndModifyDB(messageTemplateRequest)
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
// @router /messageTemplate [put]
// @param oid query uint true "Organization ID"
// @accept  json
// @param data body model.MessageTemplateRequest true "Message Template Infomation"
// @success 200
func addMessageTemplate(c echo.Context) error {
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

	utils.AddMessageTemplate(c.Get("oid").(uint), &messageTemplateRequest)
	// TODO(TO/GA): error handling

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: nil,
	})
}

// @tags MessageTemplate
// @summary Update a message template
// @description Update a message template
// @router /messageTemplate [post]
// @param oid query uint true "Organization ID"
// @param tid query uint true "Message Template ID"
// @accept  json
// @param data body model.MessageTemplateRequest true "Message Template Infomation"
// @success 200
func setMessageTemplate(c echo.Context) error {
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

	utils.UpdateMessageTemplate(c.Get("tid").(uint), &messageTemplateRequest)
	// TODO(TO/GA): error handling

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: nil,
	})
}

// @tags MessageTemplate
// @summary Get a message template
// @description Get information of a specific message template
// @router /messageTemplate [get]
// @param oid query uint true "Organization ID"
// @param tid query uint true "Message Template ID"
// @produce json
// @success 200 {object} model.MessageTemplateAPI
func getMessageTemplate(c echo.Context) error {
	messageTemplate, _ := utils.GetMessageTemplate(c.Get("&messageTemplate").(*model.MessageTemplate).IDInSMSService)
	// TODO(TO/GA): error handling

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &messageTemplate,
	})
}

// @tags MessageTemplate
// @summary Get all message templates
// @description Get information of all message templates of a specific organization
// @router /messageTemplate/all [get]
// @param oid query uint true "Organization ID"
// @produce json
// @success 200 {object} []model.AllMessageTemplateAPI
func getAllMessageTemplate(c echo.Context) error {
	oid := c.Get("oid").(uint)

	messageTemplates, _ := utils.GetAllMessageTemplate(oid)
	// TODO(TO/GA): error handling

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &messageTemplates,
	})
}
