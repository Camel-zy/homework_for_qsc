package controller

import (
	"errors"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func addMessage(c echo.Context) error {

	// text, textErr := utils.FillPlaceholder(messageTemplateRequest)
	// if textErr != nil {
	// 	return c.JSON(http.StatusBadRequest, &utils.Error{
	// 		Code: "BAD_REQUEST",
	// 		Data: textErr.Error(),
	// 	})
	// }

	// TODO(TO/GA): finish it
	return nil
}

// @tags Message
// @summary Get a message
// @description Get information of a specific message
// @router /message/{mid} [get]
// @param mid query uint true "Message ID"
// @produce json
// @success 200 {object} model.MessageApi
func getMessage(c echo.Context) error {
	mid, typeErr := utils.IsUnsignedInteger(c.QueryParam("mid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "mid need to be an unsigned integer"},
		)
	}

	message, msgErr := model.QueryMessageById(mid)
	if errors.Is(msgErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "message not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &message,
	})
}

// @tags MessageTemplate
// @summary Add a message template
// @description Add a message template
// @router /message/template [put]
// @param tid body model.MessageTemplateRequest true "Message Template Infomation"
// @produce json
// @success 200 {object} model.MessageTemplateApi
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

	messageTemplate := &model.MessageTemplate{
		Description:    messageTemplateRequest.Description,
		Text:           messageTemplateRequest.Text,
		OrganizationID: messageTemplateRequest.OrganizationID,
	}
	model.CreateMessageTemplate(messageTemplate)

	// TODO(jy): call rop-sms

	var messageTemplateApi model.MessageTemplateApi
	copier.Copy(&messageTemplateApi, &messageTemplate)
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: messageTemplateApi,
	})
}

func setMessageTemplate(c echo.Context) error {
	// TODO(TO/GA): finish it
	return nil
}

// @tags MessageTemplate
// @summary Get a message template
// @description Get information of a specific message template
// @router /message/template/{tid} [get]
// @param tid query uint true "Message Template ID"
// @produce json
// @success 200 {object} model.MessageTemplateApi
func getMessageTemplate(c echo.Context) error {
	tid, typeErr := utils.IsUnsignedInteger(c.QueryParam("tid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "tid need to be an unsigned integer",
		})
	}

	messageTemplate, msgTplErr := model.QueryMessageTemplateById(tid)
	if errors.Is(msgTplErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "message not found",
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
// @success 200 {object} []model.AllMessageTemplateApi
func getAllMessageTemplate(c echo.Context) error {
	oid, typeErr := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid need to be an unsigned integer",
		})
	}

	messageTemplates, msgTplsErr := model.QueryAllMessageTemplateInOrganization(oid)
	if errors.Is(msgTplsErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "organization not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &messageTemplates,
	})
}
