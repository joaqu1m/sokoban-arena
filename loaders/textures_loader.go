package loaders

import (
	"log"
	"rl-go/utils"
	"strings"
)

type ParsedTexture struct {
	Path     string     `json:"path"`
	ItemSize utils.Size `json:"item_size"`
}

var ParsedTextureMap = map[string]ParsedTexture{}

func GetParsedTexture(name string) ParsedTexture {
	if !strings.Contains(name, ":") {
		name = "default:" + name
	}
	item, ok := ParsedTextureMap[name]
	if !ok {
		log.Fatal("Texture not found: " + name)
	}
	return item
}

func init() {
	loadFiles("assets/textures/*.json", &ParsedTextureMap)
	log.Println("Textures loaded successfully")
}
