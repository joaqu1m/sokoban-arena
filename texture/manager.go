package texture

import (
	"fmt"
	"log"
	"rl-go/loaders"
	"rl-go/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureManager struct {
	Textures map[string]Texture
}

func NewTextureManager(loadAllTextures bool) TextureManager {

	textures := map[string]Texture{}

	textureManager := TextureManager{textures}

	if loadAllTextures {
		textureManager.LoadAllTextures()
	}

	return textureManager
}

func (tr *TextureManager) GetTexture(name string) Texture {
	name = utils.NormalizeName(name)

	texture, ok := tr.Textures[name]
	if !ok {
		log.Fatal("Texture not found: " + name)
	}

	return texture
}

func (tr *TextureManager) GetTextureFromIndex(name string, index int) rl.Rectangle {
	name = utils.NormalizeName(name)

	texture, ok := tr.Textures[name]
	if !ok {
		log.Fatal("Texture not found: " + name)
	}

	return texture.GetFromIndex(index)
}

func (tr *TextureManager) LoadTexture(name string) {
	name = utils.NormalizeName(name)

	if _, ok := tr.Textures[name]; ok {
		return // Texture already loaded
	}

	parsedTexture := loaders.GetParsedTexture(name)

	tr.Textures[name] = NewTexture(parsedTexture)
}

func (tr *TextureManager) UnloadTexture(name string) {
	name = utils.NormalizeName(name)

	texture, ok := tr.Textures[name]
	if !ok {
		log.Fatal("Texture not found: " + name)
	}

	rl.UnloadTexture(texture.LoadedTexture)
	delete(tr.Textures, name)
}

func (tr *TextureManager) LoadAllTextures() {
	for name, texture := range loaders.ParsedTextureMap {
		tex := NewTexture(texture)
		fmt.Printf("Loaded texture '%s': %dx%d pixels (item size: %dx%d)\n",
			name, tex.Size.W, tex.Size.H, tex.ItemSize.W, tex.ItemSize.H)
		tr.Textures[name] = tex
	}
}

func (tr *TextureManager) UnloadAllTextures() {
	for name, texture := range tr.Textures {
		rl.UnloadTexture(texture.LoadedTexture)
		delete(tr.Textures, name)
	}
}
