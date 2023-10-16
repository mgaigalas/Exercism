package booking

import "time"

const (
	scheduleLayout             string = "1/02/2006 15:04:05"
	passedLayout               string = "January 2, 2006 15:04:05"
	afternoonAppointmentLayout string = "Monday, January 2, 2006 15:04:05"
	descriptionLayout          string = "1/2/2006 15:04:05"
	descriptionDateLayout      string = "Monday, January 2, 2006"
	descriptionTimeLayout      string = "15:04"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	return toTime(scheduleLayout, date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return time.Now().After(toTime(passedLayout, date))
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dateTime := toTime(afternoonAppointmentLayout, date)
	twelfthHour := time.Date(
		dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		12,
		0,
		0,
		0,
		dateTime.Location())
	eighteenthHour := time.Date(
		dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		18,
		0,
		0,
		0,
		dateTime.Location())
	return (dateTime.Equal(twelfthHour) || dateTime.After(twelfthHour)) && dateTime.Before(eighteenthHour)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	dateTime := toTime(descriptionLayout, date)
	return "You have an appointment on " +
		dateTime.Format(descriptionDateLayout) +
		", at " +
		dateTime.Format(descriptionTimeLayout) +
		"."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(
		time.Now().Year(),
		9,
		15,
		0,
		0,
		0,
		0,
		time.UTC)
}

func toTime(layout, date string) time.Time {
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return t
}
