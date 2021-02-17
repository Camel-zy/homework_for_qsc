package middleware

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func AuthOrganization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		oid, err := utils.IsUnsignedInteger(c.QueryParam("oid"))
		if err != nil {
			var ok bool
			oid, ok = c.Get("oid").(uint)
			if !ok {
				return c.JSON(http.StatusBadRequest, &utils.Error{
					Code: "BAD_REQUEST",
					Data: "oid need to be an unsigned integer",
				})
			}
		}
		c.Set("oid", oid)
		uid := c.Get("uid").(uint)
		if !model.UserIsInOrganization(uid, oid) {
			return c.JSON(http.StatusUnauthorized, &utils.Error{
				Code: "NO_PRIVILEGE",
				Data: "no privilege to access the information in this organization",
			})
		}
		return next(c)
	}
}

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

func SetEventStruct(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.FormValue("Name") == "" {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "field not enough in form-data",
			})
		}

		startTime, errStart := time.Parse(time.RFC3339, c.FormValue("StartTime"))
		endTime, errEnd := time.Parse(time.RFC3339, c.FormValue("EndTime"))
		if errStart != nil || errEnd != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "time string in form-data must be in RFC 3339 format and set correctly",
			})
		}

		eventReq := model.Event{
			Name: c.FormValue("Name"),
			Description: c.FormValue("Description"),
			OrganizationID: c.Get("oid").(uint),
			OtherInfo: c.FormValue("OtherInfo"),
			StartTime: startTime,
			EndTime: endTime,
		}
		if c.FormValue("Status") != "" {
			statusCode, err := utils.IsUnsignedInteger(c.FormValue("Status"))
			if err != nil {
				return c.JSON(http.StatusBadRequest, &utils.Error{
					Code: "BAD_REQUEST",
					Data: "status needs to be an unsigned integer",
				})
			} else if statusCode > 2 {
				return c.JSON(http.StatusBadRequest, &utils.Error{
					Code: "BAD_REQUEST",
					Data: "the value in status field is illegal",
				})
			}
			eventReq.Status = statusCode
		}

		c.Set("event", eventReq)
		return next(c)
	}
}
