package navigation

import (
	"main/blocks"
	"main/world"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RaycastHit struct {
	Hit        bool
	X, Y, Z    int
	NX, NY, NZ int
	Distance   float32
}

func Raycast(chunk world.Chunk, origin, direction rl.Vector3, maxDist float32) RaycastHit {
	dir := direction
	length := rl.Vector3Length(dir)
	if length < 1e-6 {
		return RaycastHit{}
	}
	if math.Abs(float64(length)-1.0) > 0.01 {
		dir = rl.Vector3Scale(dir, 1.0/length)
	}

	x := intFloor(origin.X)
	y := intFloor(origin.Y)
	z := intFloor(origin.Z)

	stepX, stepY, stepZ := 0, 0, 0
	var tDeltaX, tDeltaY, tDeltaZ float32
	tMaxX, tMaxY, tMaxZ := float32(math.MaxFloat32), float32(math.MaxFloat32), float32(math.MaxFloat32)

	if dir.X > 0 {
		stepX = 1
		tDeltaX = 1.0 / dir.X
		tMaxX = (float32(x+1) - origin.X) / dir.X
	} else if dir.X < 0 {
		stepX = -1
		tDeltaX = -1.0 / dir.X
		tMaxX = (origin.X - float32(x)) / -dir.X
	} else {
		tDeltaX = float32(math.MaxFloat32)
	}

	if dir.Y > 0 {
		stepY = 1
		tDeltaY = 1.0 / dir.Y
		tMaxY = (float32(y+1) - origin.Y) / dir.Y
	} else if dir.Y < 0 {
		stepY = -1
		tDeltaY = -1.0 / dir.Y
		tMaxY = (origin.Y - float32(y)) / -dir.Y
	} else {
		tDeltaY = float32(math.MaxFloat32)
	}

	if dir.Z > 0 {
		stepZ = 1
		tDeltaZ = 1.0 / dir.Z
		tMaxZ = (float32(z+1) - origin.Z) / dir.Z
	} else if dir.Z < 0 {
		stepZ = -1
		tDeltaZ = -1.0 / dir.Z
		tMaxZ = (origin.Z - float32(z)) / -dir.Z
	} else {
		tDeltaZ = float32(math.MaxFloat32)
	}

	prevStepX, prevStepY, prevStepZ := 0, 0, 0

	for {
		lx, ly, lz, ok := WorldToLocal(chunk, x, y, z)
		if !ok {
			return RaycastHit{}
		}

		if chunk.Blocks[lx][ly][lz] != blocks.Air {
			return RaycastHit{
				Hit:      true,
				X:        lx,
				Y:        ly,
				Z:        lz,
				NX:       -prevStepX,
				NY:       -prevStepY,
				NZ:       -prevStepZ,
				Distance: nextT(tMaxX, tMaxY, tMaxZ, tDeltaX, tDeltaY, tDeltaZ),
			}
		}

		if tMaxX < tMaxY {
			if tMaxX < tMaxZ {
				if tMaxX > maxDist {
					return RaycastHit{}
				}
				prevStepX, prevStepY, prevStepZ = stepX, 0, 0
				x += stepX
				tMaxX += tDeltaX
			} else {
				if tMaxZ > maxDist {
					return RaycastHit{}
				}
				prevStepX, prevStepY, prevStepZ = 0, 0, stepZ
				z += stepZ
				tMaxZ += tDeltaZ
			}
		} else {
			if tMaxY < tMaxZ {
				if tMaxY > maxDist {
					return RaycastHit{}
				}
				prevStepX, prevStepY, prevStepZ = 0, stepY, 0
				y += stepY
				tMaxY += tDeltaY
			} else {
				if tMaxZ > maxDist {
					return RaycastHit{}
				}
				prevStepX, prevStepY, prevStepZ = 0, 0, stepZ
				z += stepZ
				tMaxZ += tDeltaZ
			}
		}
	}
}

func nextT(tMaxX, tMaxY, tMaxZ, tDeltaX, tDeltaY, tDeltaZ float32) float32 {
	t := tMaxX - tDeltaX
	if tMaxY-tDeltaY < t {
		t = tMaxY - tDeltaY
	}
	if tMaxZ-tDeltaZ < t {
		t = tMaxZ - tDeltaZ
	}
	if t < 0 {
		return 0
	}
	return t
}
