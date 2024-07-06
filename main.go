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
	DueDate OptionalTime `yaml:"due_date"`
	Task    string       `yaml:"task"`
	Status  Status       `yaml:"status"`
}

type model struct {
	todoPath            string
	todos               []Todo
	graveyard           []Todo
	addInputs           []textinput.Model
	cursor              int
	activeView          ViewId
	addInputsFocusIndex int
	hideCompleted       bool
}

func initialModel() model {
	// We store Todos in YAML in ~/.config/faire/todos.yaml
	// I chose to use a YAML file because it's human readable
	// so if something gets borked you can fix it manually
	// It's also easy to marshal and unmarshal in Go and convert
	// to other formats like JSON if we eventually build a web
	// interface or something
	todoPath, err := InitTodosFile()
	if err != nil || todoPath == "" {
		log.Fatalf("Could not initialize or find existing todo file: %v", err)
	}
	todos, err := LoadTodos(todoPath)
	if err != nil {
		// If we can't load the todos file, we should exit
		// because we can't do anything without it
		log.Fatalf("Could not fetch todos: %v", err)
	}
	// Text Input is a Bubble, a reusable component built in Bubble Tea
	// for use in Bubble Tea programs. Saves us writing our own.
	nt := textinput.New()
	nt.Placeholder = "Enter a new task"
	dd := textinput.New()
	dd.Placeholder = "Enter a due date in natural language"
	return model{
		todos:               todos,
		todoPath:            todoPath,
		graveyard:           []Todo{},
		activeView:          SelectViewId,
		addInputs:           []textinput.Model{nt, dd},
		addInputsFocusIndex: 0,
		hideCompleted:       false,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.activeView {
		case SelectViewId:
			return m.UpdateSelect(msg)
		case AddViewId:
			return m.UpdateAdd(msg)
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.activeView {
	case SelectViewId:
		return m.ViewSelect()
	case AddViewId:
		return m.ViewAdd()
	default:
		return fmt.Sprintf("Unknown view: %v", m.activeView)
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Bummer! Error: %v", err)
		os.Exit(1)
	}
}
