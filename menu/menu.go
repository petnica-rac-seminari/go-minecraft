package menu

import (
	"os"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var IsMenu bool = true
var Seed int = 100
var pozadina rl.Texture2D
var creditsSlika rl.Texture2D

var prikaziCredits bool = false

var seedTekst string = "100"
var unosAktivan bool = false

func UcitajMenuSliku() {
	pozadina = rl.LoadTexture("menufinal.png")
	creditsSlika = rl.LoadTexture("credits.png")
}

func UnloadujMenuSliku() {
	rl.UnloadTexture(pozadina)
	rl.UnloadTexture(creditsSlika)
}

func Crtaj() {
	if prikaziCredits {
		rl.ClearBackground(rl.Black)
		izvor := rl.NewRectangle(0, 0, float32(creditsSlika.Width), float32(creditsSlika.Height))
		odrediste := rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
		centar := rl.NewVector2(0, 0)
		rl.DrawTexturePro(creditsSlika, izvor, odrediste, centar, 0.0, rl.White)

		if rl.IsKeyPressed(rl.KeyEscape) || rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			prikaziCredits = false
		}
		return
	}

	rl.ClearBackground(rl.DarkGray)

	izvor := rl.NewRectangle(0, 0, float32(pozadina.Width), float32(pozadina.Height))
	odrediste := rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
	centar := rl.NewVector2(0, 0)
	rl.DrawTexturePro(pozadina, izvor, odrediste, centar, 0.0, rl.White)

	misPozicija := rl.GetMousePosition()

	playRect := rl.NewRectangle(720, 500, 400, 100)
	playBoja := rl.Gray
	if rl.CheckCollisionPointRec(misPozicija, playRect) {
		playBoja = rl.LightGray
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			IsMenu = false
			rl.DisableCursor()
		}
	}
	rl.DrawRectangleRec(playRect, playBoja)
	rl.DrawText("PLAY", 720+(400-rl.MeasureText("PLAY", 30))/2, 535, 30, rl.Black)

	creditsRect := rl.NewRectangle(720, 650, 400, 100)
	creditsBoja := rl.Gray
	if rl.CheckCollisionPointRec(misPozicija, creditsRect) {
		creditsBoja = rl.LightGray
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			prikaziCredits = true
		}
	}
	rl.DrawRectangleRec(creditsRect, creditsBoja)
	rl.DrawText("CREDITS", 720+(400-rl.MeasureText("CREDITS", 30))/2, 685, 30, rl.Black)

	quitRect := rl.NewRectangle(720, 800, 400, 100)
	quitBoja := rl.Gray
	if rl.CheckCollisionPointRec(misPozicija, quitRect) {
		quitBoja = rl.Maroon
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			rl.CloseWindow()
			os.Exit(0)
		}
	}
	rl.DrawRectangleRec(quitRect, quitBoja)
	rl.DrawText("QUIT", 720+(400-rl.MeasureText("QUIT", 30))/2, 835, 30, rl.Black)

	seedRect := rl.NewRectangle(1200, 500, 150, 100)

	if rl.CheckCollisionPointRec(misPozicija, seedRect) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		unosAktivan = true
	} else if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		unosAktivan = false
	}

	if unosAktivan {
		rl.DrawRectangleRec(seedRect, rl.Black)
		rl.DrawRectangleLinesEx(seedRect, 3, rl.Blue)
	} else {
		rl.DrawRectangleRec(seedRect, rl.Black)
	}

	if unosAktivan {
		kljuc := rl.GetCharPressed()
		for kljuc > 0 {
			if kljuc >= 48 && kljuc <= 57 && len(seedTekst) < 6 {
				seedTekst += string(kljuc)
			}
			kljuc = rl.GetCharPressed()
		}

		if rl.IsKeyPressed(rl.KeyBackspace) && len(seedTekst) > 0 {
			seedTekst = seedTekst[:len(seedTekst)-1]
		}

		if len(seedTekst) > 0 {
			if s, err := strconv.Atoi(seedTekst); err == nil {
				Seed = s
			}
		} else {
			Seed = 0
		}
	}

	if len(seedTekst) == 0 && unosAktivan {
		rl.DrawText("|", 1210, 535, 30, rl.White)
	} else {
		rl.DrawText(seedTekst, 1200+(150-rl.MeasureText(seedTekst, 30))/2, 535, 30, rl.White)
	}

	rl.DrawText("SEED:", 1200, 465, 24, rl.White)

	if rl.IsKeyPressed(rl.KeyEnter) {
		IsMenu = false
		rl.DisableCursor()
	}
}
