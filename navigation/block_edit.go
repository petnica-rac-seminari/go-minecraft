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
