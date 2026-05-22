package navigation

import (
	//"main/blocks"
	"main/world"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetScreenCenter() rl.Vector2 {
	return rl.NewVector2(
		float32(rl.GetScreenWidth())/2.0,
		float32(rl.GetScreenHeight())/2.0,
	)
}

func HandleBlockManipulation(camera *rl.Camera3D) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) || rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		pointRay := rl.GetScreenToWorldRay(GetScreenCenter(), *camera)
		maxRange := float32(10)
		step := float32(0.25)

		pos := pointRay.Position
		for i := float32(0); i < maxRange; i += step {
			blok := world.GetGlobalBlock(NJK(pos.X), NJK(pos.Y), NJK(pos.Z))
			if blok >= 2 {
				if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
					world.SetGlobalBlock(NJK(pos.X), NJK(pos.Y), NJK(pos.Z), 1)

				} else {
					/*newBlockBox := rl.BoundingBox{
						Min: blok.PlaceBlock,
						Max: rl.NewVector3(
							blok.PlaceBlock.X+1.0,
							blok.PlaceBlock.Y+1.0,
							blok.PlaceBlock.Z+1.0,
							),
					}*/
				}
				return
			}
			pos.X += pointRay.Direction.X * step
			pos.Y += pointRay.Direction.Y * step
			pos.Z += pointRay.Direction.Z * step
		}
	}

}

func NJK(f float32) int { //NJK znaci NEKA JEBENA KONVERZIJA IDK
	return int(math.Round(float64(f)))
}
