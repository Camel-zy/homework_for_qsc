package middleware

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type TempOrganization struct {
	OrganizationID uint `json:"OrganizationID"`
}

// oid must be set in advance in echo context
func AuthOrganization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		/*
		getOrganizationId, ok := c.Get("getOrganizationIdFunc").(func(c echo.Context) (uint, bool))
		if !ok {
			getOrganizationId = getOrganizationIdFromParam
		}

		oid, ok := getOrganizationId(c)
		 */

		oid, ok := c.Get("oid").(uint)
		if !ok {
			logrus.Error("oid hasn't been set properly in echo context")
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "oid hasn't been set in the context",
			})
		}

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

func GetOrganizationIdFromParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var oid uint
		err := echo.QueryParamsBinder(c).
			MustUint("oid", &oid).
			BindError()
		if err != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "oid needs to be set correctly",
			})
		}
		/*
		oid, err := utils.IsUnsignedInteger(c.QueryParam("oid"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "oid needs to be set correctly",
			})
		}
		 */
		c.Set("oid", oid)
		return next(c)
	}
}
