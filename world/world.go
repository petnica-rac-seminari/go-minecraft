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
	color := rl.Gray
	switch block {
	case blocks.Air:
		return
	case blocks.Grass:
		color = rl.Green
	case blocks.Stone:
		color = rl.DarkGray
	case blocks.Dirt:
		color = rl.DarkBrown
	case blocks.Water:
		color = rl.Blue
	case blocks.Bedrock:
		color = rl.Black
	case blocks.Snow:
		color = rl.White
	}
	rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 1.0, 1.0, 1.0, color)
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
