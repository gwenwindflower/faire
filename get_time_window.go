package main

import "time"

func GetTimeWindow() (time.Time, time.Time) {
	// We want dates to be local to the user
	// Truncate assumes UTC, so we convert to local time
	now := time.Now().In(time.Local)
	// This works better than Truncate for our localized time
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	aWeekBeforeActiveHabitDay := today.AddDate(0, 0, -7)
	return today, aWeekBeforeActiveHabitDay
}
