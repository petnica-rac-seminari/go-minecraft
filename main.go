package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	//"main/blocks"
	reljef "main/Reljef"
	"main/world"
)

func main() {
	rl.InitWindow(800, 600, "Raylib Go - 3D Kocka i Skakanje")
	defer rl.CloseWindow()

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 24.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 90.0
	camera.Projection = rl.CameraPerspective

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	var verticalVelocity float32 = 0.0
	const gravity float32 = -1
	const jumpForce float32 = 1
	const groundLevel float32 = 24.0
	var isGrounded bool = true
	var renderDistance int = 4

	var generatedChunks [][]world.Chunk

	generatedChunks = make([][]world.Chunk, renderDistance)
	for i := 0; i < renderDistance; i++ {
		generatedChunks[i] = make([]world.Chunk, renderDistance)
		for j := 0; j < renderDistance; j++ {
			generatedChunks[i][j] = reljef.GenerateChunk(i, j, 0.05, 8, 4, 0)
		}
	}

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		if rl.IsKeyPressed(rl.KeySpace) && isGrounded {
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

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		// rl.DrawCube(rl.NewVector3(0.0, 1.0, 0.0), 2.0, 2.0, 2.0, rl.Blue)
		// rl.DrawCubeWires(rl.NewVector3(0.0, 1.0, 0.0), 2.0, 2.0, 2.0, rl.DarkBlue)
		for i := 0; i < renderDistance; i++ {
			for j := 0; j < renderDistance; j++ {
				world.RenderChunk(generatedChunks[i][j], i, j)
			}
		}

		rl.DrawGrid(20, 1.0)

		rl.EndMode3D()

		rl.DrawFPS(10, 10)
		rl.DrawText("WASD - Kretanje | Mis - Okretanje | Space - Skakanje", 10, 40, 20, rl.DarkGray)

		rl.EndDrawing()
	}
}
