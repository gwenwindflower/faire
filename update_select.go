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
			WriteAppData(m.dataFilePath, m.appData)
			return m, tea.Quit
		case "j", "down", "tab", "ctrl+n":
			if m.todoCursor < len(*m.todos)-1 {
				m.todoCursor++
			} else {
				m.todoCursor = 0
			}
		case "k", "up", "shift+tab", "ctrl+p":
			if m.todoCursor > 0 {
				m.todoCursor--
			} else {
				m.todoCursor = len(*m.todos) - 1
			}
		case "enter", " ":
			switch (*m.todos)[m.todoCursor].Status {
			case NotStarted:
				(*m.todos)[m.todoCursor].Status = InProgress
			case InProgress:
				(*m.todos)[m.todoCursor].Status = Done
			case Done:
				(*m.todos)[m.todoCursor].Status = NotStarted
			}
		case "1":
			(*m.todos)[m.todoCursor].Status = NotStarted
		case "2":
			(*m.todos)[m.todoCursor].Status = InProgress
		case "3":
			(*m.todos)[m.todoCursor].Status = Done
		case "d":
			m.graveyard, m.todoCursor = deleteTodo(m.todos, m.graveyard, m.todoCursor)
		case "D":
			m.graveyard, m.todoCursor = deleteCompletedTodos(m.todos, m.graveyard, m.todoCursor)
		case "u":
			*m.todos, m.graveyard = undoDeleteTodo(*m.todos, m.graveyard)
		case "a":
			m.addInputsFocusIndex = 0
			for i := range m.addInputs {
				m.addInputs[i].SetValue("")
				if i == m.addInputsFocusIndex {
					cmd = m.addInputs[i].Focus()
				} else {
					m.addInputs[i].Blur()
				}
			}
			m.activeView = AddViewId
		case "h":
			m.hideCompleted = !m.hideCompleted
		case "r":
			m.activeView = HabitSelectViewId
		case "?":
			m.previousViewFromHelp = SelectViewId
			m.activeView = HelpViewId
		}
	}
	return m, cmd
}
