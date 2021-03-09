package controller

import (
	"errors"
	"net/http"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/labstack/echo/v4"
)

// @tags Relation
// @summary Create a relation
// @description Create a relation
// @router /relation/event/form [put]
// @param fid query uint true "Form ID"
// @param eid query uint true "Event ID"
// @success 200
func createEventHasForm(c echo.Context) error {
	var fid, eid uint
	err := echo.QueryParamsBinder(c).
		MustUint("fid", &fid).
		MustUint("eid", &eid).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid, eid needs to be an unsigned integer",
		})
	}
	if _, err := model.QueryEventHasForm(fid, eid); err == nil {
		return c.JSON(http.StatusOK, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "relation has already been stored",
		})
	}
	relation, err := model.CreateEventHasForm(fid, eid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "create relation failed",
		})
	}
	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: relation,
	})
}

// @tags Relation
// @summary Validate relation
// @description Validate a relation
// @router /relation/event/form [get]
// @param fid query uint true "Form ID"
// @param eid query uint true "Event ID"
func validateEventHasForm(c echo.Context) error {
	var fid, eid uint
	err := echo.QueryParamsBinder(c).
		MustUint("fid", &fid).
		MustUint("eid", &eid).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid, eid needs to be unsigned integers",
		})
	}
	_, err = model.QueryEventHasForm(fid, eid)
	if err != nil {
		return utils.GetApiReturnNotFoundOrInternalError(c, "relation", err)
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "relation exists",
	})
}

// @tags Relation
// @summary Delete relation
// @router /relation/event/form [delete]
// @param fid query uint true "Form ID"
// @param eid query uint true "Event ID"
func deleteEventHasForm(c echo.Context) error {
	var fid, eid uint
	err := echo.QueryParamsBinder(c).
		MustUint("fid", &fid).
		MustUint("eid", &eid).
		BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.Error{
			Code: "BAD_REQUEST",
			Data: "fid, eid needs to be unsigned integers",
		})
	}

	err = model.DeleteEventHasForm(fid, eid)
	if err != nil {
		if errors.Is(err, model.ErrNoRowsAffected) {
			return c.JSON(http.StatusNotFound, &utils.Error{
				Code: "NOT_FOUND",
				Data: "relation not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, &utils.Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: "delete relation failed",
		})
	}

	return c.JSON(http.StatusOK, &utils.Error{
		Code: "SUCCESS",
		Data: "delete relation success",
	})

}
