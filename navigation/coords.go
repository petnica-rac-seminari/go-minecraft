package navigation

import (
	"main/world"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func InBounds(x, y, z int) bool {
	return x >= 0 && x < ChunkSizeX &&
		y >= 0 && y < ChunkSizeY &&
		z >= 0 && z < ChunkSizeZ
}

func WorldToLocal(chunk world.Chunk, wx, wy, wz int) (lx, ly, lz int, ok bool) {
	lx = wx - chunk.GlobalX*ChunkSizeX
	ly = wy
	lz = wz - chunk.GlobalZ*ChunkSizeZ
	return lx, ly, lz, InBounds(lx, ly, lz)
}

func CameraDirection(camera rl.Camera3D) rl.Vector3 {
	dir := rl.Vector3Subtract(camera.Target, camera.Position)
	len := rl.Vector3Length(dir)
	if len < 1e-6 {
		return rl.NewVector3(0, 0, -1)
	}
	return rl.Vector3Scale(dir, 1.0/len)
}

func intFloor(v float32) int {
	return int(math.Floor(float64(v)))
}
