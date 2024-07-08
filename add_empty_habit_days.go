package main

import (
	"sort"
	"time"
)

func AddEmptyHabitDays(habitDays []Habit, start time.Time, end time.Time) []Habit {
	existingDays := make(map[string]bool)
	for _, habit := range habitDays {
		dateKey := habit.Date.Format("2006-01-02")
		existingDays[dateKey] = true
	}
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dateKey := d.Format("2006-01-02")
		if !existingDays[dateKey] {
			habitDays = append(habitDays, Habit{Date: d, Completed: false})
		}
	}
	sort.Slice(habitDays, func(i, j int) bool {
		return habitDays[i].Date.Before(habitDays[j].Date)
	})
	return habitDays
}
