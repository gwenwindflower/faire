package main

import "time"

func (m model) MoveActiveHabitDay(direction int) model {
	// TODO: Create m.ActiveHabitDateRange, a slice of dates
	// that way we can creat empty days for days that don't have
	// a Habit record yet. For example, let's say you don't use the
	// app for a week, you'd want to see the days that you missed
	// in the UI so you can catch up on your habits.
	// If activeHabitDay is zero for some reason, we initialize it to today
	if m.activeHabitDay.IsZero() {
		m.activeHabitDay = m.today
		m.aWeekBeforeActiveHabitDay = m.activeHabitDay.AddDate(0, 0, -7)
		return m
	}

	// If direction is 0 for some reason, reset to today
	if direction == 0 {
		m.activeHabitDay = m.today
		m.aWeekBeforeActiveHabitDay = m.activeHabitDay.AddDate(0, 0, -7)
		return m
	}

	// Prevent moving forward past today
	if m.activeHabitDay.Equal(m.today) && direction > 0 {
		return m
	}

	// Calculate the new active day
	newActiveDay := m.activeHabitDay.AddDate(0, 0, direction)
	// Calculate the new cutoff
	newWeekBefore := newActiveDay.AddDate(0, 0, -7)

	// Guard against going before the freaking beginning of time
	// (aka January 1, 1970 hehe)
	if newWeekBefore.Before(time.Unix(0, 0)) {
		return m
	}

	// Update the model
	m.activeHabitDay = newActiveDay
	m.aWeekBeforeActiveHabitDay = newWeekBefore

	for _, name := range m.habitList {
		(*m.habits)[name] = AddEmptyHabitDays((*m.habits)[name], m.aWeekBeforeActiveHabitDay, m.activeHabitDay)
	}

	return m
}
