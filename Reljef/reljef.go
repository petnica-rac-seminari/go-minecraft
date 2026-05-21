package reljef

import (
	"main/blocks"
	"main/world"
	"math/rand"

	"github.com/KEINOS/go-noise"
)

func BlockAtLocation(x, y, z int, frequency, amplitude, baseHeight float64, seed int64) blocks.Block {
	n, _ := noise.New(noise.OpenSimplex, seed)
	amplitude /= 2
	height := (n.Eval64(float64(x)*frequency, float64(z)*frequency)+1)*amplitude + baseHeight
	if height < float64(y) {
		return blocks.Air
	} else {
		if rand.Float64() < 0.5 {
			return blocks.Stone
		} else {
			return blocks.Grass
		}
	}
}

func GenerateChunk(chunkX, chunkZ int, frequency, amplitude, baseHeight float64, seed int64) world.Chunk {
	//frequency = 0.1
	//amplitude = 4
	//baseHeight = 4
	chunk := make([][][]blocks.Block, 16)
	for i := 0; i < 16; i++ {
		chunk[i] = make([][]blocks.Block, 64)
		for j := 0; j < 64; j++ {
			chunk[i][j] = make([]blocks.Block, 16)
			for k := 0; k < 16; k++ {
				block := BlockAtLocation(i+16*chunkX, j, k+16*chunkZ, frequency, amplitude, baseHeight, seed)
				chunk[i][j][k] = block
			}
		}
	}

	return world.Chunk{GlobalX: chunkX, GlobalZ: chunkZ, Blocks: chunk}
}
