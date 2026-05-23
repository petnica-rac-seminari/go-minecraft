package menu

import (
	"main/blocks"
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

	playRect := rl.NewRectangle(100, 600, 400, 100)
	playBoja := rl.Gray
	if rl.CheckCollisionPointRec(misPozicija, playRect) {
		playBoja = rl.LightGray
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			IsMenu = false
			rl.DisableCursor()
		}
	}
	rl.DrawRectangleRec(playRect, playBoja)
	rl.DrawText("PLAY", int32(playRect.X)+(int32(playRect.Width)-rl.MeasureText("PLAY", 30))/2, int32(playRect.Y)+35, 30, rl.Black)

	creditsRect := rl.NewRectangle(100, 750, 400, 100)
	creditsBoja := rl.Gray
	if rl.CheckCollisionPointRec(misPozicija, creditsRect) {
		creditsBoja = rl.LightGray
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			prikaziCredits = true
		}
	}
	rl.DrawRectangleRec(creditsRect, creditsBoja)
	rl.DrawText("CREDITS", int32(creditsRect.X)+(int32(creditsRect.Width)-rl.MeasureText("CREDITS", 30))/2, int32(creditsRect.Y)+35, 30, rl.Black)

	quitRect := rl.NewRectangle(100, 900, 400, 100)
	quitBoja := rl.Gray
	if rl.CheckCollisionPointRec(misPozicija, quitRect) {
		quitBoja = rl.Maroon
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			rl.CloseWindow()
			os.Exit(0)
		}
	}
	rl.DrawRectangleRec(quitRect, quitBoja)
	rl.DrawText("QUIT", int32(quitRect.X)+(int32(quitRect.Width)-rl.MeasureText("QUIT", 30))/2, int32(quitRect.Y)+35, 30, rl.Black)

	seedRect := rl.NewRectangle(580, 600, 150, 100)

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
		rl.DrawText("|", int32(seedRect.X)+10, int32(seedRect.Y)+35, 30, rl.White)
	} else {
		rl.DrawText(seedTekst, int32(seedRect.X)+(int32(seedRect.Width)-rl.MeasureText(seedTekst, 30))/2, int32(seedRect.Y)+35, 30, rl.White)
	}

	rl.DrawText("SEED:", int32(seedRect.X), int32(seedRect.Y)-35, 24, rl.White)

	if rl.IsKeyPressed(rl.KeyEnter) {
		IsMenu = false
		rl.DisableCursor()
	}
}

func CrtajHotbar(aktivniBlok blocks.Block) {
	sirinaEkrana := float32(rl.GetScreenWidth())
	visinaEkrana := float32(rl.GetScreenHeight())

	slotovi := [9]blocks.Block{1, 2, 3, 4, 5, 6, 7, 8, 9}
	imena := [9]string{"Water", "Grass", "Dirt", "Stone", "Snow", "Log", "Leaves", "Niggerrack", "Sand"}
	tasteri := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	velicinaSlota := float32(70)
	razmak := float32(10)
	ukupnaSirina := (velicinaSlota * 9) + (razmak * 4)

	startX := (sirinaEkrana - ukupnaSirina) / 2
	startY := visinaEkrana - velicinaSlota - 25

	for i := 0; i < 9; i++ {
		rect := rl.NewRectangle(startX+float32(i)*(velicinaSlota+razmak), startY, velicinaSlota, velicinaSlota)

		bojaPozadine := rl.Fade(rl.DarkGray, 0.7)
		bojaOkvira := rl.Gray
		var debljinaOkvira float32 = 2

		if slotovi[i] == aktivniBlok {
			bojaPozadine = rl.Fade(rl.LightGray, 0.8)
			bojaOkvira = rl.Gold
			debljinaOkvira = 4
		}

		rl.DrawRectangleRec(rect, bojaPozadine)
		rl.DrawRectangleLinesEx(rect, debljinaOkvira, bojaOkvira)

		rl.DrawText(tasteri[i], int32(rect.X)+6, int32(rect.Y)+6, 12, rl.White)

		sirinaTeksta := rl.MeasureText(imena[i], 12)
		visinaTeksta := int32(12)
		rl.DrawText(
			imena[i],
			int32(rect.X)+(int32(velicinaSlota)-sirinaTeksta)/2,
			int32(rect.Y)+(int32(velicinaSlota)-visinaTeksta)/2,
			12,
			rl.Black,
		)
	}
}
