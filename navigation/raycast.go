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

func Raycast(origin, direction rl.Vector3, maxDist float32) RaycastHit {
	steps := int(maxDist * 5)
	stepSize := maxDist / float32(steps)

	for i := 0; i < steps; i++ {
		currentPos := rl.Vector3Add(origin, rl.Vector3Scale(direction, float32(i)*stepSize))

		x := int(math.Floor(float64(currentPos.X)))
		y := int(math.Floor(float64(currentPos.Y)))
		z := int(math.Floor(float64(currentPos.Z)))

		if world.GetGlobalBlock(x, y, z) != blocks.Air {
			prevPos := rl.Vector3Add(origin, rl.Vector3Scale(direction, float32(i-1)*stepSize))
			px := int(math.Floor(float64(prevPos.X)))
			py := int(math.Floor(float64(prevPos.Y)))
			pz := int(math.Floor(float64(prevPos.Z)))

			return RaycastHit{
				Hit:      true,
				X:        x,
				Y:        y,
				Z:        z,
				NX:       px - x,
				NY:       py - y,
				NZ:       pz - z,
				Distance: float32(i) * stepSize,
			}
		}
	}
	return RaycastHit{}
}

func DestroyBlock(x, y, z int) bool {
	return world.SetGlobalBlock(x, y, z, blocks.Air)
}

func PlaceAdjacent(hit RaycastHit, b blocks.Block) bool {
	if !hit.Hit || b == blocks.Air {
		return false
	}
	return world.SetGlobalBlock(hit.X+hit.NX, hit.Y+hit.NY, hit.Z+hit.NZ, b)
}
