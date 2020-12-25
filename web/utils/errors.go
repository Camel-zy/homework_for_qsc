package utils

import "strconv"

type Error struct {
	Code         int     `json:"code"`
	Description  string  `json:"description"`
}

func IsUnsignedInteger(input string) bool {
	if convertedInt, err := strconv.Atoi(input); err != nil {
		return false
	} else if convertedInt < 0 {
		return false
	}
	return true
}
