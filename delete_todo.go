package main

func deleteTodo(inputTodos []Todo, inputGraveyard []Todo, cursor int) (todos []Todo, graveyard []Todo) {
	graveyard = append(inputGraveyard, inputTodos[cursor])
	todos = append(inputTodos[:cursor], inputTodos[cursor+1:]...)
	return todos, graveyard
}
