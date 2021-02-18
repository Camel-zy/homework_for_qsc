package middleware

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"net/http"
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