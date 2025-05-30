package datetimeService

import (
	"errors"
	"time"
)

// GetDateInt returns an integer representation of the date in YYYYMMDD format.
func GetDateInt(dateTime *time.Time) (int, error) {
	if dateTime != nil {
		return dateTime.Year()*10000 + int(dateTime.Month())*100 + dateTime.Day(), nil
	}
	return 0, errors.New("DateTime cannot be nil")
}

// GetTimeInt returns an integer representation of the time in HHMMSS format.
func GetTimeInt(dateTime *time.Time) (int, error) {
	if dateTime != nil {
		return dateTime.Hour()*10000 + dateTime.Minute()*100 + dateTime.Second(), nil
	}
	return 0, errors.New("DateTime cannot be nil")
}
