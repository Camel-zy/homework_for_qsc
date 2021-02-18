package middleware

import (
	"errors"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthMessageTemplate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tid, typeErr := utils.IsUnsignedInteger(c.QueryParam("tid"))
		if typeErr != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "tid need to be an unsigned integer",
			})
		}
		c.Set("tid", tid)
		oid := c.Get("oid").(uint)

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
		if messageTemplate.OrganizationID != oid {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "NO_PRIVILEGE",
				Data: "message template belongs to another organization",
			})
		}

		c.Set("&messageTemplate", messageTemplate) // TODO(TO/GA): lifecycle?
		return next(c)
	}
}
