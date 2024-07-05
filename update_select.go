package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) UpdateSelect(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			WriteTodos(m.todoPath, m.todos)
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.choices) - 1
			}
		case "enter", " ":
			switch m.todos[m.cursor].Status {
			case NotStarted:
				m.todos[m.cursor].Status = InProgress
			case InProgress:
				m.todos[m.cursor].Status = Done
			case Done:
				m.todos[m.cursor].Status = NotStarted
			}
		case "d":
			m.todos, m.choices = deleteTodo(m.todos, m.choices, m.cursor)
			if m.cursor > len(m.choices)-1 {
				m.cursor = len(m.choices) - 1
			}
		case "a":
			m.activeView = AddViewId
		}
	}
	return m, nil
}
