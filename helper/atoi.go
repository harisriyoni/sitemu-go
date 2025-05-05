package helper

import (
	"strconv"
)

// Atoi mengonversi string ke integer, jika gagal akan mengembalikan 0
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
