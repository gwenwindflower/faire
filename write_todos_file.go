package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func WriteTodos(todoPath string, todos []Todo) error {
	todoFile, err := os.Create(todoPath)
	if err != nil {
		return fmt.Errorf("could not truncate todos file for rewrite: %v", err)
	}
	defer todoFile.Close()
	data, err := yaml.Marshal(todos)
	if err != nil {
		return fmt.Errorf("could not marshal todos to YAML: %v", err)
	}
	_, err = todoFile.Write(data)
	if err != nil {
		return fmt.Errorf("could not write new data to todos file: %v", err)
	}
	return nil
}
