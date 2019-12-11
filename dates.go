package frutils

import "time"

func GetMonthsSince(since time.Time) int {
	now := time.Now()

	months := 0
	month := since.Month()

	for since.Before(now) {
		since = since.Add(time.Hour * 24)
		nextMonth := since.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}

func GetDaysBetween(date1 time.Time, date2 time.Time) int {
	return int(date1.Sub(date2).Hours() / 24)
}
