package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ViewId int

const (
	SelectViewId ViewId = iota
	AddViewId
)

type Status int

const (
	NotStarted Status = iota
	InProgress
	Done
)

type Todo struct {
	Task   string `yaml:"task"`
	Status Status `yaml:"status"`
}

type model struct {
	todoPath         string
	choices          []string
	todos            []Todo
	newTaskTextInput textinput.Model
	cursor           int
	activeView       ViewId
}

func initialModel() model {
	todoPath, err := InitTodoFile()
	if err != nil || todoPath == "" {
		log.Fatalf("Could not initialize or find existing todo file: %v", err)
	}
	todos, err := FetchTodos(todoPath)
	if err != nil {
		log.Fatalf("Could not fetch todos: %v", err)
	}
	choices := []string{}
	for _, todo := range todos {
		choices = append(choices, todo.Task)
	}
	ti := textinput.New()
	ti.Placeholder = "Enter a new task"
	ti.Focus()
	return model{
		choices:          choices,
		todos:            todos,
		todoPath:         todoPath,
		activeView:       SelectViewId,
		newTaskTextInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func deleteTodo(inputTodos []Todo, inputChoices []string, cursor int) (todos []Todo, choices []string) {
	todos = append(inputTodos[:cursor], inputTodos[cursor+1:]...)
	choices = append(inputChoices[:cursor], inputChoices[cursor+1:]...)
	return todos, choices
}

func (m model) selectUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.todos[m.cursor].Status == NotStarted {
				m.todos[m.cursor].Status = Done
			} else {
				m.todos[m.cursor].Status = NotStarted
			}
		case "d":
			m.todos, m.choices = deleteTodo(m.todos, m.choices, m.cursor)
		case "a":
			m.activeView = AddViewId
		}
	}
	return m, nil
}

func (m model) addUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.activeView {
		case SelectViewId:
			return m.selectUpdate(msg)
		case AddViewId:
			return m.addUpdate(msg)
		}
	}
	return m, nil
}

func (m model) selectView() string {
	s := "What do you want to accomplish?\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = "➔"
		}
		checked := " "
		if m.todos[i].Status == Done {
			checked = "✔"
		}
		s += fmt.Sprintf("%s %s %s\n", cursor, checked, choice)
	}
	s += "\nPress 'a' to add a new todo, 'q' to quit."
	return s
}

func (m model) addView() string {
	return fmt.Sprintf(
		"Enter a new todo:\n\n%s\n\nPress Enter to add, or Esc to cancel.", m.newTaskTextInput.View(),
	)
}

func (m model) View() string {
	switch m.activeView {
	case SelectViewId:
		return m.selectView()
	case AddViewId:
		return m.addView()
	default:
		return "Unknown view"
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Bummer! Error: %v", err)
		os.Exit(1)
	}
}
