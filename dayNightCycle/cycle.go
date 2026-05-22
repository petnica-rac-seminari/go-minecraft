package daynightcycle

import (
	"fmt"
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Color = rl.Color{R: 255, G: 255, B: 255, A: 255}

func SkyColor(currentTick float32) color.RGBA {
	currentTime := int(currentTick - float32(int(currentTick)/240*240))
	Cycle := (math.Cos((math.Pi * float64(currentTime%480) / 120.0)) + 1) * 255 / 2
	c := color.RGBA{
		R: uint8(Cycle * 120 / 255),
		G: uint8(Cycle * 223 / 255),
		B: uint8((Cycle + 25) * 255 / 280),
		A: 255,
	}

	fmt.Println(c)
	return c
}

func RenderSun() (model *rl.Model) {
	sunce_texture := rl.LoadTexture("sunce.png")
	sunce_mesh := rl.GenMeshSphere(2.0, 32, 32)
	sunce_model := rl.LoadModelFromMesh(sunce_mesh)
	rl.SetMaterialTexture(sunce_model.Materials, rl.MapDiffuse, sunce_texture)
	return &sunce_model
}

func RenderMoon() (model *rl.Model) {
	moonTexture := rl.LoadTexture("moon.png")
	moonMesh := rl.GenMeshSphere(1.5, 32, 32)
	moonModel := rl.LoadModelFromMesh(moonMesh)
	rl.SetMaterialTexture(moonModel.Materials, rl.MapDiffuse, moonTexture)

	return &moonModel
}

func SkyBodyAngle(currentTick float32) float32 {
	currentTime := (currentTick - float32(int(currentTick)/240*240))
	temp := ((currentTime / 240.0) * 360.0) * math.Pi / 180.0
	fmt.Println(temp)
	return temp
}

func MoveSun(skyBodyAngle float64, camera rl.Camera) rl.Vector3 {
	temp := rl.Vector3{
		X: (float32(math.Cos(skyBodyAngle-1))*50 + camera.Position.X),
		Y: float32(math.Sin(skyBodyAngle-1))*50 + 30,
		Z: camera.Position.Z,
	}

	fmt.Println(temp)
	return temp
}

func MoveMoon(skyBodyAngle float64, camera rl.Camera) rl.Vector3 {
	return rl.Vector3{
		X: -(float32(math.Cos(skyBodyAngle))*30 + camera.Position.X),
		Y: -(float32(math.Sin(skyBodyAngle))*30 + 30),
		Z: -camera.Position.Z,
	}
}
