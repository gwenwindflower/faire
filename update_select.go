package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) UpdateSelect(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			WriteTodos(m.todoPath, m.todos)
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.todos) - 1
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
		case "1":
			m.todos[m.cursor].Status = NotStarted
		case "2":
			m.todos[m.cursor].Status = InProgress
		case "3":
			m.todos[m.cursor].Status = Done
		case "d":
			m.todos, m.graveyard = deleteTodo(m.todos, m.graveyard, m.cursor)
			if m.cursor > len(m.todos)-1 {
				m.cursor = len(m.todos) - 1
			}
		case "u":
			m.todos, m.graveyard = undoDeleteTodo(m.todos, m.graveyard)
		case "a":
			m.activeView = AddViewId
			cmd = m.addInputs[m.addInputsFocusIndex].Focus()
		case "h":
			m.hideCompleted = !m.hideCompleted
		}
	}
	return m, cmd
}
