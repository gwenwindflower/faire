package main

func deleteTodo(inputTodos []Todo, inputChoices []string, cursor int) (todos []Todo, choices []string) {
	todos = append(inputTodos[:cursor], inputTodos[cursor+1:]...)
	choices = append(inputChoices[:cursor], inputChoices[cursor+1:]...)
	return todos, choices
}
