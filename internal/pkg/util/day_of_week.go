package util

import "time"

func GetDayOfWeek(date time.Time) string {
	switch date.Weekday() {
	case time.Monday:
		return "T2"
	case time.Tuesday:
		return "T3"
	case time.Wednesday:
		return "T4"
	case time.Thursday:
		return "T5"
	case time.Friday:
		return "T6"
	case time.Saturday:
		return "T7"
	case time.Sunday:
		return "CN"
	default:
		return ""
	}
}
