package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) UpdateHabitSelect(msg tea.Msg) (tea.Model, tea.Cmd) {
	today := time.Now().Truncate(24 * time.Hour) // Get today's date without time
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down", "ctrl+n", "tab":
			if m.habitCursor < len(m.habitList)-1 {
				m.habitCursor++
			} else {
				m.habitCursor = 0
			}
		case "k", "up", "shift+tab", "ctrl+p":
			if m.habitCursor > 0 {
				m.habitCursor--
			} else {
				m.habitCursor = len(m.habitList) - 1
			}

		case " ":
			selectedHabitName := m.habitList[m.habitCursor]
			selectedHabit := m.habits[selectedHabitName]

			var todaysEntry *Habit
			for i, h := range selectedHabit {
				if h.Date.Equal(today) {
					todaysEntry = &selectedHabit[i]
					break
				}
			}
			if todaysEntry == nil {
				m.habits[selectedHabitName] = append(selectedHabit, Habit{Date: today, Completed: true})
			} else {
				todaysEntry.Completed = !todaysEntry.Completed
			}
		case "t":
			m.activeView = SelectViewId
		case "a":
			m.addHabitInput.Focus()
			m.activeView = HabitAddViewId
		case "q", "ctrl+c":
			WriteAppData(m.dataFilePath, m.appData)
			return m, tea.Quit
		case "?":
			m.previousViewFromHelp = HabitSelectViewId
			m.activeView = HelpViewId
		}
	}
	return m, nil
}
