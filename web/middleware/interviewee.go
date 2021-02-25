package middleware

import (
	"errors"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthInterviewee(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var vid uint
		err := echo.QueryParamsBinder(c).MustUint("vid", &vid).BindError()
		if err != nil {
			return c.JSON(http.StatusBadRequest, &utils.Error{
				Code: "BAD_REQUEST",
				Data: "vid needs to be an unsigned integer",
			})
		}

		interviewee, err := model.QueryIntervieweeById(vid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusBadRequest, &utils.Error{
					Code: "BAD_REQUEST",
					Data: "no such interviewee",
				})
			}
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "find interviewee fail",
			})
		}

		department, err := model.QueryDepartmentById(interviewee.DepartmentID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "check authentication fail",
			})
		}

		if ok := model.UserIsInOrganization(c.Get("uid").(uint), department.OrganizationID); !ok {
			return c.JSON(http.StatusForbidden, &utils.Error{
				Code: "NO_PRIVILEGE",
				Data: "no privilege to access the information of the interviewee",
			})
		}

		c.Set("vid", vid)
		return next(c)
	}
}
