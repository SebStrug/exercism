package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
// input string is in the format "7/25/2019 13:45:00"
func Schedule(date string) time.Time {
	form := "1/02/2006 15:04:05"
	t, _ := time.Parse(form, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const dateFormat = "January 2, 2006 15:04:05"
	dateObj, _ := time.Parse(dateFormat, date)
	return time.Now().After(dateObj)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const dateFormat = "Monday, January 2, 2006 15:04:05"
	dateObj, _ := time.Parse(dateFormat, date)
	return dateObj.Hour() >= 12 && dateObj.Hour() <= 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const dateFormat = "1/2/2006 15:04:05"
	dateObj, _ := time.Parse(dateFormat, date)
	dateStr := dateObj.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", dateStr)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	const dateFormat = "2006-01-02"
	desiredFormat := fmt.Sprintf("%v-09-15", time.Now().Year())
	dateObj, _ := time.Parse(dateFormat, desiredFormat)
	return dateObj
}
