package main

func deleteCompletedTodos(
	inputTodos []Todo, inputGraveyard []Todo, inputCursor int,
) (todos []Todo, graveyard []Todo, cursor int) {
	todos = []Todo{}
	graveyard = inputGraveyard
	cursor = inputCursor
	deletedCount := 0
	for _, t := range inputTodos {
		if t.Status == Done {
			graveyard = append(graveyard, t)
			deletedCount++
		} else {
			todos = append(todos, t)
		}
	}
	cursor -= deletedCount
	if cursor < 0 {
		cursor = 0
	}
	return todos, graveyard, cursor
}
