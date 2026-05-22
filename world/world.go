package world

import (
	"main/blocks"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Chunk struct {
	GlobalX  int
	GlobalZ  int
	Blocks   [][][]blocks.Block
	Rendered bool
}

type ChunkPos struct {
	X, Z int
}

var LoadedChunks = make(map[ChunkPos]*Chunk)
var LoadedStructures = make(map[ChunkPos]*Chunk)

var BlockModelRegistryInstance *BlockRegistry

type BlockRegistry struct {
	BlockModels map[blocks.Block]rl.Model
}

func (br *BlockRegistry) RegisterNewModel(block blocks.Block, cubeMesh rl.Mesh, path string) {
	new_model := rl.LoadModelFromMesh(cubeMesh)
	rl.SetMaterialTexture(new_model.Materials, rl.MapDiffuse, rl.LoadTexture(path))
	br.BlockModels[block] = new_model
}

func BlockModelRegistry() *BlockRegistry {
	if BlockModelRegistryInstance == nil {
		cubeMesh := rl.GenMeshCube(1.0, 1.0, 1.0)
		BlockModelRegistryInstance = &BlockRegistry{BlockModels: make(map[blocks.Block]rl.Model)}
		BlockModelRegistryInstance.RegisterNewModel(blocks.Grass, cubeMesh, "assets/grass.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Dirt, cubeMesh, "assets/dirt.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Water, cubeMesh, "assets/water.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Stone, cubeMesh, "assets/stone.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Bedrock, cubeMesh, "assets/bedrock.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Log, cubeMesh, "assets/log.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Leaves, cubeMesh, "assets/leaves.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Netherrack, cubeMesh, "assets/netherrack.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Sand, cubeMesh, "assets/sand.png")
	}
	return BlockModelRegistryInstance
}

func GetGlobalBlock(worldX, worldY, worldZ int) blocks.Block {
	if worldY < 0 || worldY >= 64 {
		return blocks.Air
	}
	cx := int(math.Floor(float64(worldX) / 16.0))
	cz := int(math.Floor(float64(worldZ) / 16.0))

	if chunk, exists := LoadedChunks[ChunkPos{cx, cz}]; exists {
		lx := worldX - cx*16
		lz := worldZ - cz*16
		if lx >= 0 && lx < 16 && lz >= 0 && lz < 16 {
			return chunk.Blocks[lx][worldY][lz]
		}
	}
	return blocks.Air
}

func SetGlobalBlock(worldX, worldY, worldZ int, b blocks.Block) bool {
	if worldY < 0 || worldY >= 64 {
		return false
	}
	cx := int(math.Floor(float64(worldX) / 16.0))
	cz := int(math.Floor(float64(worldZ) / 16.0))

	if chunk, exists := LoadedChunks[ChunkPos{cx, cz}]; exists {
		lx := worldX - cx*16
		lz := worldZ - cz*16
		if lx >= 0 && lx < 16 && lz >= 0 && lz < 16 {
			chunk.Blocks[lx][worldY][lz] = b
			return true
		}
	}
	return false
}

func RenderBlock(block blocks.Block, x, y, z int) {
	RegistryInstance := BlockModelRegistry()
	val, ok := RegistryInstance.BlockModels[block]

	if !ok {
		return
	}

	rl.DrawModel(val, rl.NewVector3(float32(x), float32(y), float32(z)), 1.0, rl.White)
}

func RenderChunk(c Chunk) {
	xLen := len(c.Blocks)
	if xLen == 0 {
		return
	}
	yLen := len(c.Blocks[0])

	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			for z := 0; z < 16; z++ {
				if c.Blocks[x][y][z] == blocks.Air {
					continue
				}

				wx := x + c.GlobalX*16
				wz := z + c.GlobalZ*16

				isVisible := false
				if GetGlobalBlock(wx-1, y, wz) == blocks.Air || GetGlobalBlock(wx+1, y, wz) == blocks.Air ||
					GetGlobalBlock(wx, y-1, wz) == blocks.Air || GetGlobalBlock(wx, y+1, wz) == blocks.Air ||
					GetGlobalBlock(wx, y, wz-1) == blocks.Air || GetGlobalBlock(wx, y, wz+1) == blocks.Air {
					isVisible = true
				}

				if isVisible {
					RenderBlock(c.Blocks[x][y][z], wx, y, wz)
				}
			}
		}
	}
}
