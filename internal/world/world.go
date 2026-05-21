package world

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const BlockSize = 1.0

type Pos struct{ X, Y, Z int }

type World struct {
	blocks map[Pos]rl.Color
}

func New() *World {
	return &World{blocks: make(map[Pos]rl.Color)}
}

func NewDefault() *World {
	w := New()
	w.Set(0, 0, 0, rl.Blue)
	w.Set(1, 0, 0, rl.Green)
	w.Set(-1, 0, 0, rl.Red)
	return w
}

func (w *World) Has(x, y, z int) bool {
	_, ok := w.blocks[Pos{x, y, z}]
	return ok
}

func (w *World) GetColor(x, y, z int) (rl.Color, bool) {
	c, ok := w.blocks[Pos{x, y, z}]
	return c, ok
}

func (w *World) Set(x, y, z int, color rl.Color) {
	w.blocks[Pos{x, y, z}] = color
}

func (w *World) Remove(x, y, z int) {
	delete(w.blocks, Pos{x, y, z})
}

func (w *World) Count() int {
	return len(w.blocks)
}

func (w *World) BlockCenter(x, y, z int) rl.Vector3 {
	return rl.NewVector3(
		float32(x)+0.5,
		float32(y)+0.5,
		float32(z)+0.5,
	)
}

func (w *World) Draw() {
	for pos, col := range w.blocks {
		c := w.BlockCenter(pos.X, pos.Y, pos.Z)
		rl.DrawCube(c, BlockSize, BlockSize, BlockSize, col)
		rl.DrawCubeWires(c, BlockSize, BlockSize, BlockSize, rl.Black)
	}
}
