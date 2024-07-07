package main

func ToggleHabitToday(m model) model {
	// TODO: this should toggle the active day
	// not just today.

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
	var todaysEntry *Habit
	// Then e iterate over the selected habit (indicated by cursor)
	// to find the entry for today (if it exists!)
	// If it does we set todaysEntry to point to it
	for i, h := range selectedHabit {
		if h.Date.Equal(m.today) {
			todaysEntry = &selectedHabit[i]
			break
		}
	}
	if todaysEntry == nil {
		// If todaysEntry is nil, we create a new entry
		(*m.habits)[selectedHabitName] = append(selectedHabit, Habit{Date: m.today, Completed: true})
	} else {
		// Otherwise if we did find an entry, we just toggle the completion
		todaysEntry.Completed = !todaysEntry.Completed
	}
	return m
}
