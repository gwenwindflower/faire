package main

import (
	"fmt"
)

func (m model) ViewHelp() string {
	header := headerStyle.Render("Shortcuts")
	return fmt.Sprintf(`%s
j/k, up/down: Move cursor
space: Cycle todo status
1/2/3: Set todo status
a: Add todo
d: Delete todo
D: Delete all completed todos
h: Show/hide completed todos
u: Undo delete
`, header)
}
