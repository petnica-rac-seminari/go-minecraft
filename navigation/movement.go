package navigation

import (
	//"fmt"
	"math"

	//"main/blocks"
	"main/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func vector3Subtract(v1, v2 rl.Vector3) rl.Vector3 {
	return rl.NewVector3(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}

func cameraAngle(camera rl.Camera3D) rl.Vector3 {
	return vector3Subtract(camera.Target, camera.Position)
}

// 1. FIX: Give these actual values!
// In 3D: Gravity pulls DOWN (-Y), Jumping pushes UP (+Y)
const (
	gravity   = -35.0 // Acceleration downwards per second
	jumpForce = 8.5   // Instant upward velocity boost
)

var speed float32 = 5.0

var yVelocity float32
var last bool

var grounded bool

// Helper to correctly floor float32 coordinates to integer block indices
// This fixes bugs when player travels into negative coordinates
func blockCoord(val float32) int {
	return int(math.Floor(float64(val)))
}

func GetPlayerBoundingBox(camPos rl.Vector3) rl.BoundingBox {
	radiusX := float32(0.35)
	radiusZ := float32(0.35)
	playerHeight := float32(1.8)
	eyeOffset := float32(1.6) // Distance from eyes (camera) down to feet

	feetY := camPos.Y - eyeOffset

	return rl.BoundingBox{
		Min: rl.NewVector3(camPos.X-radiusX, feetY, camPos.Z-radiusZ),
		Max: rl.NewVector3(camPos.X+radiusX, feetY+playerHeight, camPos.Z+radiusZ),
	}
}

func HandleMovement(camera *rl.Camera3D) {
	var mousePositionDelta = rl.GetMouseDelta()
	deltaTime := rl.GetFrameTime()

	var rotateUp uint8

	// Camera rotation
	rl.CameraYaw(camera, -mousePositionDelta.X*0.003, 0)
	rl.CameraPitch(camera, -mousePositionDelta.Y*0.003, 1, 0, rotateUp)

	// --- PHYSICS & VERTICAL MOVEMENT ---

	// First, check if we are standing on something
	grounded = IsGrounded(camera.Position)

	if grounded {
		// If we are grounded and moving downwards, stop falling
		if yVelocity < 0 {
			yVelocity = 0
		}

		// Jump input handling
		if rl.IsKeyPressed(rl.KeySpace) {
			yVelocity = jumpForce
			grounded = false
		}
	} else {
		// Apply gravity over time (subtracting because down is -Y)
		yVelocity += gravity * deltaTime
	}

	// 2. FIX: You have to actually move the camera and its target!
	movementY := yVelocity * deltaTime
	camera.Position.Y += movementY
	camera.Target.Y += movementY

	// 3. FIX: Snap to floor if we clipped slightly below it while landing
	if grounded && yVelocity <= 0 {
		eyeOffset := float32(1.6)
		// Get the integer block height of the block directly under our feet
		feetY := camera.Position.Y - eyeOffset
		currentBlockY := blockCoord(feetY - 0.01)

		// Perfect floor position is the top of that block (block index + 1)
		topOfFloor := float32(currentBlockY + 1)

		// Snap camera perfectly so you don't jitter or sink
		camera.Position.Y = topOfFloor + eyeOffset
	}

	//
	//

	forwardX := camera.Target.X - camera.Position.X
	forwardZ := camera.Target.Z - camera.Position.Z

	// Normalize forward vector so diagonal movement isn't faster
	mag := float32(math.Sqrt(float64(forwardX*forwardX + forwardZ*forwardZ)))
	if mag > 0 {
		forwardX /= mag
		forwardZ /= mag
	}

	// Right vector is perpendicular to the forward vector
	rightX := -forwardZ
	rightZ := forwardX

	// 3. Accumulate Keyboard Input
	var moveDirX, moveDirZ float32
	if rl.IsKeyDown(rl.KeyW) {
		moveDirX += forwardX
		moveDirZ += forwardZ
	}
	if rl.IsKeyDown(rl.KeyS) {
		moveDirX -= forwardX
		moveDirZ -= forwardZ
	}
	if rl.IsKeyDown(rl.KeyD) {
		moveDirX += rightX
		moveDirZ += rightZ
	}
	if rl.IsKeyDown(rl.KeyA) {
		moveDirX -= rightX
		moveDirZ -= rightZ
	}

	// 4. Test and Apply Movement (with Frame Rate Independence)
	moveX := moveDirX * speed * deltaTime
	moveZ := moveDirZ * speed * deltaTime

	// Slide on X Axis
	if moveX != 0 {
		testPos := camera.Position
		testPos.X += moveX
		if !CheckHorizontalCollision(testPos) {
			camera.Position.X += moveX
			camera.Target.X += moveX
		}
	}

	// Slide on Z Axis
	if moveZ != 0 {
		testPos := camera.Position
		testPos.Z += moveZ
		if !CheckHorizontalCollision(testPos) {
			camera.Position.Z += moveZ
			camera.Target.Z += moveZ
		}
	}

	// Keyboard movement (X and Z axis)
	/*if rl.IsKeyDown(rl.KeyW) {
		rl.CameraMoveForward(camera, 0.09, 1)
	}
	if rl.IsKeyDown(rl.KeyA) {
		rl.CameraMoveRight(camera, -0.09, 1)
	}
	if rl.IsKeyDown(rl.KeyS) {
		rl.CameraMoveForward(camera, -0.09, 1)
	}
	if rl.IsKeyDown(rl.KeyD) {
		rl.CameraMoveRight(camera, 0.09, 1)
	}*/

}

func IsGrounded(camPos rl.Vector3) bool {
	bbox := GetPlayerBoundingBox(camPos)

	// Check just a hair below the feet bounding box
	checkY := bbox.Min.Y - 0.02

	// 4. FIX: Use math.Floor via our helper to prevent native int casting errors
	minX := blockCoord(bbox.Min.X)
	maxX := blockCoord(bbox.Max.X)
	minZ := blockCoord(bbox.Min.Z)
	maxZ := blockCoord(bbox.Max.Z)
	gridY := blockCoord(checkY)

	// Check all blocks underneath the player's horizontal bounding box footprint
	for x := minX; x <= maxX; x++ {
		for z := minZ; z <= maxZ; z++ {
			if world.GetGlobalBlock(x, gridY, z) >= 2 {
				return true
			}
		}
	}

	return false
}

func CheckHorizontalCollision(camPos rl.Vector3) bool {
	bbox := GetPlayerBoundingBox(camPos)

	// Get integer block ranges that the player overlaps
	offset := float32(0.25)

	minX := blockCoord(bbox.Min.X + offset)
	maxX := blockCoord(bbox.Max.X + offset)
	minY := blockCoord(bbox.Min.Y + 0.5) // Slight offset so we don't catch the floor
	maxY := blockCoord(bbox.Max.Y - 0.5) // Slight offset so we don't catch the ceiling
	minZ := blockCoord(bbox.Min.Z + offset + 0.5)
	maxZ := blockCoord(bbox.Max.Z + offset + 0.5)

	// Loop through all blocks in the player's current vertical and horizontal space
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			for z := minZ; z <= maxZ; z++ {
				if world.GetGlobalBlock(x, y, z) >= 2 {
					return true // Collision detected!
				}
			}
		}
	}
	return false
}
