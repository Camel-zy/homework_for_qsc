package controller

import (
	"errors"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @tags Message
// @summary get message cost and balance
// @description get message cost and balance
// @router /message/cost [get]
// @param oid query uint true "Organization ID"
// @produce json
// @success 200 {object} model.MessageCostAPI
func getMessageCost(c echo.Context) error {
	balance, _ := utils.GetMessageBalance()
	// TODO(TO/GA): error handling

	organization, _ := model.QueryOrganizationById(c.Get("oid").(uint))
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &model.MessageCostAPI{
			Cost:    organization.MessageCost,
			Balance: balance,
		},
	})
}

// @tags Message
// @summary Send a message
// @description send a message
// @router /message [put]
// @accept json
// @param oid query uint true "Organization ID"
// @param data body model.MessageRequest true "Message Information"
// @success 200
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

	department, departErr := model.QueryDepartmentById(messageRequest.DepartmentID)
	if departErr != nil {
		if errors.Is(departErr, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "department not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "send message fail",
		})
	}
	if department.OrganizationID != c.Get("oid").(uint) {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "department and organization don't match",
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
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: nil,
	})
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

// @tags MessageTemplate
// @summary Get all message signs
// @description Get information of all message signs
// @router /messageSign/all [get]
// @produce json
// @success 200 {object} []model.AllMessageSignAPI
func getAllMessageSign(c echo.Context) error {
	signs, _ := utils.GetAllMessageSign()
	// TODO(TO/GA): error handling

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &signs,
	})
}
