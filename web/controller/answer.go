package controller

import (
	"encoding/json"
	"errors"
	"fmt"
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
// @description Create an answer, while also create interviewees
// @router /answer [put]
// @accept  json
// @param fid query uint true "Form ID"
// @param zjuid query string true "ZJU ID"
// @param eid query uint true "Event ID"
// @param data body model.AnswerRequest_ true "Answer information"
// @success 200
func createAnswer(c echo.Context) error {
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
	answer := model.Answer{}
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
	newIntention, err := HandleIntention(&answerRequest.Intention, eid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	} else if len(*newIntention) == 0 {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "no valid intention has been found",
		})
	}

	jsonIntention, err := json.Marshal(*newIntention)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "marshal intention failed",
		})
	}

	answer.Name = answerRequest.Name
	answer.Intention = jsonIntention
	answer.Mobile = answerRequest.Mobile
	answer.Content =  answerRequest.Content

	if _, err := model.QueryAnswer(fid, zjuid, eid); err == nil {
		if subErr := model.UpdateAnswer(&answer, fid, zjuid, eid); subErr != nil {
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

	aid, err := model.CreateAnswer(&answer, fid, zjuid, eid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create answer failed",
		})
	}

	var oneOfTheIntervieweeID uint // 具体是哪一个不重要，只需要其中一个提取 Interviewee 信息就好

	for _, v := range *newIntention {
		interviewee := model.Interviewee{
			EventID:      eid,
			AnswerID:     aid,
			DepartmentID: v.DepartmentID,
			IntentRank:   v.IntentRank,
		}
		oneOfTheIntervieweeID, err = model.CreateInterviewee(&interviewee)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "create interviewee failed",
			})
		}
	}

	_, err = utils.SendMessage(oneOfTheIntervieweeID, 0)
	if err != nil {
		logrus.Errorf("send form submission confirmation message fail(vid=%v): %v", oneOfTheIntervieweeID, err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "send message fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: aid,
	})
}

// @tags Answer
// @summary Get an answer
// @description Get information of an answer by AnswerID
// @router /answer [get]
// @param aid query uint true "Answer ID"
// @produce json
// @success 200 {object} model.Answer_
func getAnswer(c echo.Context) error {
	var aid uint
	err := echo.QueryParamsBinder(c).MustUint("aid", &aid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "aid needs to be an unsigned integer",
		})
	}

	answer, aswErr := model.QueryAnswerByID(aid)
	if errors.Is(aswErr, gorm.ErrRecordNotFound) {
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

func HandleIntention(origArray *[]model.IntentionRequest, eid uint) (*[]model.Intention, error) {
	var newIntention []model.Intention
	hasRank := false
	event, err := model.QueryEventByID(eid)
	if err != nil {
		return nil, errors.New("event not found")
	}
	brief, err := model.QueryAllDepartmentInOrganization(event.OrganizationID)
	for _, v := range *origArray {
		did, err := FindDidByName(brief, v.DepartmentName)
		if err != nil {
			return nil, err
		}
		newIntention = append(newIntention, model.Intention{DepartmentID: did, IntentRank: v.IntentRank})
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

func FindDidByName(brief *[]model.Brief, name string) (did uint, err error) {
	for _, v := range *brief {
		if name == v.Name {
			return v.ID, nil
		}
	}
	return 0, errors.New(
		fmt.Sprintf(`department named "%s" has not been found in the current event`, name))
}
