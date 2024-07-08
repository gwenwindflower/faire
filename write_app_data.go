package main

import (
	"fmt"
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

func WriteAppData(filepath string, appData *AppData) error {
	// Sort habits by date
	for _, h := range appData.Habits {
		sort.Slice(h, func(i, j int) bool {
			return h[i].Date.Before(h[j].Date)
		})
	}
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
