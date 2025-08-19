package block

import (
	"log"
	"rl-go/texture"
	"rl-go/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Coords           rl.Vector2             `json:"coords"`
	Size             utils.Size             `json:"size"`
	Collision        utils.Collision        `json:"collision"`
	TextureReference utils.TextureReference `json:"texture_reference"`
}

func NewBlock(x, y float32, w, h int32, textureReference utils.TextureReference) Block {
	return Block{
		Coords:           rl.Vector2{X: x, Y: y},
		Size:             utils.Size{W: w, H: h},
		TextureReference: textureReference,
	}
}

func (b *Block) Plot(textureManager texture.TextureManager) {

	destRect := rl.Rectangle{
		X:      b.Coords.X,
		Y:      b.Coords.Y,
		Width:  float32(b.Size.W),
		Height: float32(b.Size.H),
	}

	texture := textureManager.GetTexture(b.TextureReference.ID)
	srcRect := texture.GetFromIndex(b.TextureReference.Index)

	rl.DrawTexturePro(
		texture.LoadedTexture,
		srcRect,
		destRect,
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}

func (b *Block) Println() {
	log.Printf("Block at (%.2f, %.2f) with size (%d, %d), texture: %s, index: %d\n",
		b.Coords.X, b.Coords.Y, b.Size.W, b.Size.H,
		b.TextureReference.ID, b.TextureReference.Index)
}
