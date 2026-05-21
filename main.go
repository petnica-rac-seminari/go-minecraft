package main

import (
	reljef "main/Reljef"

	rl "github.com/gen2brain/raylib-go/raylib"

	"main/blocks"
	"main/navigation"
	"main/world"
)

func main() {
	rl.InitWindow(800, 600, "Raylib Go - 3D Kocka i Skakanje")
	defer rl.CloseWindow()

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 10.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Projection = rl.CameraPerspective

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	var verticalVelocity float32 = 0.0
	const gravity float32 = -0.6
	const jumpForce float32 = 0.15
	const groundLevel float32 = 1.0
	var isGrounded bool = true

	generatedChunk := reljef.GenerateChunk(0, 0, 0.1, 8, 0, 1)

	const maxReach = navigation.DefaultMaxReach
	var lastHit navigation.RaycastHit

	var jumpCtrl navigation.JumpInput
	const eyeHeight = navigation.DefaultEyeHeight

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		navigation.ApplyHorizontalCollision(&camera, generatedChunk, eyeHeight, navigation.PlayerHalfWidth)

		dir := navigation.CameraDirection(camera)
		hit := navigation.Raycast(generatedChunk, camera.Position, dir, maxReach)
		lastHit = hit

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && hit.Hit {
			navigation.DestroyBlock(&generatedChunk, hit.X, hit.Y, hit.Z)
		}
		if rl.IsMouseButtonPressed(rl.MouseButtonRight) && hit.Hit {
			navigation.PlaceAdjacent(&generatedChunk, hit, blocks.Grass)
		}

		// if rl.IsKeyPressed(rl.KeySpace) && isGrounded {
		// 	verticalVelocity = jumpForce
		// 	isGrounded = false
		// }

		if navigation.IsAirborne(generatedChunk, camera.Position, eyeHeight, navigation.PlayerHalfWidth) {
			isGrounded = false
		}

		canJump := isGrounded && !navigation.IsAirborne(generatedChunk, camera.Position, eyeHeight, navigation.PlayerHalfWidth)
		if navigation.TryDoubleTapJump(&jumpCtrl, rl.GetTime(), rl.IsKeyPressed(rl.KeySpace), canJump) {
			verticalVelocity = jumpForce
			isGrounded = false
		}

		if !isGrounded {
			verticalVelocity += gravity * rl.GetFrameTime()
			camera.Position.Y += verticalVelocity
			camera.Target.Y += verticalVelocity

			if camera.Position.Y <= groundLevel {
				diff := groundLevel - camera.Position.Y
				camera.Position.Y = groundLevel
				camera.Target.Y += diff
				verticalVelocity = 0.0
				isGrounded = true
			}
		}

		navigation.ApplyVerticalBlockPhysics(&camera, &verticalVelocity, &isGrounded, generatedChunk, eyeHeight)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		// rl.DrawCube(rl.NewVector3(0.0, 1.0, 0.0), 2.0, 2.0, 2.0, rl.Blue)
		// rl.DrawCubeWires(rl.NewVector3(0.0, 1.0, 0.0), 2.0, 2.0, 2.0, rl.DarkBlue)
		world.RenderChunk(generatedChunk)

		if lastHit.Hit {
			navigation.DrawBlockOutline(lastHit.X, lastHit.Y, lastHit.Z, rl.Yellow)
		}

		rl.DrawGrid(20, 1.0)

		rl.EndMode3D()

		rl.DrawFPS(10, 10)
		rl.DrawText("WASD - Kretanje | Mis - Okretanje | Space - Skakanje", 10, 40, 20, rl.DarkGray)
		rl.DrawText("LMB - Ukloni | RMB - Postavi Grass", 10, 70, 20, rl.DarkGray)
		rl.DrawText("Space x2 - Skok | Kretanje po blokovima", 10, 100, 20, rl.DarkGray)

		rl.EndDrawing()
	}
}
