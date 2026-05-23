package nebo

import (
	"image/color"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	dayDuration = 240
	tick30sec   = 30
	tick60sec   = 60
	tick120sec  = 120
)

type STAR struct {
	Position rl.Vector3
	Size     float32
}

var Stars []STAR

func GenerateStars(count int) {

	Stars = nil

	for i := 0; i < count; i++ {

		angle := rand.Float64() * math.Pi * 2
		height := rand.Float64() * math.Pi

		radius := float32(80)

		x := float32(math.Cos(angle)*math.Sin(height)) * radius
		y := float32(math.Cos(height))*radius + 30
		z := float32(math.Sin(angle)*math.Sin(height)) * radius

		Stars = append(Stars, STAR{
			Position: rl.Vector3{
				X: x,
				Y: y,
				Z: z,
			},
			Size: rand.Float32()*0.15 + 0.05,
		})
	}
}

func DrawStars(camera rl.Camera) {

	for _, star := range Stars {

		worldPos := rl.Vector3{
			X: star.Position.X + camera.Position.X,
			Y: star.Position.Y,
			Z: star.Position.Z + camera.Position.Z,
		}

		rl.DrawSphere(
			worldPos,
			star.Size,
			color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			},
		)
	}
}

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

func IsNight(currentTick int64) bool {
	day := currentTick % dayDuration
	return day >= tick60sec && day < tick120sec+tick30sec
}

func IsDay(currentTick int64) bool {
	return !IsNight(currentTick)
}

func IsSunVisible(currentTick int64) bool {
	day := currentTick % dayDuration
	return day >= tick120sec || day < tick60sec
}

func IsMoonVisible(currentTick int64) bool {
	day := currentTick % dayDuration
	return day >= tick60sec && day < tick120sec+tick30sec
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
	return temp + math.Pi/2 + 1.0
}

func MoveSun(skyBodyAngle int64, camera rl.Camera) rl.Vector3 {
	angle := float64(skyBodyAngle-1) + math.Pi/2
	return rl.Vector3{
		X: float32(math.Cos(angle))*30 + camera.Position.X,
		Y: float32(math.Sin(angle))*30 + 30,
		Z: camera.Position.Z,
	}
}

func MoveSunByAngle(skyBodyAngle float32, camera rl.Camera) rl.Vector3 {
	angle := float64(skyBodyAngle - 1)
	return rl.Vector3{
		X: float32(math.Cos(angle))*50 + camera.Position.X,
		Y: float32(math.Sin(angle))*50 + 30,
		Z: camera.Position.Z,
	}
}
func MoveMoonByAngle(skyBodyAngle float32, camera rl.Camera) rl.Vector3 {
	angle := float64(skyBodyAngle-1) + math.Pi
	return rl.Vector3{
		X: float32(math.Cos(angle))*30 + camera.Position.X,
		Y: float32(math.Sin(angle))*20 + 10,
		Z: camera.Position.Z,
	}
}

func MoveMoon(currentTick int64, camera rl.Camera) rl.Vector3 {
	day := currentTick % dayDuration
	phase := float64(day-tick60sec) / float64(tick120sec+tick30sec-tick60sec)
	angle := phase * math.Pi
	return rl.Vector3{
		X: float32(math.Cos(angle))*30 + camera.Position.X,
		Y: float32(math.Sin(angle))*30 + 30,
		Z: camera.Position.Z,
	}
}
