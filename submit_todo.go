package main

import (
	"log"
	"time"

	"github.com/tj/go-naturaldate"
)

func SubmitTodo(m model) ViewId {
	task := m.addInputs[0].Value()
	dueDate := m.addInputs[1].Value()
	if task == "" {
		return SelectViewId
	} else {
		if dueDate == "" {
			dd := NewOptionalTime(time.Time{})
			dd.Set = false
			*m.todos = append(*m.todos, Todo{Task: task, Status: NotStarted, DueDate: dd})
		} else {
			parsedDate, err := naturaldate.Parse(dueDate, time.Now(), naturaldate.WithDirection(naturaldate.Future))
			if err != nil {
				log.Printf("Could not parse date from due date input: %v", err)
			}
			*m.todos = append(*m.todos, Todo{Task: task, Status: NotStarted, DueDate: NewOptionalTime(parsedDate)})
		}
		err := WriteAppData(m.dataFilePath, m.appData)
		if err != nil {
			log.Fatalf("Could not write new todo: %v", err)
		}
		return SelectViewId
	}
}
