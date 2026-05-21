package navigation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CameraForward(cam *rl.Camera3D) rl.Vector3 {
	diff := rl.Vector3Subtract(cam.Target, cam.Position)
	length := rl.Vector3Length(diff)
	if length < 1e-6 {
		return rl.NewVector3(0, 0, -1)
	}
	return rl.Vector3Scale(diff, 1.0/length)
}

func CameraCardinal(cam *rl.Camera3D) string {
	f := CameraForward(cam)
	f.Y = 0
	if rl.Vector3Length(f) < 1e-6 {
		return "N"
	}
	f = rl.Vector3Normalize(f)
	if abs(f.X) > abs(f.Z) {
		if f.X > 0 {
			return "E"
		}
		return "W"
	}
	if f.Z > 0 {
		return "S"
	}
	return "N"
}

func abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
