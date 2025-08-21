package loaders

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func loadFile[T any](filePath string, out *map[string]T) {
	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error reading %s: %v\n", filePath, err))
	}

	var serialized map[string]T
	err = yaml.Unmarshal(yamlData, &serialized)
	if err != nil {
		fmt.Printf("Error parsing %s: %v\n", filePath, err)
		return
	}

	filename := filepath.Base(filePath)
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))

	result := map[string]T{}
	for key, data := range serialized {
		newKey := filename + ":" + key
		result[newKey] = data
	}

	*out = result
}

func loadFiles[T any](filePath string, out *map[string]T) {
	files, err := filepath.Glob(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error reading level directory %s: %v\n", filePath, err))
	}

	for _, file := range files {
		yamlData, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", file, err)
			continue
		}

		var serialized map[string]T
		err = yaml.Unmarshal(yamlData, &serialized)
		if err != nil {
			fmt.Printf("Error parsing %s: %v\n", file, err)
			continue
		}

		filename := filepath.Base(file)
		filename = strings.TrimSuffix(filename, filepath.Ext(filename))

		for key, data := range serialized {
			newKey := filename + ":" + key
			if _, exists := (*out)[newKey]; exists {
				fmt.Printf("Duplicate level ID found: %s\n", newKey)
				continue
			}
			(*out)[newKey] = data
		}
	}
}
