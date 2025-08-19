package loaders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func loadFile[T any](filePath string, out *map[string]T) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error reading textures.json: %v\n", err))
	}

	var serialized map[string]T
	err = json.Unmarshal(jsonData, &serialized)
	if err != nil {
		fmt.Printf("Error parsing textures.json: %v\n", err)
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
		panic(fmt.Sprintf("Error reading level directory: %v\n", err))
	}

	for _, file := range files {
		jsonData, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", file, err)
			continue
		}

		var serialized map[string]T
		err = json.Unmarshal(jsonData, &serialized)
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
