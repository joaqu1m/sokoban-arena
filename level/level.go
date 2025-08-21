package level

import (
	"rl-go/block"
	"rl-go/loaders"
	"rl-go/texture"
)

type Level struct {
	Title   string          `json:"title"`
	Author  string          `json:"author"`
	Group   string          `json:"group"`
	Map     [][]block.Block `json:"map"`
	Texture texture.Texture `json:"texture"`
}

func NewLevel(levelMap loaders.ParsedLevel) Level {

	rowCount := len(levelMap.Map)
	colCount := len(levelMap.Map[0])

	blocksMatrix := [][]block.Block{}

	for rowIndex, row := range levelMap.Map {
		blockRow := []block.Block{}
		for colIndex, blockName := range row {

			parsedBlock := loaders.GetParsedBlock(blockName)

			blockRow = append(blockRow, block.NewBlock(
				float32((colIndex-(colCount/2))*int(parsedBlock.DefaultSize.H)),
				float32((rowIndex-(rowCount/2))*int(parsedBlock.DefaultSize.W)),
				parsedBlock.DefaultSize.W,
				parsedBlock.DefaultSize.H,
				parsedBlock.Texture,
			))

		}

		blocksMatrix = append(blocksMatrix, blockRow)
	}

	return Level{
		Title:  levelMap.Title,
		Author: levelMap.Author,
		Group:  levelMap.Group,
		Map:    blocksMatrix,
	}
}
