package navigation

import (
	"main/blocks"
	"main/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func blockTopY(blockIndexY int) float32 {
	return float32(blockIndexY) + 0.5
}

func isSolid(b blocks.Block) bool {
	if b == blocks.Air {
		return false
	} else if b == blocks.Water {
		return false
	}
	return true
}
func columnGroundY(worldX, worldZ float32, feetY float32) (groundY float32, ok bool) {
	wx := intFloor(worldX)
	wz := intFloor(worldZ)

	// Počinjemo pretragu od bloka u kojem su noge/struk prema dole
	startWY := intFloor(feetY + 0.5)
	if startWY >= ChunkSizeY {
		startWY = ChunkSizeY - 1
	}
	if startWY < 0 {
		return 0, false
	}

	for wy := startWY; wy >= 0; wy-- {
		if world.GetGlobalBlock(wx, wy, wz) != blocks.Air && world.GetGlobalBlock(wx, wy, wz) != blocks.Water {
			return blockTopY(wy), true
		}
	}
	return 0, false
}

func HighestGroundY(pos rl.Vector3, halfWidth, eyeHeight float32) (groundY float32, ok bool) {
	feetY := pos.Y - eyeHeight

	samples := [4][2]float32{
		{pos.X - halfWidth, pos.Z - halfWidth},
		{pos.X + halfWidth, pos.Z - halfWidth},
		{pos.X - halfWidth, pos.Z + halfWidth},
		{pos.X + halfWidth, pos.Z + halfWidth},
	}

	found := false
	for _, s := range samples {
		// Prosleđujemo feetY direktno umesto maxAllowedY
		top, columnOK := columnGroundY(s[0], s[1], feetY)
		if !columnOK {
			continue
		}
		if !found || top > groundY {
			groundY = top
			found = true
		}
	}
	return groundY, found
}

func IsAirborne(pos rl.Vector3, eyeHeight, halfWidth float32) bool {
	groundY, ok := HighestGroundY(pos, halfWidth, eyeHeight)
	if !ok {
		return true
	}
	feetY := pos.Y - eyeHeight
	return feetY > groundY+GroundSnapEpsilon
}

func aabbOverlapsBlock(pos rl.Vector3, halfWidth, feetY, headY float32, wx, wy, wz int) bool {
	bMinX := float32(wx) - 0.5
	bMaxX := float32(wx) + 0.5
	bMinY := float32(wy) - 0.5
	bMaxY := float32(wy) + 0.5
	bMinZ := float32(wz) - 0.5
	bMaxZ := float32(wz) + 0.5

	pMinX := pos.X - halfWidth
	pMaxX := pos.X + halfWidth
	pMinZ := pos.Z - halfWidth
	pMaxZ := pos.Z + halfWidth

	return pMaxX > bMinX && pMinX < bMaxX &&
		headY > bMinY && feetY < bMaxY &&
		pMaxZ > bMinZ && pMinZ < bMaxZ
}

func resolveHorizontalAxis(pos *float32, halfWidth float32, blockCenter int) {
	blockMin := float32(blockCenter) - 0.5
	blockMax := float32(blockCenter) + 0.5
	playerMin := *pos - halfWidth
	playerMax := *pos + halfWidth

	overlapLeft := playerMax - blockMin
	overlapRight := blockMax - playerMin

	if overlapLeft > 0 && overlapRight > 0 {
		if overlapLeft < overlapRight {
			*pos -= overlapLeft
		} else {
			*pos += overlapRight
		}
	}
}

func ApplyHorizontalCollision(camera *rl.Camera3D, eyeHeight, halfWidth float32) {
	feetY := camera.Position.Y - eyeHeight
	headY := camera.Position.Y

	minWX := intFloor(camera.Position.X - halfWidth + 0.1)
	maxWX := intFloor(camera.Position.X + halfWidth - 0.1)
	minWZ := intFloor(camera.Position.Z - halfWidth + 0.1)
	maxWZ := intFloor(camera.Position.Z + halfWidth - 0.1)
	minLY := intFloor(feetY)
	maxLY := intFloor(headY)

	for wx := minWX; wx <= maxWX; wx++ {
		for wz := minWZ; wz <= maxWZ; wz++ {
			for ly := minLY; ly <= maxLY; ly++ {
				if world.GetGlobalBlock(wx, ly, wz) == blocks.Air || world.GetGlobalBlock(wx, ly, wz) == blocks.Water {
					continue
				}

				if !aabbOverlapsBlock(camera.Position, halfWidth, feetY+0.1, headY, wx, ly, wz) {
					continue
				}

				resolveHorizontalAxis(&camera.Position.X, halfWidth, wx)
				resolveHorizontalAxis(&camera.Target.X, halfWidth, wx)
				resolveHorizontalAxis(&camera.Position.Z, halfWidth, wz)
				resolveHorizontalAxis(&camera.Target.Z, halfWidth, wz)
			}
		}
	}
}

func ApplyVerticalBlockPhysics(camera *rl.Camera3D, velocity *float32, grounded *bool, eyeHeight float32) {
	halfWidth := float32(PlayerHalfWidth)

	if IsAirborne(camera.Position, eyeHeight, halfWidth) {
		*grounded = false
	}

	groundY, ok := HighestGroundY(camera.Position, halfWidth, eyeHeight)
	if !ok {
		return
	}

	feetY := camera.Position.Y - eyeHeight

	headWX := intFloor(camera.Position.X)
	headWZ := intFloor(camera.Position.Z)
	headWY := intFloor(camera.Position.Y)
	if *velocity > 0 && world.GetGlobalBlock(headWX, headWY, headWZ) != blocks.Air && world.GetGlobalBlock(headWX, headWY, headWZ) != blocks.Water {
		*velocity = 0
	}

	if *velocity <= 0 && feetY <= groundY+GroundSnapEpsilon {
		dy := groundY - feetY
		camera.Position.Y += dy
		camera.Target.Y += dy
		*velocity = 0
		*grounded = true
	}
}
