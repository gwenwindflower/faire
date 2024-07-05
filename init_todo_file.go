package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitTodoFile() (todoPath string, err error) {
	todoPath = ""
	home, err := os.UserHomeDir()
	if err != nil {
		return todoPath, fmt.Errorf("could not get user home directory: %v", err)
	}
	todoPath = filepath.Join(home, ".config", "faire", "todos.yaml")
	_, err = os.Stat(todoPath)
	if os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(todoPath), 0755)
		_, err = os.Create(todoPath)
		if err != nil {
			return todoPath, fmt.Errorf("could not create todos file: %v", err)
		}
	} else if err != nil {
		return todoPath, fmt.Errorf("error checking for todos file: %v", err)
	}
	return todoPath, nil
}
