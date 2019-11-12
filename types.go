package frutils

import (
	"strconv"
	"strings"
	"time"
)

func ToString(i int) string {
	return strconv.Itoa(i)
}

func ToString64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func ToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func ToBool(s string) bool {
	return strings.TrimSpace(s) == "true" || strings.TrimSpace(s) == "1" || strings.TrimSpace(s) == "TRUE"
}

func DateToString(date time.Time) string {
	return date.Format("2006-01-02")
}

func DateToStringFull(date time.Time) string {
	return date.Format("2006-01-02T15:04:05-07:00")
}

func IntToAlphabetPosition(i int) string {
	return string('A' - 1 + i)
}

/* dates */
func GetDaysBetween(date1 time.Time, date2 time.Time) int {
	return int(date1.Sub(date2).Hours() / 24)
}
