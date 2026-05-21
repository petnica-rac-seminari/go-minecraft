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
	return b != blocks.Air
}

func columnGroundY(chunk world.Chunk, worldX, worldZ float32) (groundY float32, ok bool) {
	lx, _, lz, inChunk := WorldToLocal(chunk, intFloor(worldX), 0, intFloor(worldZ))
	if !inChunk {
		return 0, false
	}

	for ly := ChunkSizeY - 1; ly >= 0; ly-- {
		if isSolid(chunk.Blocks[lx][ly][lz]) {
			return blockTopY(ly), true
		}
	}
	return 0, false
}

func HighestGroundY(chunk world.Chunk, pos rl.Vector3, halfWidth, eyeHeight float32) (groundY float32, ok bool) {
	_ = eyeHeight
	samples := [4][2]float32{
		{pos.X - halfWidth, pos.Z - halfWidth},
		{pos.X + halfWidth, pos.Z - halfWidth},
		{pos.X - halfWidth, pos.Z + halfWidth},
		{pos.X + halfWidth, pos.Z + halfWidth},
	}

	found := false
	for _, s := range samples {
		top, columnOK := columnGroundY(chunk, s[0], s[1])
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

func IsAirborne(chunk world.Chunk, pos rl.Vector3, eyeHeight, halfWidth float32) bool {
	groundY, ok := HighestGroundY(chunk, pos, halfWidth, eyeHeight)
	if !ok {
		return true
	}
	feetY := pos.Y - eyeHeight
	return feetY > groundY+GroundSnapEpsilon
}

func aabbOverlapsBlock(pos rl.Vector3, halfWidth, feetY, headY float32, lx, ly, lz int) bool {
	bMinX := float32(lx) - 0.5
	bMaxX := float32(lx) + 0.5
	bMinY := float32(ly) - 0.5
	bMaxY := float32(ly) + 0.5
	bMinZ := float32(lz) - 0.5
	bMaxZ := float32(lz) + 0.5

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

func ApplyHorizontalCollision(camera *rl.Camera3D, chunk world.Chunk, eyeHeight, halfWidth float32) {
	feetY := camera.Position.Y - eyeHeight
	headY := camera.Position.Y

	minWX := intFloor(camera.Position.X - halfWidth)
	maxWX := intFloor(camera.Position.X + halfWidth)
	minWZ := intFloor(camera.Position.Z - halfWidth)
	maxWZ := intFloor(camera.Position.Z + halfWidth)
	minLY := intFloor(feetY)
	maxLY := intFloor(headY)

	for wx := minWX; wx <= maxWX; wx++ {
		for wz := minWZ; wz <= maxWZ; wz++ {
			lx, _, lz, inChunk := WorldToLocal(chunk, wx, 0, wz)
			if !inChunk {
				continue
			}
			for ly := minLY; ly <= maxLY; ly++ {
				if ly < 0 || ly >= ChunkSizeY {
					continue
				}
				if !InBounds(lx, ly, lz) {
					continue
				}
				if !isSolid(chunk.Blocks[lx][ly][lz]) {
					continue
				}
				if !aabbOverlapsBlock(camera.Position, halfWidth, feetY, headY, wx, ly, wz) {
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

func headHitsSolid(chunk world.Chunk, pos rl.Vector3, halfWidth float32) bool {
	headY := pos.Y
	minWX := intFloor(pos.X - halfWidth)
	maxWX := intFloor(pos.X + halfWidth)
	minWZ := intFloor(pos.Z - halfWidth)
	maxWZ := intFloor(pos.Z + halfWidth)
	ly := intFloor(headY)

	for wx := minWX; wx <= maxWX; wx++ {
		for wz := minWZ; wz <= maxWZ; wz++ {
			lx, _, lz, inChunk := WorldToLocal(chunk, wx, ly, wz)
			if !inChunk || ly < 0 || ly >= ChunkSizeY {
				continue
			}
			if isSolid(chunk.Blocks[lx][ly][lz]) {
				return true
			}
		}
	}
	return false
}

func ApplyVerticalBlockPhysics(camera *rl.Camera3D, velocity *float32, grounded *bool, chunk world.Chunk, eyeHeight float32) {
	halfWidth := float32(PlayerHalfWidth)

	if IsAirborne(chunk, camera.Position, eyeHeight, halfWidth) {
		*grounded = false
	}

	groundY, ok := HighestGroundY(chunk, camera.Position, halfWidth, eyeHeight)
	if !ok {
		return
	}

	feetY := camera.Position.Y - eyeHeight

	if *velocity > 0 && headHitsSolid(chunk, camera.Position, halfWidth) {
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
