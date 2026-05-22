package world

import (
	"main/blocks"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Chunk struct {
	GlobalX  int
	GlobalZ  int
	Blocks   [][][]blocks.Block
	Rendered bool
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

func RenderChunk(c Chunk, x_offset int, y_offset int) {
	xLen := len(c.Blocks)
	if xLen == 0 {
		return
	}
	yLen := len(c.Blocks[0])
	if yLen == 0 {
		return
	}
	zLen := len(c.Blocks[0][0])
	if zLen == 0 {
		return
	}

	for x, xBlock := range c.Blocks {
		for y, yBlock := range xBlock {
			for z, zBlock := range yBlock {
				// 2. Proveravamo da li je blok vidljiv (da li dodiruje Air)
				isVisible := false

				// Provera X ose (levo i desno)
				if x == 0 || x == xLen-1 || c.Blocks[x-1][y][z] == blocks.Air || c.Blocks[x+1][y][z] == blocks.Air {
					isVisible = true
				} else
				// Provera Y ose (dole i gore)
				if y == 0 || y == yLen-1 || c.Blocks[x][y-1][z] == blocks.Air || c.Blocks[x][y+1][z] == blocks.Air {
					isVisible = true
				} else
				// Provera Z ose (napred i nazad)
				if z == 0 || z == zLen-1 || c.Blocks[x][y][z-1] == blocks.Air || c.Blocks[x][y][z+1] == blocks.Air {
					isVisible = true
				}

				// 3. Ako je bar jedna strana izložena vazduhu (ili je na ivici chunka), renderujemo ga
				if isVisible {
					RenderBlock(zBlock, x+x_offset*16, y, z+y_offset*16)
				}
			}
		}

	}
}
