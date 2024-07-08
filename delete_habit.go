package main

func (m *model) DeleteHabit() {
	selectedHabitName := m.habitList[m.habitCursor]
	delete(*m.habits, selectedHabitName)
	m.habitList = m.habitList[:m.habitCursor+copy(m.habitList[m.habitCursor:], m.habitList[m.habitCursor+1:])]
	if m.habitCursor >= len(m.habitList) {
		m.habitCursor = len(m.habitList) - 1
	}
	if m.habitCursor < 0 {
		m.habitCursor = 0
	}
}
