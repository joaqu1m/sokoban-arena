package loaders

import (
	"log"
	"rl-go/utils"
	"strings"
)

type ParsedBlock struct {
	Texture     utils.TextureReference `json:"texture"`
	DefaultSize utils.Size             `json:"default_size"`
	Collision   utils.Collision        `json:"collision"`
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
	loadFiles("assets/blocks/*.json", &ParsedBlockMap)
	log.Println("Blocks loaded successfully")
}
