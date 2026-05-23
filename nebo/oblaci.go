package nebo

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
	for i := 0; i < 100; i++ {

		c := CLOUDS{
			x: float32(rand.Intn(80) - 40),
			y: 20,
			z: float32(rand.Intn(80) - 40),
		}

		clouds = append(clouds, c)
	}

	return clouds
}

func MoveClouds(clouds []CLOUDS) {

	for i := range clouds {

		clouds[i].x += 0.02 //brzina oblaka

		if clouds[i].x > 40 {
			clouds[i].x = -40
		}
	}
}

func RenderCloud(c CLOUDS, playerPos rl.Vector3) {

	transparentWhite := rl.Fade(rl.White, 0.8)

	rl.DrawCube(
		rl.NewVector3(c.x+playerPos.X, c.y, c.z+playerPos.Z),
		3.0,
		1.2,
		3.0,
		transparentWhite,
	)

	rl.DrawCube(
		rl.NewVector3(c.x-2.0+playerPos.X, c.y+0.5, c.z+playerPos.Z),
		2.0,
		1.0,
		2.0,
		transparentWhite,
	)

	rl.DrawCube(
		rl.NewVector3(c.x+2.0+playerPos.X, c.y+0.5, c.z+playerPos.Z),
		2.0,
		1.0,
		2.0,
		transparentWhite,
	)
}

func DrawClouds(clouds []CLOUDS, playerPos rl.Vector3) {

	for _, c := range clouds {
		RenderCloud(c, playerPos)
	}
}
