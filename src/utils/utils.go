package utils

import (
	"strconv"
)

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
