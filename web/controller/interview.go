package controller

import (
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func getInterviewInEvent(c echo.Context) error {
	eid, errEid := utils.IsUnsignedInteger(c.QueryParam("eid"))
	iid, errIid := utils.IsUnsignedInteger(c.QueryParam("iid"))

	if errEid != nil || errIid != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: "BAD_REQUEST", Data: "eid and iid need to be an unsigned integer"})
	}

	_, evtErr := model.QueryEventById(eid)
	if errors.Is(evtErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "EVT_NOT_FOUND", Data: "event not found"})
	}

	interview, itvErr := model.QueryInterviewByIdInEvent(eid, iid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "ITV_NOT_FOUND", Data: "interview not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &interview})
}

func getAllInterviewInEvent(c echo.Context) error {
	eid, typeErr := utils.IsUnsignedInteger(c.QueryParam("eid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{Code: "BAD_REQUEST", Data: "eid need to be an unsigned integer"})
	}

	interviews, itvErr := model.QueryAllInterviewInEvent(eid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound){
		return c.JSON(http.StatusNotFound, &utils.Error{Code: "EVT_NOT_FOUND", Data: "event not found"})
	}

	return c.JSON(http.StatusOK, &utils.Error{Code: "SUCCESS", Data: &interviews})
}
