package middleware

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckDepartmentInOrganization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var did uint
		err := echo.QueryParamsBinder(c).MustUint("did", &did).BindError()
		if err != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "did needs to be an unsigned integer",
			})
		}

		oid, ok := c.Get("oid").(uint)  // assume it has been set by the previous procedure
		if !ok {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "oid hasn't been set properly",
			})
		}

		department, err := model.QueryDepartmentById(did)
		if err != nil || department.OrganizationID != oid {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "NOT_FOUND",
				Data: "no such department under current organization which inferred from requested event",
			})
		}

		return next(c)
	}
}

func SetInterviewOrganization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var iid uint
		err := echo.QueryParamsBinder(c).MustUint("iid", &iid).BindError()
		if err != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "iid needs to be an unsigned integer",
			})
		}

		interview, err := model.QueryInterviewByIDWithPreload(iid)
		if err != nil {
			return c.JSON(http.StatusForbidden, &utils.Error{
				Code: "NO_PRIVILEGE",
				Data: "no privilege to access the information in this interview",
			})
		} else if interview.Event.OrganizationID == 0 {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "oid hasn't been set properly",
			})
		}

		c.Set("oid", interview.Event.OrganizationID)

		return next(c)
	}
}
