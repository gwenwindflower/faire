package main

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadTodos(todoPath string) ([]Todo, error) {
	t := []Todo{}
	tf, err := os.Open(todoPath)
	if err != nil {
		log.Fatalf("Could not open todos file: %v", err)
	}
	defer tf.Close()
	data, err := io.ReadAll(tf)
	if err != nil {
		log.Fatalf("Could not read todos file: %v", err)
	}
	err = yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalf("Could not unmarshal todos file: %v", err)
	}
	return t, nil
}
