package navigation

import (
	"main/blocks"
	"main/world"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var SelectedBlock blocks.Block = blocks.Grass

func getScreenCenter() rl.Vector2 {
	return rl.NewVector2(
		float32(rl.GetScreenWidth())/2.0,
		float32(rl.GetScreenHeight())/2.0,
	)
}

func selectBlock() {
	// Loop through keys 1 to 9 (rl.KeyOne is 49, rl.KeyNine is 57)
	for i := 0; i < 9; i++ { // Fixed loop bounds to cleanly match keys 1-9
		targetKey := int32(rl.KeyOne) + int32(i)

		if rl.IsKeyPressed(targetKey) {
			SelectedBlock = blocks.Block(i + 1)
		}
	}
}

func HandleBlockManipulation(camera *rl.Camera3D) {
	selectBlock()

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) || rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		pointRay := rl.GetScreenToWorldRay(getScreenCenter(), *camera)
		maxRange := float32(12)
		step := float32(0.1) // Shrunk slightly to make edge-detection significantly less buggy

		pos := pointRay.Position
		prevPos := pos

		for i := float32(0); i < maxRange; i += step {
			targetX, targetY, targetZ := NJK(pos.X), NJK(pos.Y), NJK(pos.Z)
			blok := world.GetGlobalBlock(targetX, targetY, targetZ)

			if blok >= 2 {
				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					if blok != blocks.Bedrock {
						world.SetGlobalBlock(targetX, targetY, targetZ, blocks.Air)
						updateNeighboringWater(targetX, targetY, targetZ)
					}
				} else {
					placeX, placeY, placeZ := NJK(prevPos.X), NJK(prevPos.Y), NJK(prevPos.Z)

					// Make sure we actually moved to a different voxel space coordinate before placing
					if placeX != targetX || placeY != targetY || placeZ != targetZ {

						// --- BULLETPROOF PLAYER COLLISION CHECK ---
						playerBox := GetPlayerBoundingBox(camera.Position)

						// Convert the prospective block index into physical 3D spaces bounds
						blockMinX := float32(placeX) - 0.5
						blockMaxX := float32(placeX) + 0.5
						blockMinY := float32(placeY) - 0.5
						blockMaxY := float32(placeY) + 0.5
						blockMinZ := float32(placeZ) - 0.5
						blockMaxZ := float32(placeZ) + 0.5

						// Check if the 3D block bounds overlap with your actual player body
						collidesWithPlayer :=
							(blockMaxX > playerBox.Min.X && blockMinX < playerBox.Max.X) &&
								(blockMaxY > playerBox.Min.Y && blockMinY < playerBox.Max.Y) &&
								(blockMaxZ > playerBox.Min.Z && blockMinZ < playerBox.Max.Z)

						// Only place the block if it's clear of your bounding box!
						if !collidesWithPlayer {
							world.SetGlobalBlock(placeX, placeY, placeZ, SelectedBlock)
							spreadFromNeighbors(placeX, placeY, placeZ)
						}
					}
				}
				return
			}
			prevPos = pos
			pos.X += pointRay.Direction.X * step
			pos.Y += pointRay.Direction.Y * step
			pos.Z += pointRay.Direction.Z * step
		}
	}
}

func NJK(f float32) int {
	return int(math.Round(float64(f)))
}

func spreadFromNeighbors(x, y, z int) {
	if world.GetGlobalBlock(x, y, z) != blocks.Water {
		return
	}

	blocksToCheck := [][3]int{{x, y - 1, z}, {x - 1, y, z}, {x + 1, y, z}, {x, y, z - 1}, {x, y, z + 1}}
	time.Sleep(time.Millisecond * 500)
	for _, d := range blocksToCheck {
		nx, ny, nz := d[0], d[1], d[2]
		if ny == y-1 && world.GetGlobalBlock(nx, ny, nz) <= 2 {
			world.SetGlobalBlock(d[0], d[1], d[2], blocks.Water)
			go spreadFromNeighbors(nx, ny, nz)
			return
		} else if world.GetGlobalBlock(nx, ny, nz) == blocks.Air {
			world.SetGlobalBlock(nx, ny, nz, blocks.Water)
			go spreadFromNeighbors(nx, ny, nz)
		}
	}
}

func updateNeighboringWater(x, y, z int) {
	blocksToCheck := [][3]int{{x, y - 1, z}, {x - 1, y, z}, {x + 1, y, z}, {x, y, z - 1}, {x, y, z + 1}, {x, y + 1, z}}
	for _, d := range blocksToCheck {
		nx, ny, nz := d[0], d[1], d[2]
		if world.GetGlobalBlock(nx, ny, nz) == blocks.Water {
			go spreadFromNeighbors(nx, ny, nz)
		}
	}
}

func Vector3Round(v rl.Vector3) rl.Vector3 {
	return rl.NewVector3(
		float32(math.Round(float64(v.X))),
		float32(math.Round(float64(v.Y))),
		float32(math.Round(float64(v.Z))),
	)
}
