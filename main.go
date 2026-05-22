package main

import (
	"math"

	reljef "main/Reljef"

	nebo "main/dayNightCycle"

	rl "github.com/gen2brain/raylib-go/raylib"

	"main/blocks"
	"main/navigation"
	"main/world"
)

const render_dist = 4

func main() {
	rl.InitWindow(1920, 1080, "Raylib Go - 3D Kocka i Skakanje")
	defer rl.CloseWindow()

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 40.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Projection = rl.CameraPerspective

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	var verticalVelocity float32 = 0.0
	const gravity float32 = -26.0
	const jumpForce float32 = 8.5
	var isGrounded bool = true
	var BlockToPlace blocks.Block = blocks.Grass

	const maxReach = navigation.DefaultMaxReach
	var lastHit navigation.RaycastHit
	var time float32 = 0

	var jumpCtrl navigation.JumpInput
	const eyeHeight = navigation.DefaultEyeHeight

	for !rl.WindowShouldClose() {
		time += rl.GetFrameTime()
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

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

		c := nebo.SkyColor(int(time))
		navigation.ApplyVerticalBlockPhysics(&camera, &verticalVelocity, &isGrounded, eyeHeight)

		rl.BeginDrawing()
		rl.ClearBackground(c)

		rl.BeginMode3D(camera)

		for z := -halfDist; z <= halfDist; z++ {
			for x := -halfDist; x <= halfDist; x++ {
				pos := world.ChunkPos{X: playerCX + x, Z: playerCZ + z}
				if chunk, exists := world.LoadedChunks[pos]; exists {
					world.RenderChunk(*chunk)
				}
			}
		}

		if lastHit.Hit {
			navigation.DrawBlockOutline(lastHit.X, lastHit.Y, lastHit.Z, rl.Yellow)
		}

		rl.EndMode3D()

		rl.DrawFPS(10, 10)
		rl.DrawText("WASD - Kretanje | Mis - Okretanje | Space - Skok", 10, 40, 20, rl.DarkGray)
		rl.DrawText("LMB - Unisti | RMB - Postavi blok", 10, 70, 20, rl.DarkGray)
		rl.DrawText("1 - Grass | 2 - Dirt | 3 - Stone | 4 - Water | 5 - Snow", 10, 550, 20, rl.Black)

		rl.EndDrawing()
	}
}
