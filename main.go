package main

import (
	reljef "main/Reljef"
	nebo "main/nebo"
	"main/oblaci"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	//"main/blocks"
	"main/world"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	rl.InitWindow(1920, 1080, "Raylib Go - 3D Kocka i Skakanje")
	defer rl.CloseWindow()

	// startTime := time.Now()

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
	const groundLevel float32 = 10.0
	var isGrounded bool = true
	var currentTick float32 = 0
	var clouds []oblaci.CLOUDS = oblaci.GenerateClouds()
	nebo.GenerateStars(150)

	sunce_model := nebo.RenderSun()
	moonModel := nebo.RenderMoon()

	generatedChunk := reljef.GenerateChunk(0, 0, 0.1, 8, 0, 1)

	// fmt.Println(clouds)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)
		currentTick += rl.GetFrameTime()
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

		oblaci.MoveClouds(clouds)

		rl.BeginDrawing()
		rl.ClearBackground(nebo.SkyColor(int64(currentTick)))

		rl.BeginMode3D(camera)

		// rl.DrawCube(rl.NewVector3(0.0, 1.0, 0.0), 2.0, 2.0, 2.0, rl.Blue)
		// rl.DrawCubeWires(rl.NewVector3(0.0, 1.0, 0.0), 2.0, 2.0, 2.0, rl.DarkBlue)
		world.RenderChunk(generatedChunk)
		oblaci.DrawClouds(clouds, camera.Position)

		if nebo.IsSunVisible(int64(currentTick)) {
			rl.DrawModel(*sunce_model, nebo.MoveSunByAngle(nebo.SkyBodyAngle(currentTick), camera), 1.0, rl.White)
		}

		if nebo.IsMoonVisible(int64(currentTick)) {
			rl.DrawModel(*moonModel, nebo.MoveMoonByAngle(nebo.SkyBodyAngle(currentTick), camera), 1.0, rl.White)
		}

		if nebo.IsNight(int64(currentTick)) {
			nebo.DrawStars(camera)
		}

		rl.DrawGrid(20, 1.0)

		// oblaci.DrawClouds()
		rl.EndMode3D()

		rl.DrawFPS(10, 10)
		rl.DrawText("WASD - Kretanje | Mis - Okretanje | Space - Skakanje", 10, 40, 20, rl.DarkGray)
		rl.EndDrawing()
	}
}
