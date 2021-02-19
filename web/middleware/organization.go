package middleware

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TempOrganization struct {
	OrganizationID uint `json:"OrganizationID"`
}

func AuthOrganization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		getOrganizationId, ok := c.Get("getOrganizationIdFunc").(func(c echo.Context) (uint, bool))
		if !ok {
			getOrganizationId = getOrganizationIdFromParam
		}

		oid, ok := getOrganizationId(c)
		if !ok {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "OrganizationID (oid) needs to be set correctly",
			})
		}

		c.Set("oid", oid)
		uid := c.Get("uid").(uint)
		if !model.UserIsInOrganization(uid, oid) {
			return c.JSON(http.StatusForbidden, &utils.Error{
				Code: "NO_PRIVILEGE",
				Data: "no privilege to access the information in this organization",
			})
		}
		return next(c)
	}
}

func getOrganizationIdFromParam(c echo.Context) (oid uint, ok bool) {
	oid, err := utils.IsUnsignedInteger(c.QueryParam("oid"))
	if err == nil {
		ok = true
	}
	return
}

func getOrganizationIdFromContext(c echo.Context) (oid uint, ok bool) {
	oid, ok = c.Get("oid").(uint)
	return
}

func getOrganizationIdFromForm(c echo.Context) (oid uint, ok bool) {
	temp := TempOrganization{}
	err := c.Bind(&temp)
	if err == nil {
		ok = true
		oid = temp.OrganizationID
	}
	return
}
