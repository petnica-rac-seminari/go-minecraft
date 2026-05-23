package reljef

import (
	"main/blocks"
	"main/world"

	"math/rand"

	"github.com/KEINOS/go-noise"
)

const chunkSize = 16
const viewChunks = 10

const width = chunkSize * viewChunks
const height = chunkSize * viewChunks

const smoothness = 15

func Octaves(n noise.Generator, v float64, x int, y int, br_o int) float64 {
	sm := 15
	value := v
	amplitude := 1.0
	frequency := 1.0
	maxAmp := 1.0

	for i := 0; i < br_o; i++ {
		nx := (float64(x) / float64(sm)) * frequency
		ny := (float64(y) / float64(sm)) * frequency

		value += n.Eval64(nx, ny) * amplitude

		maxAmp += amplitude

		amplitude *= 0.5
		frequency *= 2.0
	}

	return value / maxAmp
}

func assignValues(n noise.Generator, br_oct, startX, startZ int, sm_mod float64) []float64 {

	values := make([]float64, 16*16)

	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {
			worldX := startX + x
			worldZ := startZ + z

			v := n.Eval64(float64(worldX)/float64(sm_mod*smoothness), float64(worldZ)/float64(sm_mod*smoothness))
			v = Octaves(n, v, worldX, worldZ, br_oct)

			values[z*16+x] = v
		}
	}

	return values
}

func newNoise(worldSeed, seedMod, br_oct, startX, startZ int, sm_mod float64) []float64 {
	n, _ := noise.New(noise.OpenSimplex, int64(worldSeed+seedMod))
	return assignValues(n, br_oct, startX, startZ, sm_mod)
}

func DetermineTrees(startX, startZ, seed int) [][]float64 {
	coords := make([][]float64, 16)
	for x := 0; x < 16; x++ {
		coords[x] = make([]float64, 16)
	}

	values := newNoise(seed, 2, 2, startX, startZ, 1)

	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {

			v := values[z*16+x]
			norm := (v + 1.0) / 2.0

			v = values[z*16+x]
			norm1 := (v + 1.0) / 2.0

			combined := (norm + norm1) / 2.0
			normf := int(combined*combined*24) + 4

			if normf > 12 {
				coords[x][z] = 1
			}
		}
	}

	return coords
}

func BiomeMap(startX, startZ, seed int) [][]int {
	biomes := make([][]int, 16)

	for x := 0; x < 16; x++ {
		biomes[x] = make([]int, 16)
	}

	values := newNoise(seed, 2, 0, startX, startZ, 10)

	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {

			v := values[z*16+x]
			norm := v + 1.0

			normf := int(norm*norm*24) + 4

			if normf < 20 {
				biomes[x][z] = 1
			} else {
				biomes[x][z] = 2
			}
		}
	}

	return biomes
}

func GenerateTree(x, y, z int, chunk [][][]blocks.Block) {
	for dx := -2; dx <= 2; dx++ {
		for dy := 3; dy <= 5; dy++ {
			for dz := -2; dz <= 2; dz++ {
				nx, nz := x+dx, z+dz
				if (dx == -2 || dx == 2) && (dz == -2 || dz == 2) {
					continue
				}
				chunk[nx][y+dy][nz] = blocks.Leaves
			}
		}
	}

	for dx := -1; dx <= 1; dx++ {
		for dy := 6; dy <= 7; dy++ {
			for dz := -1; dz <= 1; dz++ {
				nx, nz := x+dx, z+dz
				if (dx == -1 || dx == 1) && (dz == -1 || dz == 1) {
					continue
				}
				chunk[nx][y+dy][nz] = blocks.Leaves
			}
		}
	}

	chunk[x][y][z] = blocks.Log
	chunk[x][y+1][z] = blocks.Log
	chunk[x][y+2][z] = blocks.Log
	chunk[x][y+3][z] = blocks.Log
	chunk[x][y+4][z] = blocks.Log
	chunk[x][y+5][z] = blocks.Log
}

func GenerateOW(startX, startZ, seed int) world.Chunk {
	chunk := make([][][]blocks.Block, 16)
	det_trees := DetermineTrees(startX, startZ, seed)
	biome_map := BiomeMap(startX, startZ, seed)
	gen_structs := make([][]bool, 16)
	for i := range gen_structs {
		gen_structs[i] = make([]bool, 16)
	}
	rng := rand.New(rand.NewSource(int64(seed)))
	for x := 0; x < 16; x++ {
		chunk[x] = make([][]blocks.Block, 64)
		for z := 0; z < 64; z++ {
			chunk[x][z] = make([]blocks.Block, 16)
		}
	}

	values := newNoise(seed, 0, 1, startX, startZ, 1)
	values1 := newNoise(seed, 1, 7, startX, startZ, 1)
	BiomeMap(startX, startZ, seed)

	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {
			v := values[z*16+x]
			norm := (v + 1.0) / 2.0

			v = values1[z*16+x]
			norm1 := (v + 1.0) / 2.0

			combined := (norm + norm1) / 2.0
			normf := int(combined*combined*24) + 4

			if normf > 7 {
				for y := 0; y <= int(normf); y++ {
					if y == int(normf) {
						if biome_map[x][z] == 1 {
							chunk[x][y][z] = blocks.Grass
						} else if biome_map[x][z] == 2 {
							chunk[x][y][z] = blocks.Sand
						}
					} else if y == int(normf)-1 {
						chunk[x][y][z] = blocks.Dirt
					} else if y == 0 {
						chunk[x][y][z] = blocks.Bedrock
					} else {
						chunk[x][y][z] = blocks.Stone
					}
				}

				if det_trees[x][z] == 1 {
					valid_tr := true

					for dx := -3; dx <= 3; dx++ {
						for dz := -3; dz <= 3; dz++ {
							nx := x + dx
							nz := z + dz

							if nx < 0 || nx >= 16 || nz < 0 || nz >= 16 {
								continue
							}

							if gen_structs[nx][nz] {
								valid_tr = false
								break
							}
						}

						if !valid_tr {
							break
						}
					}

					if valid_tr && biome_map[x][z] == 1 {
						gen_structs[x][z] = true
						if rng.Intn(2) > 0 && !(x < 2 || x > 13 || z < 2 || z > 13) {
							GenerateTree(x, normf+1, z, chunk)
						}
					}
				} else if biome_map[x][z] == 1 {
					if rng.Intn(128) > 126 && !(x < 2 || x > 13 || z < 2 || z > 13) {
						GenerateTree(x, normf+1, z, chunk)
					}
				}
			} else {
				for y := 0; y <= 8; y++ {
					if y == int(normf) {
						chunk[x][y][z] = blocks.Sand
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

	return world.Chunk{GlobalX: startX / 16, GlobalZ: startZ / 16, Blocks: chunk, Rendered: false, IsDirty: false}
}

func GenerateNether(startX, startZ, seed int) world.Chunk {
	chunk := make([][][]blocks.Block, 16)
	for x := 0; x < 16; x++ {
		chunk[x] = make([][]blocks.Block, 64)
		for z := 0; z < 64; z++ {
			chunk[x][z] = make([]blocks.Block, 16)
		}
	}

	values := newNoise(seed, 0, 1, startX, startZ, 0.25)
	values1 := newNoise(seed, 1, 7, startX, startZ, 0.25)
	BiomeMap(startX, startZ, seed)

	for x := 0; x < 16; x++ {
		for z := 0; z < 16; z++ {
			v := values[z*16+x]
			norm := (v + 1.0) / 2.0

			v = values1[z*16+x]
			norm1 := (v + 1.0) / 2.0

			combined := (norm + norm1) / 2.0
			normf := int(combined*combined*24) + 4

			for y := 0; y <= normf; y++ {
				if y == 0 {
					chunk[x][y][z] = blocks.Bedrock
				} else {
					chunk[x][y][z] = blocks.Netherrack
				}
			}
		}
	}

	return world.Chunk{GlobalX: startX / 16, GlobalZ: startZ / 16, Blocks: chunk, Rendered: false, IsDirty: false}
}
