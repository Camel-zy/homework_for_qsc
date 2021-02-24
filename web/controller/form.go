package controller

import (
	"encoding/json"
	"errors"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"sort"
)

// @tags Form
// @summary Create a form
// @description Create a form
// @router /form [put]
// @accept  json
// @param data body model.FormApi_ true "Form information"
// @success 200
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

// @tags Form
// @summary Update a form
// @description Update a form
// @router /form [post]
// @accept  json
// @param data body model.FormApi_ true "Form information"
// @success 200
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

// @tags Form
// @summary Get a form
// @description Get a form
// @router /form [get]
// @param fid query uint true "Form ID"
// @success 200 {object} model.Form_
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

// @tags Answer
// @summary Create an answer
// @description Create an answer
// @router /answer [put]
// @accept  json
// @param fid query uint true "Form ID"
// @param zjuid query string true "ZJU ID"
// @param eid query uint true "Event ID"
// @param data body model.AnswerRequest_ true "Answer information"
// @success 200
func updateAnswer(c echo.Context) error {
	var fid, eid uint
	var zjuid string
	err := echo.QueryParamsBinder(c).
		MustUint("fid", &fid).
		MustString("zjuid", &zjuid).
		MustUint("eid", &eid).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid, zjuid, eid needs to be an unsigned integer",
		})
	}

	answerRequest := model.AnswerRequest{}
	if err := c.Bind(&answerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	} else if err = c.Validate(&answerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	}
	if newIntention, err := SortIntention(&answerRequest.Intention); err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	} else {
		answerRequest.Intention = *newIntention
	}

	if _, err := model.QueryAnswer(fid, zjuid, eid); err == nil {
		// TODO: write update logic instead of this
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "the answer already exists",
		})
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create answer failed due to some error while creating",
		})
	}

	aid, err := model.CreateAnswer(&answerRequest, fid, zjuid, eid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create answer failed",
		})
	}

	for _, v := range answerRequest.Intention {
		interviewee := model.Interviewee{
			EventID: eid,
			AnswerID: aid,
			DepartmentID: v.DepartmentID,
			IntentRank: v.IntentRank,
		}
		_, err := model.CreateInterviewee(&interviewee)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "create interviewee failed",
			})
		}
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: aid,
	})
}

func SortIntention(origArray *[]model.Intention) (*[]model.Intention, error) {
	var newIntention []model.Intention
	hasRank := false
	for _, v := range *origArray {
		if v.DepartmentID != 0 {
			newIntention = append(newIntention, v)
		}
		if v.IntentRank != 0 {
			hasRank = true
		}
	}

	if hasRank {
		sort.Slice(newIntention, func(i, j int) bool {
			return newIntention[i].IntentRank < newIntention[j].IntentRank
		})
		for k, v := range newIntention {
			if k + 1 != int(v.IntentRank) {
				err := errors.New("interview rank must be a string of continuous positive, or all 0 to disable rank")
				return nil, err
			}
		}
	}
	return &newIntention, nil
}
