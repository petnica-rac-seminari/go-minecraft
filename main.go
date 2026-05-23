package main

import (
	"math"
	"math/rand"

	reljef "main/Reljef"
	"main/blocks"
	nebo "main/dayNightCycle"
	"main/menu"
	"main/navigation"
	"main/oblaci"
	"main/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const render_dist = 4

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1920, 1080, "Raylib Go - 3D Kocka i Skakanje")
	defer rl.CloseWindow()

	menu.UcitajMenuSliku()
	defer menu.UnloadujMenuSliku()

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 40.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 90.0
	camera.Projection = rl.CameraPerspective
	rl.EnableCursor()
	rl.SetTargetFPS(60)
	rng := rand.New(rand.NewSource(int64(menu.Seed)))
	_ = rng

	var verticalVelocity float32 = 0.0
	const gravity float32 = -26.0
	const jumpForce float32 = 8.5
	var isGrounded bool = true
	var BlockToPlace blocks.Block = blocks.Grass

	const maxReach = navigation.DefaultMaxReach
	var lastHit navigation.RaycastHit
	var currentTick float32 = 0
	var clouds []oblaci.CLOUDS = oblaci.GenerateClouds()
	sunce_model := nebo.RenderSun()

	var jumpCtrl navigation.JumpInput
	const eyeHeight = navigation.DefaultEyeHeight

	for !rl.WindowShouldClose() {
		currentTick += rl.GetFrameTime()

		if !menu.IsMenu {
			rl.UpdateCamera(&camera, rl.CameraFirstPerson)

			playerCX := int(math.Floor(float64(camera.Position.X) / 16.0))
			playerCZ := int(math.Floor(float64(camera.Position.Z) / 16.0))

			halfDist := render_dist / 2
			for z := -halfDist; z <= halfDist; z++ {
				for x := -halfDist; x <= halfDist; x++ {
					pos := world.ChunkPos{X: playerCX + x, Z: playerCZ + z}

					if _, exists := world.LoadedChunks[pos]; !exists {
						c := reljef.GenerateChunk(pos.X*16, pos.Z*16, menu.Seed)
						world.LoadedChunks[pos] = &c
					}
				}
			}

			navigation.ApplyHorizontalCollision(&camera, eyeHeight, navigation.PlayerHalfWidth)

			dir := navigation.CameraDirection(camera)
			hit := navigation.Raycast(camera.Position, dir, maxReach)
			lastHit = hit

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

			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && hit.Hit {
				navigation.DestroyBlock(hit.X, hit.Y, hit.Z)
			}
			if rl.IsMouseButtonPressed(rl.MouseButtonRight) && hit.Hit {
				navigation.PlaceAdjacent(hit, BlockToPlace)
			}

			if rl.IsKeyPressed(rl.KeySpace) && isGrounded {
				verticalVelocity = jumpForce
				isGrounded = false
			}

			if navigation.IsAirborne(camera.Position, eyeHeight, navigation.PlayerHalfWidth) {
				isGrounded = false
			}

			canJump := isGrounded && !navigation.IsAirborne(camera.Position, eyeHeight, navigation.PlayerHalfWidth)
			if navigation.TryDoubleTapJump(&jumpCtrl, rl.GetTime(), rl.IsKeyPressed(rl.KeySpace), canJump) {
				verticalVelocity = jumpForce
				isGrounded = false
			}

			if !isGrounded {
				verticalVelocity += gravity * rl.GetFrameTime()
				camera.Position.Y += verticalVelocity * rl.GetFrameTime()
				camera.Target.Y += verticalVelocity * rl.GetFrameTime()
			}

			navigation.ApplyVerticalBlockPhysics(&camera, &verticalVelocity, &isGrounded, eyeHeight)
			oblaci.MoveClouds(clouds, camera)
		}

		rl.BeginDrawing()

		if !menu.IsMenu {
			rl.DisableCursor()
			rl.ClearBackground(nebo.SkyColor(int(currentTick)))

			rl.BeginMode3D(camera)

			playerCX := int(math.Floor(float64(camera.Position.X) / 16.0))
			playerCZ := int(math.Floor(float64(camera.Position.Z) / 16.0))
			halfDist := render_dist / 2

			for z := -halfDist; z <= halfDist; z++ {
				for x := -halfDist; x <= halfDist; x++ {
					pos := world.ChunkPos{X: playerCX + x, Z: playerCZ + z}
					if chunk, exists := world.LoadedChunks[pos]; exists {
						world.RenderChunk(*chunk)
					}
					if structure, exists := world.LoadedStructures[pos]; exists {
						world.RenderChunk(*structure)
					}
				}
			}

			oblaci.DrawClouds(clouds)
			if lastHit.Hit {
				navigation.DrawBlockOutline(lastHit.X, lastHit.Y, lastHit.Z, rl.Yellow)
			}

			rl.DrawModel(*sunce_model, nebo.MoveSun(float64(nebo.SkyBodyAngle(currentTick)), camera), 1.0, rl.White)

			rl.EndMode3D()

			rl.DrawFPS(10, 10)
			rl.DrawText("WASD - Kretanje | Mis - Okretanje | Space - Skok", 10, 40, 20, rl.DarkGray)
			rl.DrawText("LMB - Unisti | RMB - Postavi blok", 10, 70, 20, rl.DarkGray)
			rl.DrawText("1 - Grass | 2 - Dirt | 3 - Stone | 4 - Water | 5 - Snow", 10, 550, 20, rl.Black)
			rl.DrawCircle(int32(rl.GetScreenWidth())/2, int32(rl.GetScreenHeight())/2, 10, rl.White)

		} else {
			rl.ClearBackground(rl.DarkGray)
			menu.Crtaj()
		}

		rl.EndDrawing()
	}
}
