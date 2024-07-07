package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) UpdateAdd(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" {
				if m.addInputsFocusIndex < len(m.addInputs)-1 {
					m.addInputsFocusIndex++
				} else {
					m.activeView = SubmitTodo(m)
				}
			}
			if s == "up" || s == "shift+tab" {
				m.addInputsFocusIndex--
			}
			if s == "down" || s == "tab" {
				m.addInputsFocusIndex++
			}

			// If we're at the end of the inputs, loop back to the beginning
			if m.addInputsFocusIndex > len(m.addInputs)-1 {
				m.addInputsFocusIndex = 0
			} else if m.addInputsFocusIndex < 0 {
				m.addInputsFocusIndex = len(m.addInputs) - 1
			}
			cmds := make([]tea.Cmd, len(m.addInputs))
			for i := 0; i < len(m.addInputs); i++ {
				if i == m.addInputsFocusIndex {
					cmds[i] = m.addInputs[i].Focus()
					continue
				}
				m.addInputs[i].Blur()
			}
			return m, tea.Batch(cmds...)
		case "?":
			m.previousViewFromHelp = AddViewId
			m.activeView = HelpViewId
			return m, nil
		case "esc":
			m.activeView = SelectViewId
			return m, nil
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.addInputs))

	for i := range m.addInputs {
		m.addInputs[i], cmds[i] = m.addInputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
