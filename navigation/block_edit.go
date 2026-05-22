package navigation

import (
	"main/blocks"
	"main/world"
)

func GetBlock(chunk *world.Chunk, x, y, z int) (blocks.Block, bool) {
	if chunk == nil || !InBounds(x, y, z) {
		return blocks.Air, false
	}
	return chunk.Blocks[x][y][z], true
}

func DestroyBlock(chunk *world.Chunk, x, y, z int) bool {
	if chunk == nil || !InBounds(x, y, z) {
		return false
	}
	if chunk.Blocks[x][y][z] == blocks.Air {
		return false
	}
	chunk.Blocks[x][y][z] = blocks.Air
	return true
}

func PlaceBlock(chunk *world.Chunk, x, y, z int, b blocks.Block) bool {
	if chunk == nil || !InBounds(x, y, z) {
		return false
	}
	if chunk.Blocks[x][y][z] != blocks.Air {
		return false
	}
	if b == blocks.Air {
		return false
	}
	chunk.Blocks[x][y][z] = b
	return true
}

func PlaceAdjacent(chunk *world.Chunk, hit RaycastHit, b blocks.Block) bool {
	if chunk == nil || !hit.Hit {
		return false
	}
	return PlaceBlock(chunk, hit.X+hit.NX, hit.Y+hit.NY, hit.Z+hit.NZ, b)
}
