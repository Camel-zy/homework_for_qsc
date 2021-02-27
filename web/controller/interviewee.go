package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"

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

	err = model.UpdateInterviewee(&model.Interviewee{InterviewOptions: intOptMarshalled, Status: 3}, vid)
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

// @tags Interviewee
// @summary Get all interviewees in an interview
// @router /interview/interviewee [get]
// @param iid query uint true "Interview ID"
// @produce json
// @success 200 {array} model.Interviewee_
func getAllInterviewees(c echo.Context) error {
	var iid uint
	iid = c.Get("iid").(uint)

	joinedInterview, err := model.QueryAllJoinedInterviewOfInterview(iid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "get interviewees failed",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: *joinedInterview,
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
	return modifyIntervieweeTemplate(4)(c)
}

// @tags Interviewee
// @summary Send an interviewee to next round
// @router /interviewee/next [post]
// @param vid query uint true "Interviewee ID"
// @success 200
func nextInterviewee(c echo.Context) error {
	return modifyIntervieweeTemplate(2)(c)
}

// @tags Interviewee
// @summary Reject an interviewee
// @router /interviewee/reject [post]
// @param vid query uint true "Interviewee ID"
// @success 200
func rejectInterviewee(c echo.Context) error {
	return modifyIntervieweeTemplate(5)(c)
}

// @tags Interviewee
// @summary Get interviewee in event of department
// @router /event/department/interviewee/all [get]
// @param eid query uint true "Event ID"
// @param did query uint true "Department ID"
// @produce json
// @success 200 {array} model.Interviewee_
func getAllIntervieweeInEventOfDepartment(c echo.Context) error {
	eid, eidErr := utils.IsUnsignedInteger(c.QueryParam("eid"))
	did, didErr := utils.IsUnsignedInteger(c.QueryParam("did"))

	if eidErr != nil || didErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "require uint eid and did",
		})
	}

	_, eventErr := model.QueryEventByID(eid)
	if errors.Is(eventErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "event not found",
		})
	}
	_, departmentErr := model.QueryDepartmentById(did)
	if errors.Is(departmentErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "department not found",
		})
	}

	interviewee, intervieweeErr := model.QueryAllIntervieweeByDidAndEid(did, eid)
	if intervieweeErr != nil {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "interviewee not found",
		})
	}
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &interviewee,
	})
}

// @tags Interviewee
// @summary Get interviewee in event of department by round
// @router /event/department/round/interviewee/all [get]
// @param eid query uint true "Event ID"
// @param did query uint true "Department ID"
// @param round query uint true "Round (0为名单总览，1为公海，2为一面，以此类推)"
// @produce json
// @success 200 {array} model.Interviewee_
func getAllIntervieweeByRound(c echo.Context) error {

	did, didErr := utils.IsUnsignedInteger(c.QueryParam("did"))
	eid, eidErr := utils.IsUnsignedInteger(c.QueryParam("eid"))
	round, roundErr := utils.IsUnsignedInteger(c.QueryParam("round"))

	if didErr != nil || eidErr != nil || roundErr != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "require uint did, eid, round and status",
		})
	}

	_, departmentErr := model.QueryDepartmentById(did)
	if departmentErr != nil {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "department not found",
		})
	}

	_, eventErr := model.QueryEventByID(eid)
	if eventErr != nil {
		return c.JSON(http.StatusNotFound, &utils.Error{
			Code: "NOT_FOUND",
			Data: "event not found",
		})
	}

	var interviewee *[]model.Interviewee
	var intervieweeErr error
	if round == 0 {
		interviewee, intervieweeErr = model.QueryAllIntervieweeByDidAndEid(did, eid)
		if intervieweeErr != nil && !errors.Is(intervieweeErr, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusInternalServerError, &utils.Error{
				Code: "INTERNAL_SERVER_ERR",
				Data: "get interviewee failed",
			})
		}
	} else {
		var ret []model.Interviewee
		for i := 1; i < 4; i++ {
			intervieweeTemp, intervieweeErr := model.QueryAllIntervieweeByRoundAndStatus(did, eid, round, uint(i))
			if intervieweeErr != nil && !errors.Is(intervieweeErr, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusInternalServerError, &utils.Error{
					Code: "INTERNAL_SERVER_ERR",
					Data: "get interviewee failed",
				})
			}
			if intervieweeTemp != nil {
				ret = append(ret, *intervieweeTemp...)
			}
		}
		interviewee = &ret
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: &interviewee,
	})
}

// @tags Interviewee
// @summary Get interviewee of admitted status
// @router /event/department/admitted/interviewee [get]
// @param eid query uint true "Event ID"
// @param did query uint true "Department ID"
// @produce json
// @success 200 {object} model.Interviewee_
// func getAllIntervieweeByAdmittedStatus(c echo.Context) error {
// 	var status uint
// 	status = 4 //Status: 4 纳入组织，5 拒绝
// 	did, didErr := utils.IsUnsignedInteger(c.QueryParam("did"))
// 	eid, eidErr := utils.IsUnsignedInteger(c.QueryParam("eid"))

// 	if didErr != nil || eidErr != nil || roundErr != nil {
// 		return c.JSON(http.StatusBadRequest, &utils.Error{
// 			Code: "BAD_REQUEST",
// 			Data: "require uint did, eid and round",
// 		})
// 	}

// 	_, departmentErr := model.QueryDepartmentById(did)
// 	if departmentErr != nil {
// 		return c.JSON(http.StatusNotFound, &utils.Error{
// 			Code: "NOT_FOUND",
// 			Data: "department not found",
// 		})
// 	}

// 	_, eventErr := model.QueryEventByID(eid)
// 	if eventErr != nil {
// 		return c.JSON(http.StatusNotFound, &utils.Error{
// 			Code: "NOT_FOUND",
// 			Data: "event not found",
// 		})
// 	}

// 	interviewee, intervieweeErr := model.QueryIntervieweeByRoundAndStatus(did, eid, round, status)
// 	if intervieweeErr != nil {
// 		return c.JSON(http.StatusNotFound, &utils.Error{
// 			Code: "NOT_FOUND",
// 			Data: "interviewee not found",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, &utils.Error{
// 		Code: "SUCCESS",
// 		Data: &interviewee,
// 	})
// }

// @tags Interviewee
// @summary Get interviewee of rejected status
// @router /event/department/rejected/interviewee [get]
// @param eid query uint true "Event ID"
// @param did query uint true "Department ID"
// @produce json
// @success 200 {object} model.Interviewee_
// func getAllIntervieweeByRejectedStatus(c echo.Context) error {
// 	var status uint
// 	status = 5 //Status: 4 纳入组织，5 拒绝
// 	did, didErr := utils.IsUnsignedInteger(c.QueryParam("did"))
// 	eid, eidErr := utils.IsUnsignedInteger(c.QueryParam("eid"))

// 	if didErr != nil || eidErr != nil || roundErr != nil {
// 		return c.JSON(http.StatusBadRequest, &utils.Error{
// 			Code: "BAD_REQUEST",
// 			Data: "require uint did, eid and round",
// 		})
// 	}

// 	_, departmentErr := model.QueryDepartmentById(did)
// 	if departmentErr != nil {
// 		return c.JSON(http.StatusNotFound, &utils.Error{
// 			Code: "NOT_FOUND",
// 			Data: "department not found",
// 		})
// 	}

// 	_, eventErr := model.QueryEventByID(eid)
// 	if eventErr != nil {
// 		return c.JSON(http.StatusNotFound, &utils.Error{
// 			Code: "NOT_FOUND",
// 			Data: "event not found",
// 		})
// 	}

// 	interviewee, intervieweeErr := model.QueryIntervieweeByRoundAndStatus(did, eid, round, status)
// 	if intervieweeErr != nil {
// 		return c.JSON(http.StatusNotFound, &utils.Error{
// 			Code: "NOT_FOUND",
// 			Data: "interviewee not found",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, &utils.Error{
// 		Code: "SUCCESS",
// 		Data: &interviewee,
// 	})
// }
