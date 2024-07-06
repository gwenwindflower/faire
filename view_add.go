package main

func (m model) ViewAdd() string {
	s := headerStyle.Render("Enter a new todo:")
	s += "\n"
	for i := range m.addInputs {
		s += m.addInputs[i].View() + "\n"
	}
	s += "\nPress '?' for shortcuts."
	return s
}
