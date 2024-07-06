package main

func deleteTodo(inputTodos []Todo, inputGraveyard []Todo, inputCursor int) (todos []Todo, graveyard []Todo, cursor int) {
	graveyard = append(inputGraveyard, inputTodos[inputCursor])
	todos = append(inputTodos[:inputCursor], inputTodos[inputCursor+1:]...)
	if inputCursor > len(todos)-1 {
		cursor = len(todos) - 1
	}
	return todos, graveyard, cursor
}
