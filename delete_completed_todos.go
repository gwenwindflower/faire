package main

func deleteCompletedTodos(
	todos *[]Todo, inputGraveyard [][]Todo, inputCursor int,
) (graveyard [][]Todo, cursor int) {
	// Preserve the graveyard, we only want to add to it
	graveyard = inputGraveyard
	cursor = inputCursor
	NotDoneTodos := []Todo{}
	// We'll use this to make sure our cursor
	// doesn't go out of bounds after deletion
	deletedCount := 0
	grave := []Todo{}
	for _, t := range *todos {
		if t.Status == Done {
			grave = append(grave, t)
			deletedCount++
		} else {
			// NB: because this is a named return, we don't need to
			// initialize an empty `[]Todos{}` slice, we can just use it
			// which is a cool feature of Go
			NotDoneTodos = append(NotDoneTodos, t)
		}
	}
	// Update the todos slice within our AppData struct
	// to only include the NotDoneTodos
	*todos = NotDoneTodos
	// Account for the deleted todos
	// in the cursor position
	cursor -= deletedCount
	// Make sure the cursor doesn't go out of bounds
	// from the above operation
	// This first guard condition shouldn't be possible,
	// but just to be safe it's pretty cheap to check
	if cursor >= len(*todos)-1 {
		cursor = len(*todos) - 1
	}
	if cursor < 0 {
		cursor = 0
	}
	graveyard = append(graveyard, grave)
	// Because we used named return values, we don't need to return them here
	// We can just `return` and Go will know what to do
	return
}
