package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tj/go-naturaldate"
)

func (m model) UpdateAdd(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" {
				task := m.addInputs[0].Value()
				dueDate := m.addInputs[1].Value()
				if task == "" {
					m.activeView = SelectViewId
				} else {
					if dueDate == "" {
						dd := NewOptionalTime(time.Time{})
						dd.Set = false
						m.todos = append(m.todos, Todo{Task: task, Status: NotStarted, DueDate: dd})
					} else {
						parsedDate, err := naturaldate.Parse(dueDate, time.Now(), naturaldate.WithDirection(naturaldate.Future))
						if err != nil {
							log.Printf("Could not parse date from due date input: %v", err)
						}
						m.todos = append(m.todos, Todo{Task: task, Status: NotStarted, DueDate: NewOptionalTime(parsedDate)})
					}
					err := WriteAppData(m.dataFilePath, m.appData)
					if err != nil {
						log.Fatalf("Could not write todos: %v", err)
					}
					m.activeView = SelectViewId
				}
			}
			if s == "up" || s == "shift+tab" {
				m.addInputsFocusIndex--
			} else {
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
