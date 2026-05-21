package reljef

import (
	"main/blocks"
	"main/world"

	"github.com/KEINOS/go-noise"
)

func BlockAtLocation(x, y, z int, frequency, amplitude, baseHeight float64, seed int64) blocks.Block {
	n, _ := noise.New(noise.OpenSimplex, seed)
	amplitude /= 2
	height := (n.Eval64(float64(x)*frequency, float64(z)*frequency)+1)*amplitude + baseHeight
	if height < float64(y) {
		return blocks.Air
	} else {
		return blocks.Stone
	}
}

func GenerateChunk(chunkX, chunkZ int, frequency, amplitude, baseHeight float64, seed int64) world.Chunk {
	chunk := make([][][]blocks.Block, 16)
	for i := 0; i < 16; i++ {
		chunk[i] = make([][]blocks.Block, 64)
		for k := 0; k < 64; k++ {
			chunk[i][k] = make([]blocks.Block, 16)
			for j := 0; j < 16; j++ {
				block := BlockAtLocation(i+16*chunkX, k, j+16*chunkZ, frequency, amplitude, baseHeight, seed)
				chunk[i][k][j] = block
			}
		}
	}

	return world.Chunk{GlobalX: chunkX, GlobalZ: chunkZ, Blocks: chunk}
}

func Gen3x3(chunkX, chunkZ int, frequency, amplitude, baseHeight float64, seed int64) [][]world.Chunk {
	chunks := make([][]world.Chunk, 3)
	for i := -1; i <= 1; i++ {
		chunks[i] = make([]world.Chunk, 3)
		for j := -1; j <= 1; j++ {
			chunks[i][j] = GenerateChunk(chunkX, chunkZ, frequency, amplitude, baseHeight, seed)
		}
	}

	return chunks
}
