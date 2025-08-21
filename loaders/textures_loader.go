package loaders

import (
	"log"
	"rl-go/utils"
	"strings"
)

type ParsedTexture struct {
	Path     string     `yaml:"path"`
	ItemSize utils.Size `yaml:"item_size"`
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
	loadFiles("assets/textures/*.yaml", &ParsedTextureMap)
	log.Println("Textures loaded successfully")
}
