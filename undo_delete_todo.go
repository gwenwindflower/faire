package main

func undoDeleteTodo(it []Todo, ig []Todo) (t []Todo, g []Todo) {
	if len(ig) == 0 {
		return it, ig
	}
	// Pop the last item off the graveyard
	undone := ig[len(ig)-1]
	g = ig[:len(ig)-1]
	// Add the undone todo back to the todos
	t = append(it, undone)
	return t, g
}
