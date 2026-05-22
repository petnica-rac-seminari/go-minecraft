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

func UpdateCamera(camera *rl.Camera, mode rl.CameraMode) {
	var mousePositionDelta = rl.GetMouseDelta()

	moveInWorldPlaneBool := mode == rl.CameraFirstPerson || mode == rl.CameraThirdPerson
	var moveInWorldPlane uint8
	if moveInWorldPlaneBool {
		moveInWorldPlane = 1
	}

	rotateAroundTargetBool := mode == rl.CameraThirdPerson || mode == rl.CameraOrbital
	var rotateAroundTarget uint8
	if rotateAroundTargetBool {
		rotateAroundTarget = 1
	}

	lockViewBool := mode == rl.CameraFirstPerson || mode == rl.CameraThirdPerson || mode == rl.CameraOrbital
	var lockView uint8
	if lockViewBool {
		lockView = 1
	}

	var rotateUp uint8

	if mode == rl.CameraOrbital {
		// Orbital can just orbit
		var rotation = rl.MatrixRotate(rl.GetCameraUp(camera), 0.5*rl.GetFrameTime())
		var view = rl.Vector3Subtract(camera.Position, camera.Target)
		view = rl.Vector3Transform(view, rotation)
		camera.Position = rl.Vector3Add(camera.Target, view)
	} else {
		// Camera rotation
		if rl.IsKeyDown(rl.KeyDown) {
			rl.CameraPitch(camera, -0.03, lockView, rotateAroundTarget, rotateUp)
		}
		if rl.IsKeyDown(rl.KeyUp) {
			rl.CameraPitch(camera, 0.03, lockView, rotateAroundTarget, rotateUp)
		}
		if rl.IsKeyDown(rl.KeyRight) {
			rl.CameraYaw(camera, -0.03, rotateAroundTarget)
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			rl.CameraYaw(camera, 0.03, rotateAroundTarget)
		}

		// Camera movement
		if !(rl.IsGamepadAvailable(0)) {
			// Camera pan (for CameraFree)
			if mode == rl.CameraFree && rl.IsMouseButtonDown(rl.MouseMiddleButton) {
				var mouseDelta = rl.GetMouseDelta()
				if mouseDelta.X > 0.0 {
					rl.CameraMoveRight(camera, 0.2, moveInWorldPlane)
				}
				if mouseDelta.X < 0.0 {
					rl.CameraMoveRight(camera, -0.2, moveInWorldPlane)
				}
				if mouseDelta.Y > 0.0 {
					rl.CameraMoveUp(camera, -0.2)
				}
				if mouseDelta.Y < 0.0 {
					rl.CameraMoveUp(camera, 0.2)
				}
			} else {
				// Mouse support
				rl.CameraYaw(camera, -mousePositionDelta.X*0.003, rotateAroundTarget)
				rl.CameraPitch(camera, -mousePositionDelta.Y*0.003, lockView, rotateAroundTarget, rotateUp)
			}

			// Keyboard support
			if rl.IsKeyDown(rl.KeyW) {
				rl.CameraMoveForward(camera, 0.09, moveInWorldPlane)
			}
			if rl.IsKeyDown(rl.KeyA) {
				rl.CameraMoveRight(camera, -0.09, moveInWorldPlane)
			}
			if rl.IsKeyDown(rl.KeyS) {
				rl.CameraMoveForward(camera, -0.09, moveInWorldPlane)
			}
			if rl.IsKeyDown(rl.KeyD) {
				rl.CameraMoveRight(camera, 0.09, moveInWorldPlane)
			}
		} else {
			// Gamepad controller support
			rl.CameraYaw(camera, -(rl.GetGamepadAxisMovement(0, rl.GamepadAxisRightX)*float32(2))*0.003, rotateAroundTarget)
			rl.CameraPitch(camera, -(rl.GetGamepadAxisMovement(0, rl.GamepadAxisRightY)*float32(2))*0.003, lockView, rotateAroundTarget, rotateUp)

			if rl.GetGamepadAxisMovement(0, rl.GamepadAxisLeftY) <= -0.25 {
				rl.CameraMoveForward(camera, 0.09, moveInWorldPlane)
			}
			if rl.GetGamepadAxisMovement(0, rl.GamepadAxisLeftX) <= -0.25 {
				rl.CameraMoveRight(camera, -0.09, moveInWorldPlane)
			}
			if rl.GetGamepadAxisMovement(0, rl.GamepadAxisLeftY) >= 0.25 {
				rl.CameraMoveForward(camera, -0.09, moveInWorldPlane)
			}
			if rl.GetGamepadAxisMovement(0, rl.GamepadAxisLeftX) >= 0.25 {
				rl.CameraMoveRight(camera, 0.09, moveInWorldPlane)
			}
		}

		if mode == rl.CameraFree {
			if rl.IsKeyDown(rl.KeySpace) {
				rl.CameraMoveUp(camera, 0.09)
			}
			if rl.IsKeyDown(rl.KeyLeftControl) {
				rl.CameraMoveUp(camera, -0.09)
			}
		}
	}

	if mode == rl.CameraThirdPerson || mode == rl.CameraOrbital || mode == rl.CameraFree {
		// Zoom target distance
		rl.CameraMoveToTarget(camera, -rl.GetMouseWheelMove())
		if rl.IsKeyPressed(rl.KeyKpSubtract) {
			rl.CameraMoveToTarget(camera, 2.0)
		}
		if rl.IsKeyPressed(rl.KeyKpAdd) {
			rl.CameraMoveToTarget(camera, -2.0)
		}
	}
}

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
					c := reljef.GenerateChunk(pos.X*16, pos.Z*16, 0.1, 4, 4, 12345)
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
			BlockToPlace = blocks.Stone
		case rl.KeyThree:
			BlockToPlace = blocks.Dirt
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
		rl.DrawText("WASD - Kretanje | Mis - Okretanje | Space - Skakanje", 10, 40, 20, rl.DarkGray)
		rl.DrawText("LMB - Ukloni | RMB - Postavi Grass", 10, 70, 20, rl.DarkGray)
		rl.DrawText("Space x2 - Skok | Kretanje po blokovima", 10, 100, 20, rl.DarkGray)
		rl.DrawText("1-grass | 2-stone | 3-dirt | 4-water | 5-snow", 10, 1000, 20, rl.White)

		rl.EndDrawing()
	}
}
