package daynightcycle

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	dayDuration = 240
	tick30sec   = 30
	tick60sec   = 60
	tick120sec  = 120
)

func SkyColor(currentTick int64) color.RGBA {
	day := currentTick % dayDuration
	dayColor := color.RGBA{R: 140, G: 184, B: 255, A: 255}
	sunsetColor := color.RGBA{R: 142, G: 77, B: 37, A: 255}
	nightColor := color.RGBA{R: 20, G: 36, B: 69, A: 255}

	var hR, hG, hB float64

	if day >= 0 && day <= tick30sec {
		hR, hG, hB = float64(dayColor.R), float64(dayColor.G), float64(dayColor.B)
	} else if day <= tick60sec {
		t := float64(day-tick30sec) / float64(tick30sec)
		hR = (1.0-t)*float64(dayColor.R) + t*float64(sunsetColor.R)
		hG = (1.0-t)*float64(dayColor.G) + t*float64(sunsetColor.G)
		hB = (1.0-t)*float64(dayColor.B) + t*float64(sunsetColor.B)
	} else if day <= tick120sec {
		t := float64(day-tick60sec) / float64(tick60sec)
		hR = (1.0-t)*float64(sunsetColor.R) + t*float64(nightColor.R)
		hG = (1.0-t)*float64(sunsetColor.G) + t*float64(nightColor.G)
		hB = (1.0-t)*float64(sunsetColor.B) + t*float64(nightColor.B)
	} else if day <= tick120sec+tick30sec {
		t := float64(day-tick120sec) / float64(tick30sec)
		hR = (1.0-t)*float64(nightColor.R) + t*float64(sunsetColor.R)
		hG = (1.0-t)*float64(nightColor.G) + t*float64(sunsetColor.G)
		hB = (1.0-t)*float64(nightColor.B) + t*float64(sunsetColor.B)
	} else {
		t := float64(day-(tick120sec+tick30sec)) / float64(dayDuration-(tick120sec+tick30sec))
		hR = (1.0-t)*float64(sunsetColor.R) + t*float64(dayColor.R)
		hG = (1.0-t)*float64(sunsetColor.G) + t*float64(dayColor.G)
		hB = (1.0-t)*float64(sunsetColor.B) + t*float64(dayColor.B)
	}

	return color.RGBA{uint8(clamp(hR, 0, 255)), uint8(clamp(hG, 0, 255)), uint8(clamp(hB, 0, 255)), 255}
}

func clamp(v float64, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
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
	return temp
}

func MoveSun(skyBodyAngle float64, camera rl.Camera) rl.Vector3 {
	return rl.Vector3{
		X: (float32(math.Cos(skyBodyAngle-1))*50 + camera.Position.X),
		Y: float32(math.Sin(skyBodyAngle-1))*50 + 30,
		Z: camera.Position.Z,
	}
}

func MoveMoon(skyBodyAngle float64, camera rl.Camera) rl.Vector3 {
	return rl.Vector3{
		X: -(float32(math.Cos(skyBodyAngle))*30 + camera.Position.X),
		Y: -(float32(math.Sin(skyBodyAngle))*30 + 30),
		Z: -camera.Position.Z,
	}
}
