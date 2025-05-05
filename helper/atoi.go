package helper

import (
	"strconv"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
