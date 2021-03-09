package utils

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Error struct {
	Code  string       `json:"code"`
	Data  interface{}  `json:"data"`
}

func IsUnsignedInteger(input string) (uint, error) {
	if convertedInt, err := strconv.ParseUint(input, 10, 64); err != nil {
		return 0, errors.New("not an unsigned integer")
	} else {
		return uint(convertedInt), nil
	}
}

func GetApiReturnNotFoundOrInternalError(c echo.Context, name string, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = c.JSON(http.StatusNotFound, &Error{
			Code: "NOT_FOUND",
			Data: fmt.Sprintf("%s not found", name),
		})
		return errors.New(fmt.Sprintf("%s not found", name))
	} else {
		_ = c.JSON(http.StatusInternalServerError, &Error{
			Code: "INTERNAL_SERVER_ERR",
			Data: fmt.Sprintf("get %s failed", name),
		})
		return errors.New(fmt.Sprintf("get %s failed", name))
	}
}

func ReturnNotFound(c echo.Context, name string) error {
	_ = c.JSON(http.StatusNotFound, &Error{
		Code: "NOT_FOUND",
		Data: fmt.Sprintf("%s not found", name),
	})
	return errors.New(fmt.Sprintf("%s not found", name))
}
