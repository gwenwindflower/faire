package main

func undoDeleteTodo(inputTodos []Todo, inputGraveyard [][]Todo) (todos []Todo, graveyard [][]Todo) {
	if len(inputGraveyard) == 0 {
		return inputTodos, inputGraveyard
	}
	mostRecentGrave := inputGraveyard[len(inputGraveyard)-1]
	graveyard = inputGraveyard[:len(inputGraveyard)-1]

	// We can save a tiiiny bit of memory by pre-allocating the slice
	// since we know the exact length we need
	todos = make([]Todo, 0, len(inputTodos)+len(mostRecentGrave))
	todos = append(todos, inputTodos...)
	todos = append(todos, mostRecentGrave...)

	return
}
