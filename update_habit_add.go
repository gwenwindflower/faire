package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) UpdateHabitAdd(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			newHabit := m.addHabitInput.Value()
			if newHabit == "" {
				m.activeView = HabitSelectViewId
			}
			(*m.habits)[newHabit] = []Habit{}
			m.habitList = append(m.habitList, newHabit)
			m.addHabitInput.SetValue("")
			m.activeView = HabitSelectViewId
			return m, nil
		case "esc":
			m.addHabitInput.SetValue("")
			m.activeView = HabitSelectViewId
		}
	}
	var cmd tea.Cmd
	m.addHabitInput, cmd = m.addHabitInput.Update(msg)
	return m, cmd
}
