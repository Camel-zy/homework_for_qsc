package controller

import (
	"errors"
	"net/http"
	"sort"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

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
	} else if len(*newIntention) == 0 {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "no valid intention has been found",
		})
	} else {
		answerRequest.Intention = *newIntention
	}

	if _, err := model.QueryAnswer(fid, zjuid, eid); err == nil {
		if suberr := model.UpdateAnswer(&answerRequest, fid, zjuid, eid); suberr != nil {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "update answer failed",
			})
		}
		return c.JSON(http.StatusOK, &utils.Error{
			Code: "SUCCESS",
			Data: "update answer success",
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
			EventID:      eid,
			AnswerID:     aid,
			DepartmentID: v.DepartmentID,
			IntentRank:   v.IntentRank,
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

	// TODO(TO/GA): Send SMS

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
			if k+1 != int(v.IntentRank) {
				err := errors.New("interview rank must be a string of continuous positive, or all 0 to disable rank")
				return nil, err
			}
		}
	}
	return &newIntention, nil
}