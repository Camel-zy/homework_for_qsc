package controller

import (
	"encoding/json"
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

func createForm(c echo.Context) error {
	FormRequest := model.FormApi{}
	if err := c.Bind(&FormRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if err := c.Validate(&FormRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if err := json.Valid(FormRequest.Content); err != true {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "wrong format, JSON required",
		})
	}
	if err := model.CreateForm(&FormRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create form fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "create form success",
	})
}

func updateForm(c echo.Context) error {
	FormRequest := model.FormApi{}
	err := c.Bind(&FormRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	fid, err := utils.IsUnsignedInteger(c.QueryParam("fid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid needs to be specified correctly",
		})
	}
	FormRequest.ID = fid
	if err := json.Valid(FormRequest.Content); err != true {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "wrong format, JSON required",
		})
	}
	if err := model.UpdateFormByID(&FormRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "update Form fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "update Form success",
	})
}

func getForm(c echo.Context) error {
	fid, typeErr := utils.IsUnsignedInteger(c.QueryParam("fid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid need to be an unsigned integer",
		})
	}

	form, itvErr := model.QueryFormById(fid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "form not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &form,
	})
}

func updateAnswer(c echo.Context) error {
	fid, typeErr := utils.IsUnsignedInteger(c.QueryParam("fid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid need to be an unsigned integer",
		})
	}
	zjuid, typeErr := utils.IsUnsignedInteger(c.QueryParam("zjuid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "zjuid need to be an unsigned integer",
		})
	}
	eid, typeErr := utils.IsUnsignedInteger(c.QueryParam("eid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid need to be an unsigned integer",
		})
	}
	_, itvErr := model.QueryAnswerByZjuidAndEvent(strconv.FormatUint(uint64(zjuid), 10), eid)
	var content datatypes.JSON
	temp := c.QueryParam("content")
	temp = strings.Replace(temp, "\n", "", -1)
	json.Unmarshal([]byte(temp), &content)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		if aid, err := model.CreateAnswer(content, fid, strconv.FormatUint(uint64(zjuid), 10), eid); err != nil {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "create answer fail",
			})
		} else {
			return c.JSON(http.StatusOK, &utils.Error{
				Code: "SUCCESS",
				Data: aid,
			})
		}
	}
	if err := model.UpdateAnswer(content, strconv.FormatUint(uint64(zjuid), 10), eid); err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "update Answer fail",
		})
	}
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "update answer success",
	})
}

func getAnswer(c echo.Context) error {
	zjuid, typeErr := utils.IsUnsignedInteger(c.QueryParam("zjuid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "zjuid need to be an unsigned integer",
		})
	}
	eid, typeErr := utils.IsUnsignedInteger(c.QueryParam("eid"))
	if typeErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "eid need to be an unsigned integer",
		})
	}
	answer, itvErr := model.QueryAnswerByZjuidAndEvent(strconv.FormatUint(uint64(zjuid), 10), eid)
	if errors.Is(itvErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "answer not found",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &answer,
	})
}
