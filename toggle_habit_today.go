package main

func ToggleHabitToday(m model) model {
	selectedHabitName := m.habitList[m.habitCursor]
	selectedHabit := m.habits[selectedHabitName]

	var todaysEntry *Habit
	for i, h := range selectedHabit {
		if h.Date.Equal(m.today) {
			todaysEntry = &selectedHabit[i]
			break
		}
	}
	if todaysEntry == nil {
		m.habits[selectedHabitName] = append(selectedHabit, Habit{Date: m.today, Completed: true})
	} else {
		todaysEntry.Completed = !todaysEntry.Completed
	}
	return m
}
