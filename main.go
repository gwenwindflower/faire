package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ViewId int

// This is a fun thing in Go:
// iota is a built-in constant generator that increments by 1 each time it's used
// It's used here to give each view a unique ID
// This is useful for switching between
// views in the update function
const (
	SelectViewId ViewId = iota
	AddViewId
	HelpViewId
	HabitSelectViewId
	HabitAddViewId
)

type Status int

// We do the same thing with Status
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

type Habit struct {
	Date      time.Time `yaml:"date"`
	Completed bool      `yaml:"completed"`
}

type AppData struct {
	Habits map[string][]Habit `yaml:"habits"`
	Todos  []Todo             `yaml:"todos"`
}

type model struct {
	habits                    map[string][]Habit
	dataFilePath              string
	appData                   AppData
	habitList                 []string
	todos                     []Todo
	graveyard                 []Todo
	addInputs                 []textinput.Model
	addHabitInput             textinput.Model
	todoCursor                int
	habitCursor               int
	activeView                ViewId
	addInputsFocusIndex       int
	previousViewFromHelp      ViewId
	hideCompleted             bool
	today                     time.Time
	activeHabitDay            time.Time
	aWeekBeforeActiveHabitDay time.Time
}

func initialModel() model {
	// We store Todos in YAML in ~/.config/faire/todos.yaml
	// I chose to use a YAML file because it's human readable
	// so if something gets borked you can fix it manually
	// It's also easy to marshal and unmarshal in Go and convert
	// to other formats like JSON if we eventually build a web
	// interface or something
	p, err := InitDataFile()
	if err != nil || p == "" {
		log.Fatalf("Could not initialize or find existing todo file: %v", err)
	}
	d, err := LoadData(p)
	if err != nil {
		// If we can't load the todos file, we should exit
		// because we can't do anything without it
		log.Fatalf("Could not load data file: %v", err)
	}
	// Text Input is a Bubble, a reusable component built in Bubble Tea
	// for use in Bubble Tea programs. Saves us writing our own.
	nt := textinput.New()
	nt.Placeholder = "Enter a new task"
	dd := textinput.New()
	dd.Placeholder = "Enter a due date in natural language"
	ah := textinput.New()
	ah.Placeholder = "Enter a new habit to track"
	// Habits are stored in a map of strings to slices of individual
	// dates and whether they were completed or not. This is so we can
	// easily add new habits and check them off for each day.
	// But it does mean we need to convert the map to a slice of
	// strings for the habit select view.
	// We also need to store the habit list in the model so we can
	// display it in the view.
	habitList := []string{}
	for k := range d.Habits {
		habitList = append(habitList, k)
	}
	// We use a window of today's date through a
	// week before in several places,
	// particularly habit tracking, so we just
	// put it in the model so we do it once.
	today, aWeekBeforeActiveHabitDay := GetTimeWindow()
	return model{
		dataFilePath:              p,
		appData:                   d,
		todos:                     d.Todos,
		habits:                    d.Habits,
		habitList:                 habitList,
		graveyard:                 []Todo{},
		activeView:                SelectViewId,
		addInputs:                 []textinput.Model{nt, dd},
		addHabitInput:             ah,
		addInputsFocusIndex:       0,
		hideCompleted:             false,
		previousViewFromHelp:      SelectViewId,
		today:                     today,
		activeHabitDay:            today,
		aWeekBeforeActiveHabitDay: aWeekBeforeActiveHabitDay,
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
		case HabitSelectViewId:
			return m.UpdateHabitSelect(msg)
		case HabitAddViewId:
			return m.UpdateHabitAdd(msg)
		case HelpViewId:
			return m.UpdateHelp(msg)
		default:
			return m.UpdateSelect(msg)
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
	case HabitSelectViewId:
		return m.ViewHabitSelect()
	case HabitAddViewId:
		return m.ViewHabitAdd()
	case HelpViewId:
		return m.ViewHelp()
	default:
		return m.ViewSelect()
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Bummer! Error: %v", err)
		os.Exit(1)
	}
}
