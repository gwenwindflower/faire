package main

func deleteTodo(todos *[]Todo, inputGraveyard [][]Todo, inputCursor int) (graveyard [][]Todo, cursor int) {
	grave := []Todo{}
	grave = append(grave, (*todos)[inputCursor])
	graveyard = append(inputGraveyard, grave)
	*todos = append((*todos)[:inputCursor], (*todos)[inputCursor+1:]...)
	// Keep the cursor in bounds
	if inputCursor > len(*todos)-1 {
		cursor = len(*todos) - 1
	}
	if inputCursor < 0 {
		cursor = 0
	}
	return
}
