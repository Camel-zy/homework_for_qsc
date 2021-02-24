// CAUTIOUS: is to be DELETED
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

// // @tags Message
// // @summary get message cost and balance
// // @description get message cost and balance
// // @router /message/cost [get]
// // @param oid query uint true "Organization ID"
// // @produce json
// // @success 200 {object} model.MessageCostAPI
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

func sendMessage(BindAndValidate func(c echo.Context) (*model.MessageRequest, uint, error)) func(c echo.Context) error {
	return func(c echo.Context) error {
		messageRequest, messageTemplateID, msgReqErr := BindAndValidate(c)
		if msgReqErr != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: msgReqErr.Error(),
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

		text, sendErr := utils.SendMessage(messageRequest, messageTemplateID)
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
			Data: text,
		})
	}
}

// // @tags Message
// // @summary send form confirm message
// // @description send form confirm message
// // @router /message/form [put]
// // @param oid query uint true "Organization ID"
// // @param data body model.SendUserMessageRequest true "Message Information"
// // @success 200 {string} MessageText "MessageText"
func sendFormConfirmMessage(c echo.Context) error {
	return sendMessage(func(c echo.Context) (*model.MessageRequest, uint, error) {
		var messageRequest model.SendUserMessageRequest
		if bindErr := c.Bind(&messageRequest); bindErr != nil {
			return nil, 0, bindErr
		}

		if validateErr := c.Validate(&messageRequest); validateErr != nil {
			return nil, 0, validateErr
		}

		var req model.MessageRequest
		copier.Copy(&req, &messageRequest)
		return &req, 0, nil
	})(c)
}

// // @tags Message
// // @summary send interview select message
// // @description send interview select message
// // @router /message/interview/select [put]
// // @param oid query uint true "Organization ID"
// // @param data body model.SendInterviewMessageRequest true "Message Information"
// // @success 200 {string} MessageText "MessageText"
func sendInterviewSelectMessage(c echo.Context) error {
	return sendMessage(func(c echo.Context) (*model.MessageRequest, uint, error) {
		var messageRequest model.SendInterviewMessageRequest
		if bindErr := c.Bind(&messageRequest); bindErr != nil {
			return nil, 0, bindErr
		}

		if validateErr := c.Validate(&messageRequest); validateErr != nil {
			return nil, 0, validateErr
		}

		var req model.MessageRequest
		copier.Copy(&req, &messageRequest)
		return &req, 1, nil
	})(c)
}

// // @tags Message
// // @summary send interview confirm message
// // @description send interview confirm message
// // @router /message/interview/confirm [put]
// // @param oid query uint true "Organization ID"
// // @param data body model.SendInterviewMessageRequest true "Message Information"
// // @success 200 {string} MessageText "MessageText"
func sendInterviewConfirmMessage(c echo.Context) error {
	return sendMessage(func(c echo.Context) (*model.MessageRequest, uint, error) {
		var messageRequest model.SendInterviewMessageRequest
		if bindErr := c.Bind(&messageRequest); bindErr != nil {
			return nil, 0, bindErr
		}

		if validateErr := c.Validate(&messageRequest); validateErr != nil {
			return nil, 0, validateErr
		}

		var req model.MessageRequest
		copier.Copy(&req, &messageRequest)
		return &req, 2, nil
	})(c)
}

// // @tags Message
// // @summary send reject message
// // @description send reject message
// // @router /message/reject [put]
// // @param oid query uint true "Organization ID"
// // @param data body model.SendUserMessageRequest true "Message Information"
// // @success 200 {string} MessageText "MessageText"
func sendRejectMessage(c echo.Context) error {
	return sendMessage(func(c echo.Context) (*model.MessageRequest, uint, error) {
		var messageRequest model.SendUserMessageRequest
		if bindErr := c.Bind(&messageRequest); bindErr != nil {
			return nil, 0, bindErr
		}

		if validateErr := c.Validate(&messageRequest); validateErr != nil {
			return nil, 0, validateErr
		}

		var req model.MessageRequest
		copier.Copy(&req, &messageRequest)
		return &req, 3, nil
	})(c)
}
