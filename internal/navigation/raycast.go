package navigation

import (
	"math"

	"main/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MaxReach = 6.0
	RayStep  = 0.05
)

type Hit struct {
	Found                  bool
	BlockX, BlockY, BlockZ int
	PlaceX, PlaceY, PlaceZ int
}

func Raycast(w *world.WorldStruct, origin, direction rl.Vector3) Hit {
	dir := rl.Vector3Normalize(direction)
	var prevX, prevY, prevZ int
	hasPrev := false

	for dist := float32(0); dist <= MaxReach; dist += RayStep {
		p := rl.Vector3Add(origin, rl.Vector3Scale(dir, dist))
		bx := int(math.Floor(float64(p.X)))
		by := int(math.Floor(float64(p.Y)))
		bz := int(math.Floor(float64(p.Z)))

		if w.Has(bx, by, bz) {
			h := Hit{
				Found:  true,
				BlockX: bx, BlockY: by, BlockZ: bz,
			}
			if hasPrev {
				h.PlaceX, h.PlaceY, h.PlaceZ = prevX, prevY, prevZ
			}
			return h
		}
		prevX, prevY, prevZ = bx, by, bz
		hasPrev = true
	}
	return Hit{}
}
