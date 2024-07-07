package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var shortcutBlockStyle = lipgloss.NewStyle().
	Align(lipgloss.Left).
	Border(lipgloss.RoundedBorder())

var helpHeaderStyle = headerStyle.
	Width(25).
	Border(lipgloss.RoundedBorder()).
	Align(lipgloss.Center)

var subHeaderStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Foreground(lipgloss.Color("#a6d189"))

func (m model) ViewHelp() string {
	header := helpHeaderStyle.Render("Shortcuts")
	selectSubheader := subHeaderStyle.Render("Todo List")
	selectShortcuts := shortcutBlockStyle.Render(`space: Cycle todo status
1/2/3: Set todo status
a: Add todo
d: Delete todo
D: Delete all completed
h: Toggle show completed
u: Undo delete
j/k: Move cursor
tab/shift+tab: Move cursor
down/up: Move cursor
ctrl+n/ctrl+p: Move cursor
q/ctrl+c: Quit`)
	addSubheader := subHeaderStyle.Render("Add todo")
	addShortcuts := shortcutBlockStyle.Render(`esc: Return to list
enter: Save todo
tab/shift+tab: Move input focus
up/down: Move input focus
ctrl+c: Quit`)
	habitHeader := subHeaderStyle.Render("Habit Tracker")
	habitShortcuts := shortcutBlockStyle.Render(`space: Cycle habit status
a: Add habit
j/k: Move cursor
down/up: Move cursor
ctrl+n/ctrl+p: Move cursor
tab/shift+tab: Move cursor
q/ctrl+c: Quit`)
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s", header, selectSubheader, selectShortcuts, addSubheader, addShortcuts, habitHeader, habitShortcuts)
}
