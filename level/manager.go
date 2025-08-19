package level

import (
	"log"
	"rl-go/loaders"
	"rl-go/texture"
	"rl-go/utils"
)

type LevelManager struct {
	ParsedLevels   map[string]loaders.ParsedLevel
	CurrentLevel   *Level
	TextureManager texture.TextureManager
}

func NewLevelManager(textureManager texture.TextureManager) *LevelManager {

	return &LevelManager{
		ParsedLevels:   loaders.ParsedLevelMap,
		CurrentLevel:   nil,
		TextureManager: textureManager,
	}
}

func (lm *LevelManager) StartLevel(name string) {
	name = utils.NormalizeName(name)

	if _, exists := lm.ParsedLevels[name]; !exists {
		log.Fatal("Level not found: " + name)
	}
	parsedLevel := lm.ParsedLevels[name]

	for _, textureName := range parsedLevel.Textures {
		lm.TextureManager.LoadTexture(textureName)
	}

	level := NewLevel(parsedLevel)

	lm.CurrentLevel = &level
}

func (lm *LevelManager) PlotLevel() {

	if lm.CurrentLevel == nil {
		log.Fatal("No level initialized")
	}

	for _, row := range lm.CurrentLevel.Map {
		for _, block := range row {
			block.Plot(lm.TextureManager)
		}
	}
}
