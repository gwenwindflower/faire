package main

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadData(filepath string) (AppData, error) {
	d := AppData{}
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Could not open todos file: %v", err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Could not read todos file: %v", err)
	}
	err = yaml.Unmarshal(data, &d)
	if err != nil {
		log.Fatalf("Could not unmarshal todos file: %v", err)
	}
	return d, nil
}
