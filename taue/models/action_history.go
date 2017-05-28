package models

import (
	"time"
)

type ActionHistory struct {
	actions []Action
	year    int
	month   time.Month
	day     int
}

func (actionHistory ActionHistory) turfColor() string {

	var count int

	for _, action := range actionHistory.actions {
		count += action.Count
	}

	switch {
	case count == 0:
		return ""
	case 1 <= count, count <= 3:
		return ""
	case 4 <= count, count <= 6:
		return ""
	case 7 <= count, count <= 9:
		return ""
	case 10 <= count, count <= 12:
		return ""
	case 13 <= count, count <= 15:
		return ""
	case 16 <= count, count <= 18:
		return ""
	case 19 <= count, count <= 20:
		return ""
	case 20 < count:
		return ""
	default:
		return ""
	}

}

func (actionHistory ActionHistory) isEqualDate(date time.Time) bool {
	return actionHistory.year == date.Year() && actionHistory.month == date.Month() && actionHistory.day == date.Day()
}
