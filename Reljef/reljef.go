package reljef

import (
	"main/blocks"
	"main/menu"
	"main/world"
	"math"

	"math/rand"

	"github.com/KEINOS/go-noise"
	rl "github.com/gen2brain/raylib-go/raylib"
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

func GeneratePortal(x, y, z int, chunk [][][]blocks.Block, dim bool) {
	for dx := -1; dx <= 2; dx++ {
		for dy := -3; dy <= 4; dy++ {
			nx := x + dx
			if (dx == 0 || dx == 1) && (dy > 0 && dy < 4) {
				continue
			} else if dy < 0 {
				if dim {
					chunk[nx][y+dy][z] = blocks.Grass
				} else {
					chunk[nx][y+dy][z] = blocks.Netherrack
				}
			} else if dy == 0 {
				chunk[nx][y+dy][z] = blocks.Portal
			} else {
				chunk[nx][y+dy][z] = blocks.Bedrock
			}
		}
	}
}

func treeChecker(x, z int, gen_structs [][]bool) bool {
	for dx := -3; dx <= 3; dx++ {
		for dz := -3; dz <= 3; dz++ {

			nx := x + dx
			nz := z + dz

			if nx < 0 || nx >= 16 || nz < 0 || nz >= 16 {
				continue
			}

			if gen_structs[nx][nz] {
				return false
			}
		}
	}

	return true
}

func GenerateRiverMap(chunkX, chunkZ, x, z int, frequency float64, seed int64) bool {
	n, _ := noise.New(noise.OpenSimplex, seed)

	riverMap := (n.Eval64(float64(x+16*chunkX)*frequency, float64(z+16*chunkZ)*frequency) + 0.866) * 0.577

	if riverMap >= 0.47 && riverMap <= 0.53 {
		return true
	} else {
		return false
	}
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
						switch biome_map[x][z] {
						case 1:
							chunk[x][y][z] = blocks.Grass
						case 2:
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
					valid_tr := treeChecker(x, z, gen_structs)

					if valid_tr && biome_map[x][z] == 1 {
						gen_structs[x][z] = true
						if rng.Intn(2) > 0 && !(x < 2 || x > 13 || z < 2 || z > 13) {
							GenerateTree(x, normf+1, z, chunk)
						}
					}
				} else if biome_map[x][z] == 1 {
					if rng.Float64() > 0.993 && !(x < 2 || x > 13 || z < 2 || z > 13) && treeChecker(x, z, gen_structs) {
						GeneratePortal(x, normf+1, z, chunk, true)
					} else if rng.Intn(64) > 62 && !(x < 2 || x > 13 || z < 2 || z > 13) {
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

			if GenerateRiverMap(startX, startZ, x, z, 0.02, int64(seed)) {

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

			if rand.Intn(100) > 95 {
				chunk[x][normf+1][z] = blocks.Sudomil_bot
				chunk[x][normf+2][z] = blocks.Sudomil_top
			}

		}
	}

	return world.Chunk{startX / 16, startZ / 16, chunk, false, nil, false}
}

func DimensionSwap(dim bool, halfDist int, playerCX, playerCZ int, camera rl.Camera3D) {
	world.LoadedChunks = make(map[world.ChunkPos]*world.Chunk)
	dim = !dim
	for z := -halfDist; z <= halfDist; z++ {
		for x := -halfDist; x <= halfDist; x++ {
			pos := world.ChunkPos{X: playerCX + x, Z: playerCZ + z}

			if _, exists := world.LoadedChunks[pos]; !exists {
				if dim {
					c := GenerateOW(pos.X*16, pos.Z*16, menu.Seed)
					world.LoadedChunks[pos] = &c
				} else {
					c := GenerateNether(pos.X*16, pos.Z*16, menu.Seed)
					world.LoadedChunks[pos] = &c
				}
			}
		}
	}

	pos := world.ChunkPos{X: playerCX, Z: playerCZ}
	playerX := (int(math.Abs(float64(camera.Position.X))) % 16)
	playerZ := (int(math.Abs(float64(camera.Position.Z))) % 16)
	playerY := int(camera.Position.Y - 1)
	camera.Position.Y += 5

	if playerX > 13 {
		if playerZ > 13 {
			GeneratePortal(playerX-4, playerY-2, playerZ-4, world.LoadedChunks[pos].Blocks, dim)
		} else if playerZ < 3 {
			GeneratePortal(playerX-4, playerY-2, playerZ+4, world.LoadedChunks[pos].Blocks, dim)
		} else {
			GeneratePortal(playerX-4, playerY-2, playerZ, world.LoadedChunks[pos].Blocks, dim)
		}
	} else if playerX < 3 {
		if playerZ > 13 {
			GeneratePortal(playerX+4, playerY-2, playerZ-4, world.LoadedChunks[pos].Blocks, dim)
		} else if playerZ < 3 {
			GeneratePortal(playerX+4, playerY-2, playerZ+4, world.LoadedChunks[pos].Blocks, dim)
		} else {
			GeneratePortal(playerX+4, playerY-2, playerZ, world.LoadedChunks[pos].Blocks, dim)
		}
	} else {
		GeneratePortal(playerX, playerY-2, playerZ, world.LoadedChunks[pos].Blocks, dim)
	}

	camera.Position.Y += 5
}
