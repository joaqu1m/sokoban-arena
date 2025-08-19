package loaders

import (
	"log"
	"strings"
)

type ParsedLevel struct {
	Title    string     `json:"title"`
	Author   string     `json:"author"`
	Group    string     `json:"group"`
	Map      [][]string `json:"map"`
	Textures []string   `json:"textures"`
}

var ParsedLevelMap = map[string]ParsedLevel{}

func GetParsedLevel(name string) ParsedLevel {
	if !strings.Contains(name, ":") {
		name = "default:" + name
	}
	item, ok := ParsedLevelMap[name]
	if !ok {
		log.Fatal("Level not found: " + name)
	}
	return item
}

func init() {
	loadFiles("assets/levels/*.json", &ParsedLevelMap)
	log.Println("Levels loaded successfully")
}
