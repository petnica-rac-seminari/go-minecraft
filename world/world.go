package world

import "main/blocks"

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
