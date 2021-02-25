package controller

import (
	"encoding/json"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
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

	err = model.UpdateInterviewee(&model.Interviewee{InterviewOptions: intOptMarshalled}, vid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "update interviewee failed",
		})
	}

	// TODO(TO/GA): Send SMS

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "",
	})
}
