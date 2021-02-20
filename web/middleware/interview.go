package middleware

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckDepartmentInOrganization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		did, err := utils.IsUnsignedInteger(c.QueryParam("did"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "DepartmentID in form-data need to be an unsigned integer",
			})
		}

		_, err = model.QueryDepartmentByIdUnderOrganization(c.Get("oid").(uint), did)
		if err != nil {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "NOT_FOUND",
				Data: "no such department under current organization which inferred from requested event",
			})
		}

		return next(c)
	}
}

func SetInterview(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var interviewReq model.InterviewResponse

		err := c.Bind(&interviewReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "unknown bad request",
			})
		}

		c.Set("interview", interviewReq)
		return next(c)
	}
}
