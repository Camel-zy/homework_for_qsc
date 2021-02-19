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
		c.Set("eid", eid)
		event, err := model.QueryEventByID(eid)
		if err != nil {
			return c.JSON(http.StatusForbidden, &utils.Error{
				Code: "NO_PRIVILEGE",
				Data: "no privilege to access the information in this organization",
			})
		}
		c.Set("oid", event.OrganizationID)
		return next(c)
	}
}
