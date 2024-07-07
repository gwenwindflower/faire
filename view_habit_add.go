package main

func (m model) ViewHabitAdd() string {
	s := headerStyle.Render("Enter a new habit to track:")
	s += "\n"
	s += m.addHabitInput.View() + "\n"
	footer := footerStyle.Render("Press '?' for shortcuts.")
	s += footer
	return s
}
