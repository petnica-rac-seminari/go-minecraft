package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"go-minecraft/blocks"
	nav "go-minecraft/internal/instructions"
	"go-minecraft/internal/world"
)

func RenderBlock(block blocks.Block, x, y, z int) {
	switch block {
	case blocks.Air:
		return
	case blocks.Grass:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.DarkGreen)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Green)
	case blocks.Stone:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Gray)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.DarkGray)
	case blocks.Dirt:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Brown)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.DarkBrown)
	case blocks.Water:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Blue)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.DarkBlue)
	case blocks.Bedrock:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Black)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Gray)
	case blocks.Snow:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.White)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Gray)
	}
}

func main() {
	rl.InitWindow(800, 600, "Petnica Minecraft")
	defer rl.CloseWindow()

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 2.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Projection = rl.CameraPerspective

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	w := world.NewDefault()

	var verticalVelocity float32 = 0.0
	const gravity float32 = -0.6
	const jumpForce float32 = 0.15
	const groundLevel float32 = 2.0
	var isGrounded bool = true

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		nav.HandleBlockInput(&camera, w)

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

		w.Draw()

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()
		rl.EndDrawing()
	}
}
