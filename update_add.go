package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) UpdateAdd(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.activeView = SelectViewId
		case "enter":
			task := m.newTaskTextInput.Value()
			if task == "" {
				m.activeView = SelectViewId
			} else {
				m.todos = append(m.todos, Todo{Task: task, Status: NotStarted})
				m.choices = append(m.choices, task)
				err := WriteTodos(m.todoPath, m.todos)
				if err != nil {
					log.Fatalf("Could not write todos: %v", err)
				}
				m.activeView = SelectViewId
			}
		default:
			m.newTaskTextInput, cmd = m.newTaskTextInput.Update(msg)
		}
	}
	return m, cmd
}
