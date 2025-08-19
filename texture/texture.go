package texture

import (
	"rl-go/loaders"
	"rl-go/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Texture struct {
	Path          string       `json:"path"`
	Size          utils.Size   `json:"size"`
	ItemSize      utils.Size   `json:"item_size"`
	LoadedTexture rl.Texture2D `json:"loaded_texture"`
}

func NewTexture(texture loaders.ParsedTexture) Texture {

	tex := rl.LoadTexture(texture.Path)
	if tex.ID == 0 {
		panic("Failed to load texture: " + texture.Path)
	}

	size := utils.Size{
		W: int32(tex.Width),
		H: int32(tex.Height),
	}

	return Texture{
		texture.Path,
		size,
		texture.ItemSize,
		tex,
	}
}

func (t *Texture) GetFromIndex(index int) rl.Rectangle {

	itemsPerRow := t.Size.W / t.ItemSize.W
	itemsPerColumn := t.Size.H / t.ItemSize.H
	maxItems := itemsPerRow * itemsPerColumn

	if index < 0 {
		index = 0
	}
	if index >= int(maxItems) {
		index = int(maxItems) - 1
	}

	row := int32(index) / itemsPerRow
	col := int32(index) % itemsPerRow

	x := col * t.ItemSize.W
	y := row * t.ItemSize.H

	return rl.NewRectangle(float32(x), float32(y), float32(t.ItemSize.W), float32(t.ItemSize.H))
}
