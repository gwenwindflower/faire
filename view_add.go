package main

import "fmt"

func (m model) ViewAdd() string {
	return fmt.Sprintf(
		"Enter a new todo:\n\n%s\n\nPress Enter to add, or Esc to cancel.", m.newTaskTextInput.View(),
	)
}
