package voda

import (
	"main/blocks"
)

func PostaviIzvore(world *World) {
	for i := 0; i < len(world.Chunks); i++ {
		chunk := world.Chunks[i]
		najvisiY := -1
		for y := 64; y > 0; y-- {
			if chunk.Blocks[y] != blocks.Air {
				najvisiY = y
				break
			}
		}

		if najvisiY != -1 && najvisiY+1 < 64 {
			if chunk.Blocks[najvisiY+1] == blocks.Air {
				chunk.Blocks[najvisiY+1] = blocks.Water
			}
		}
	}
}
