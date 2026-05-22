package oblaci

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CLOUDS struct {
	x float32
	y float32
	z float32
}

func GenerateClouds() []CLOUDS {

	var clouds []CLOUDS
	for i := 0; i < 30; i++ {
		c := CLOUDS{
			x: float32(rand.Intn(80) - 40),
			y: 20, // konstantna visina
			z: float32(rand.Intn(80) - 40),
		}
		clouds = append(clouds, c)
	}

	return clouds
}

func RenderCloud(clouds []CLOUDS) {
	transparentWhite := rl.Fade(rl.White, 0.8)

	for _, c := range clouds {
		rl.DrawCube(rl.NewVector3(c.x, c.y, c.z), 3.0, 1.2, 3.0, transparentWhite)
		rl.DrawCube(rl.NewVector3(c.x-2.0, c.y+0.5, c.z), 2.0, 1.0, 2.0, transparentWhite)
		rl.DrawCube(rl.NewVector3(c.x+2.0, c.y+0.5, c.z), 2.0, 1.0, 2.0, transparentWhite)
	}
}
