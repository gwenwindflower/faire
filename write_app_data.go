package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func WriteAppData(filepath string, appData *AppData) error {
	tf, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("could not clear file for rewrite: %v", err)
	}
	defer tf.Close()
	data, err := yaml.Marshal(appData)
	if err != nil {
		return fmt.Errorf("could not marshal data to YAML: %v", err)
	}
	_, err = tf.Write(data)
	if err != nil {
		return fmt.Errorf("could not write new data to file: %v", err)
	}
	return nil
}
