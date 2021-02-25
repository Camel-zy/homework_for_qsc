package controller

import (
	"encoding/json"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// @tags Interviewee
// @summary Update the interview option of a interviewee
// @description The frontend is required to send multiple iid in a JSON array,
// @description which are the optional interviews for an interviewee to choose from
// @router /interviewee/options [post]
// @param vid query uint true "Interviewee ID"
// @param data body model.IntervieweeRequest true "The InterviewOption field is required, in JSON array"
// @produce json
func updateIntervieweeOptions(c echo.Context) error {
	var vid uint
	err := echo.QueryParamsBinder(c).MustUint("vid", &vid).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "vid needs to be an unsigned integer",
		})
	}

	var intervieweeRequest model.IntervieweeRequest
	err = c.Bind(&intervieweeRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: err.Error(),
		})
	} else if len(intervieweeRequest.InterviewOptions) == 0 {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "requires a uint InterviewOptions array",
		})
	}

	intOptMarshalled, err := json.Marshal(intervieweeRequest.InterviewOptions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "error occurs when marshalling JSON",
		})
	}

	err = model.UpdateInterviewee(&model.Interviewee{InterviewOptions: intOptMarshalled, SentMessage: 2}, vid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "update interviewee failed",
		})
	}

	_, err = utils.SendMessage(vid, 1)
	if err != nil {
		logrus.Errorf("send reject message fail(vid=%v): %v", vid, err)
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "send message fail",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "",
	})
}

func modifyIntervieweeTemplate(newStatus uint) func(echo.Context) error {
	return func(c echo.Context) error {
		vid := c.Get("vid").(uint)
		err := model.UpdateInterviewee(&model.Interviewee{Status: newStatus}, vid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "update interviewee fail",
			})
		}
		if newStatus == 4 {
			_, err = utils.SendMessage(vid, 3)
			if err != nil {
				logrus.Errorf("send reject message fail(vid=%v): %v", vid, err)
				return c.JSON(http.StatusInternalServerError, &utils.Error{
					Code: "INTERNAL_SERVER_ERR",
					Data: "send message fail",
				})
			}
		}
		return c.JSON(http.StatusOK, &utils.Error{
			Code: "SUCCESS",
			Data: "update interviewee success",
		})
	}
}

// @tags Interviewee
// @summary Admit an interviewee
// @router /interviewee/admit [post]
// @param vid query uint true "Interviewee ID"
// @success 200
func admitInterviewee(c echo.Context) error {
	return modifyIntervieweeTemplate(3)(c);
}

// @tags Interviewee
// @summary Send an interviewee to next round
// @router /interviewee/next [post]
// @param vid query uint true "Interviewee ID"
// @success 200
func nextInterviewee(c echo.Context) error {
	return modifyIntervieweeTemplate(1)(c);
}

// @tags Interviewee
// @summary Reject an interviewee
// @router /interviewee/reject [post]
// @param vid query uint true "Interviewee ID"
// @success 200
func rejectInterviewee(c echo.Context) error {
	return modifyIntervieweeTemplate(4)(c);
}
