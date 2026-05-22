package navigation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawBlockOutline(x, y, z int, color rl.Color) {
	rl.DrawCubeWires(
		rl.NewVector3(float32(x), float32(y), float32(z)),
		1.0, 1.0, 1.0,
		color,
	)
}
