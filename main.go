package main

import (
	//"fmt"
	"math"

	reljef "main/Reljef"
	nav "main/navigation"
	"main/oblaci"

	nebo "main/dayNightCycle"

	rl "github.com/gen2brain/raylib-go/raylib"

	"main/blocks"
	//"main/navigation"
	"main/world"
)

const render_dist = 4

func main() {
	rl.InitWindow(1600, 900, "Raylib Go - 3D Kocka i Skakanje")
	defer rl.CloseWindow()

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 40.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Projection = rl.CameraPerspective
	rl.DisableCursor()
	rl.SetTargetFPS(60)

	var BlockToPlace blocks.Block = blocks.Grass

	var time float32 = 0
	var clouds []oblaci.CLOUDS = oblaci.GenerateClouds()

	for !rl.WindowShouldClose() {
		time += rl.GetFrameTime()

		nav.HandleBlockManipulation(&camera)
		nav.HandleMovement(&camera)

		playerCX := int(math.Floor(float64(camera.Position.X) / 16.0))
		playerCZ := int(math.Floor(float64(camera.Position.Z) / 16.0))

		halfDist := render_dist / 2
		for z := -halfDist; z <= halfDist; z++ {
			for x := -halfDist; x <= halfDist; x++ {
				pos := world.ChunkPos{X: playerCX + x, Z: playerCZ + z}

				if _, exists := world.LoadedChunks[pos]; !exists {
					c := reljef.GenerateChunk(pos.X*16, pos.Z*16)
					world.LoadedChunks[pos] = &c
				}
			}
		}

		switch rl.GetKeyPressed() {
		case rl.KeyOne:
			BlockToPlace = blocks.Grass
		case rl.KeyTwo:
			BlockToPlace = blocks.Dirt
		case rl.KeyThree:
			BlockToPlace = blocks.Stone
		case rl.KeyFour:
			BlockToPlace = blocks.Water
		case rl.KeyFive:
			BlockToPlace = blocks.Snow
		}
		if BlockToPlace == blocks.Air {
			BlockToPlace = blocks.Bedrock
		}

		rl.BeginDrawing()
		rl.ClearBackground(nebo.SkyColor(int(time)))

		rl.BeginMode3D(camera)

		for z := -halfDist; z <= halfDist; z++ {
			for x := -halfDist; x <= halfDist; x++ {
				pos := world.ChunkPos{X: playerCX + x, Z: playerCZ + z}
				if chunk, exists := world.LoadedChunks[pos]; exists {
					world.RenderChunk(*chunk)
				}
			}
		}

		oblaci.RenderCloud(clouds)

		rl.EndMode3D()

		rl.DrawFPS(10, 10)
		rl.DrawText("WASD - Kretanje | Mis - Okretanje | Space - Skok", 10, 40, 20, rl.DarkGray)
		rl.DrawText("LMB - Unisti | RMB - Postavi blok", 10, 70, 20, rl.DarkGray)
		rl.DrawText("1 - Grass | 2 - Dirt | 3 - Stone | 4 - Water | 5 - Snow", 10, 550, 20, rl.Black)

		rl.EndDrawing()
	}
}
