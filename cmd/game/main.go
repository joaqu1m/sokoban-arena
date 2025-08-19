package main

import (
	"log"
	"rl-go/level"
	"rl-go/texture"
	"rl-go/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SokobanBlock struct {
	X            float32
	Y            float32
	TextureIndex int
}

const MAX_DELTA_TIME float32 = 0.05
const BLOCK_SIZE float32 = 32

var GAME_SIZE = utils.Size{
	W: 800,
	H: 800,
}

func main() {
	// rl.SetConfigFlags(rl.FlagFullscreenMode); // fullscreen
	// rl.SetConfigFlags(rl.FlagWindowTopmost | rl.FlagWindowUndecorated) // windowed fullscreen

	log.Println("Starting Sokoban Game...")

	rl.InitWindow(GAME_SIZE.W, GAME_SIZE.H, "Sokoban Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	textureManager := texture.NewTextureManager(false)
	defer textureManager.UnloadAllTextures()

	levelManager := level.NewLevelManager(textureManager)

	for !rl.WindowShouldClose() {
		deltaTime := rl.GetFrameTime()

		if deltaTime > MAX_DELTA_TIME {
			deltaTime = MAX_DELTA_TIME
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if levelManager.CurrentLevel == nil {
			levelManager.StartLevel("tutorial-1")
			log.Println("Initialized tutorial-1")
		}

		levelManager.PlotLevel()

		rl.EndDrawing()
	}
}
