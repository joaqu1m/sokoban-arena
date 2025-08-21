package loaders

import (
	"log"
	"strings"
)

type ParsedLevel struct {
	Title    string     `yaml:"title"`
	Author   string     `yaml:"author"`
	Group    string     `yaml:"group"`
	Map      [][]string `yaml:"map"`
	Textures []string   `yaml:"textures"`
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
	loadFiles("assets/levels/*.yaml", &ParsedLevelMap)
	log.Println("Levels loaded successfully")
}
