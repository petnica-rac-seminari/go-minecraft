package world

import (
	"main/blocks"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Chunk struct {
	GlobalX int
	GlobalZ int
	Blocks  [][][]blocks.Block
}

type WorldStruct struct {
	//singleton
	Chunks [3][3]Chunk
}

var WorldInstance *WorldStruct

func World() *WorldStruct {
	if WorldInstance == nil {
		WorldInstance = &WorldStruct{}
		return WorldInstance
	} else {
		return WorldInstance
	}
}

/*func NewChunk(x, y int) *Chunk {
	new_chunk := &Chunk{globalX: x, globalY: y}
	World().Chunks[x][y] = *new_chunk
	return new_chunk
}*/

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
	for i, xBlock := range c.Blocks {
		for j, yBlock := range xBlock {
			for k, zBlock := range yBlock {
				RenderBlock(zBlock, i, j, k)
			}
		}

	}
}
