package main

import (
	"main/blocks"

	rl "github.com/gen2brain/raylib-go/raylib"
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
	// 1. Inicijalizacija prozora
	rl.InitWindow(800, 600, "Petnica Minecraft")
	defer rl.CloseWindow()

	// 2. Podešavanje 3D kamere
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 2.0, 4.0) // Početna pozicija kamere
	camera.Target = rl.NewVector3(0.0, 1.0, 0.0)   // Tačka u koju kamera gleda
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)       // Vektor koji definiše gde je "gore"
	camera.Fovy = 60.0                             // Vidno polje (Field of View)
	camera.Projection = rl.CameraPerspective       // Vrsta projekcije

	// Sakrivamo i zaključavamo kursor kako bismo mogli normalno da se okrećemo mišem (kao u FPS igrama)
	rl.DisableCursor()

	rl.SetTargetFPS(60)

	var verticalVelocity float32 = 0.0
	const gravity float32 = -0.6
	const jumpForce float32 = 0.15
	const groundLevel float32 = 2.0
	var isGrounded bool = true

	// 3. Glavna petlja
	for !rl.WindowShouldClose() {
		// Automatski ažurira kameru na osnovu WASD tastera i pokreta miša
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

		// --- POČETAK CRTANJA ---
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Aktivacija 3D moda
		rl.BeginMode3D(camera)

		// Crtamo jednu plavu kocku na poziciji (0, 1, 0)
		rl.DrawCube(rl.NewVector3(0.0, 1.0, 0.0), 2.0, 2.0, 2.0, rl.Blue)
		// Crtamo ivice kocke kako bi se lakše video 3D oblik
		rl.DrawCubeWires(rl.NewVector3(0.0, 1.0, 0.0), 2.0, 2.0, 2.0, rl.DarkBlue)

		// Crtamo mrežu na tlu kako bismo imali osećaj za prostor i kretanje
		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()
		rl.EndDrawing()
	}
}
