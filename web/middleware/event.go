package middleware

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetEventOrganization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		eid, err := utils.IsUnsignedInteger(c.QueryParam("eid"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "eid need to be an unsigned integer",
			})
		}
		event, err := model.QueryEventByID(eid)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, &utils.Error{
				Code: "NO_PRIVILEGE",
				Data: "no privilege to access the information in this organization",
			})
		}
		c.Set("oid", event.OrganizationID)
		return next(c)
	}
}

func CheckRequiredField(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.FormValue("Name") == "" ||
			c.FormValue("StartTime") == "" ||
			c.FormValue("EndTime") == "" {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "some required fields are not set",
			})
		}
		return next(c)
	}
}

func SetEvent(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var eventReq model.EventApi

		err := c.Bind(&eventReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "unknown bad request",
			})
		}
		if eventReq.Status > 2 {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "the value in status field is illegal",
			})
		}

		c.Set("event", eventReq)
		return next(c)
	}
}
