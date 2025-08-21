package loaders

import (
	"log"
	"rl-go/utils"
	"strings"
)

type ParsedBlock struct {
	Texture     utils.TextureReference `yaml:"texture"`
	DefaultSize utils.Size             `yaml:"default_size"`
	Collision   utils.Collision        `yaml:"collision"`
}

var ParsedBlockMap = map[string]ParsedBlock{}

func GetParsedBlock(name string) ParsedBlock {
	if !strings.Contains(name, ":") {
		name = "default:" + name
	}
	item, ok := ParsedBlockMap[name]
	if !ok {
		log.Fatal("Block not found: " + name)
	}
	return item
}

func init() {
	loadFiles("assets/blocks/*.yaml", &ParsedBlockMap)
	log.Println("Blocks loaded successfully")
}
