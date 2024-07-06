package main

func (m model) ViewAdd() string {
	s := "Enter a new todo:\n"
	for i := range m.addInputs {
		s += m.addInputs[i].View() + "\n"
	}
	s += "\nPress 'tab' to switch inputs, 'esc' to go back, 'enter' to save."
	return s
}
