package blocks

import rl "github.com/gen2brain/raylib-go/raylib"

type Block int

const (
	Air Block = iota
	Bedrock
	Water
	Grass
	Dirt
	Stone
	Snow

// Log
// Leaves
)

func RenderBlock(block Block, x, y, z int) {
	switch block {
	case Air:
		return
	case Grass:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.DarkGreen)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Green)
	case Stone:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Gray)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.DarkGray)
	case Dirt:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Brown)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.DarkBrown)
	case Water:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Blue)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.DarkBlue)
	case Bedrock:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Black)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Gray)
	case Snow:
		rl.DrawCube(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.White)
		rl.DrawCubeWires(rl.NewVector3(float32(x), float32(y), float32(z)), 2.0, 2.0, 2.0, rl.Gray)
	}
}
