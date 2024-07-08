package main

import "sort"

func ToggleActiveHabit(m model) model {
	// m.habitList is a slice of strings containing
	// the names of the habits, which are map keys
	// so cannot otherwise be iterated over.
	// NB: we use separate cursors for habits and todos
	// so that when you switch between views, the cursor
	// state is preserved.
	selectedHabitName := m.habitList[m.habitCursor]
	// This is the actual habit data
	selectedHabit := (*m.habits)[selectedHabitName]

	// We create a pointer to a Habit struct
	var activeEntry *Habit
	// Then e iterate over the selected habit (indicated by cursor)
	// to find the entry for today (if it exists!)
	// If it does we set todaysEntry to point to it
	for i, h := range selectedHabit {
		if h.Date.Equal(m.activeHabitDay) {
			activeEntry = &selectedHabit[i]
			break
		}
	}
	if activeEntry == nil {
		// If activeEntry is nil, we create a new entry
		selectedHabit = append(selectedHabit, Habit{Date: m.activeHabitDay, Completed: true})
		sort.Slice(selectedHabit, func(i, j int) bool {
			return selectedHabit[i].Date.Before(selectedHabit[j].Date)
		})
	} else {
		// Otherwise if we did find an entry, we just toggle the completion
		activeEntry.Completed = !activeEntry.Completed
	}
	return m
}
