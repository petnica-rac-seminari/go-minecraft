package reljef

import (
	"main/blocks"
	"main/world"

	"github.com/KEINOS/go-noise"
)

const seed = 100

const chunkSize = 16
const viewChunks = 10

const width = chunkSize * viewChunks
const height = chunkSize * viewChunks

const smoothness = 15

func Octaves(n noise.Generator, v float64, x int, y int, br_o int) float64 {
	value := v
	amplitude := 1.0
	frequency := 1.0
	maxAmp := 1.0

	for i := 0; i < br_o; i++ {
		nx := (float64(x) / smoothness) * frequency
		ny := (float64(y) / smoothness) * frequency

		value += n.Eval64(nx, ny) * amplitude

		maxAmp += amplitude

		amplitude *= 0.5
		frequency *= 2.0
	}

	return value / maxAmp
}

func assignValues(n noise.Generator, br_oct int, startX, startY int) []float64 {

	values := make([]float64, 16*16)

	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {

			worldX := startX + x
			worldY := startY + y

			v := n.Eval64(float64(worldX)/smoothness, float64(worldY)/smoothness)
			v = Octaves(n, v, worldX, worldY, br_oct)

			values[y*16+x] = v
		}
	}

	return values
}

func newNoise(seedMod int, br_oct int, startX, startY int) []float64 {
	n, _ := noise.New(noise.OpenSimplex, int64(seed+seedMod))
	return assignValues(n, br_oct, startX, startY)
}

func GenerateChunk(startX, startZ int) world.Chunk {
	chunk := make([][][]blocks.Block, 16)
	for x := 0; x < 16; x++ {
		chunk[x] = make([][]blocks.Block, 64)
		for y := 0; y < 64; y++ {
			chunk[x][y] = make([]blocks.Block, 16)
		}
	}

	values := newNoise(0, 1, startX, startZ)
	values1 := newNoise(1, 7, startX, startZ)

	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {

			v := values[z*16+x]
			norm := (v + 1.0) / 2.0 * 16

			v = values1[z*16+x]
			norm1 := (v + 1.0) / 2.0 * 16

			normf := int(norm+norm1) / 2
			if normf > 6 {
				for y := 0; y <= int(normf); y++ {
					if y == int(normf) {
						chunk[x][y][z] = blocks.Grass
					} else if y == int(normf)-1 {
						chunk[x][y][z] = blocks.Dirt
					} else if y == 0 {
						chunk[x][y][z] = blocks.Bedrock
					} else {
						chunk[x][y][z] = blocks.Stone
					}
				}
			} else {
				for y := 0; y <= 6; y++ {
					if y == int(normf) {
						chunk[x][y][z] = blocks.Grass
					} else if y == int(normf)-1 {
						chunk[x][y][z] = blocks.Dirt
					} else if y == 0 {
						chunk[x][y][z] = blocks.Bedrock
					} else if y < int(normf)-1 {
						chunk[x][y][z] = blocks.Stone
					} else {
						chunk[x][y][z] = blocks.Water
					}
				}
			}
		}
	}

	return world.Chunk{GlobalX: startX / 16, GlobalZ: startZ / 16, Blocks: chunk, Rendered: false}
}
