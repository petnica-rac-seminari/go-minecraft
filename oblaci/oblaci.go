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
	for i := 0; i < 800; i++ {

		c := CLOUDS{
			x: float32(rand.Intn(480) - 240),
			y: 40,
			z: float32(rand.Intn(480) - 240),
		}

		clouds = append(clouds, c)
	}

	return clouds
}

func MoveClouds(clouds []CLOUDS, playerY float32) {

	for i := range clouds {

		clouds[i].x += 0.01 //brzina oblaka
		clouds[i].y = playerY + 40

		if clouds[i].x > 480 {
			clouds[i].x = -480
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
