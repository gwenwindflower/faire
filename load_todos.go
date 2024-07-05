package main

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadTodos(todoPath string) ([]Todo, error) {
	todos := []Todo{}
	todoFile, err := os.Open(todoPath)
	if err != nil {
		log.Fatalf("Could not open todos file: %v", err)
	}
	defer todoFile.Close()
	data, err := io.ReadAll(todoFile)
	if err != nil {
		log.Fatalf("Could not read todos file: %v", err)
	}
	err = yaml.Unmarshal(data, &todos)
	if err != nil {
		log.Fatalf("Could not unmarshal todos file: %v", err)
	}
	return todos, nil
}
