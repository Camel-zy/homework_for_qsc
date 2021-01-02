package utils

import (
	"errors"
	"strconv"
)

type Error struct {
	Code  string       `json:"code"`
	Data  interface{}  `json:"data"`
}

func IsUnsignedInteger(input string) (uint, error) {
	if convertedInt, err := strconv.Atoi(input); err != nil {
		return 0, errors.New("not an integer")
	} else if convertedInt < 0 {
		return 1, errors.New("not an unsigned integer")
	} else {
		return uint(convertedInt), nil
	}
}