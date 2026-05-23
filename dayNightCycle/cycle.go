package dayNightCycle

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Color = rl.Color{R: 255, G: 255, B: 255, A: 255}

func SkyColor(currentTick int, dimension bool) color.RGBA {
	Cycle := (math.Cos((math.Pi * float64(currentTick%480) / 120.0)) + 1) * 255 / 2
	var c color.RGBA
	if dimension {
		c := color.RGBA{
			R: uint8(Cycle * 120 / 255),
			G: uint8(Cycle * 223 / 255),
			B: uint8((Cycle + 25) * 255 / 280),
			A: 255,
		}
		return c
	} else {
		return color.RGBA{150, 0, 0, 255}
	}

	return c
}

func RenderSun() (model *rl.Model) {
	sunce_texture := rl.LoadTexture("sunce.png")
	sunce_mesh := rl.GenMeshSphere(2.0, 32, 32)
	sunce_model := rl.LoadModelFromMesh(sunce_mesh)
	rl.SetMaterialTexture(sunce_model.Materials, rl.MapDiffuse, sunce_texture)
	return &sunce_model
}

func SkyBodyAngle(currentTick float32) float32 {
	currentTime := (currentTick - float32(int(currentTick)/240*240))
	return ((currentTime / 240.0) * 360.0) * math.Pi / 180.0
}

func MoveSun(skyBodyAngle float64, camera rl.Camera) rl.Vector3 {
	v := rl.Vector3{
		X: float32(math.Cos(skyBodyAngle))*150 + camera.Position.X,
		Y: float32(math.Sin(skyBodyAngle))*150 + 30 + camera.Position.Y,
		Z: camera.Position.Z,
	}
	return v
}
