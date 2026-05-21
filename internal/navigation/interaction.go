package navigation

import (
	"fmt"
	"go-minecraft/internal/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	worldMinY = 0
	worldMaxY = 32
)

// HandleBlockInput: lijevi = ukloni, desni = postavi.
func HandleBlockInput(cam *rl.Camera3D, w *world.World) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		RemoveBlockAtCrosshair(cam, w)
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		PlaceBlockAtCrosshair(cam, w, rl.Gray)
	}
}

func RemoveBlockAtCrosshair(cam *rl.Camera3D, w *world.World) {
	hit := Raycast(w, cam.Position, CameraForward(cam))
	if !hit.Found {
		return
	}
	w.Remove(hit.BlockX, hit.BlockY, hit.BlockZ)
}

func PlaceBlockAtCrosshair(cam *rl.Camera3D, w *world.World, color rl.Color) {
	hit := Raycast(w, cam.Position, CameraForward(cam))
	if !hit.Found {
		return
	}
	if hit.PlaceY < worldMinY || hit.PlaceY >= worldMaxY {
		return
	}
	if w.Has(hit.PlaceX, hit.PlaceY, hit.PlaceZ) {
		return
	}
	if tooCloseToCamera(cam, w, hit.PlaceX, hit.PlaceY, hit.PlaceZ) {
		return
	}
	w.Set(hit.PlaceX, hit.PlaceY, hit.PlaceZ, color)
}

func tooCloseToCamera(cam *rl.Camera3D, w *world.World, x, y, z int) bool {
	c := w.BlockCenter(x, y, z)
	return rl.Vector3Distance(cam.Position, c) < 1.2
}

func DrawHUD(cam *rl.Camera3D, w *world.World) {
	rl.DrawText(fmt.Sprintf("Blocks: %d", w.Count()), 10, 10, 20, rl.Black)
	rl.DrawText(fmt.Sprintf("Facing: %s", CameraCardinal(cam)), 10, 40, 20, rl.Black)

	centerX := int32(400)
	centerY := int32(300)
	rl.DrawLine(centerX-10, centerY, centerX+10, centerY, rl.Black)
	rl.DrawLine(centerX, centerY-10, centerX, centerY+10, rl.Black)
}
