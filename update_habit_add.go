package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) UpdateHabitAdd(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			habit := m.addHabitInput.Value()
			if habit == "" {
				m.activeView = HabitSelectViewId
			}
			m.habits[habit] = []Habit{}
			m.habitList = append(m.habitList, habit)
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
