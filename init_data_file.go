package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitDataFile() (tp string, err error) {
	tp = ""
	home, err := os.UserHomeDir()
	if err != nil {
		return tp, fmt.Errorf("could not get user home directory: %v", err)
	}
	tp = filepath.Join(home, ".config", "faire", "todos.yaml")
	_, err = os.Stat(tp)
	if os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(tp), 0755)
		_, err = os.Create(tp)
		if err != nil {
			return tp, fmt.Errorf("could not create data file: %v", err)
		}
	} else if err != nil {
		return tp, fmt.Errorf("error checking for data file: %v", err)
	}
	return tp, nil
}
