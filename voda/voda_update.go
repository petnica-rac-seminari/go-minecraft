package voda

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func generisanjeVode(poz *Cord) {
	PostaviIzvore(poz)
	for k, i := range listaIzvora {
		if i {
			GeneratePotok(k)
		}
	}
}

//jebeno ne znam kako sta izgleda
//potrebno mi je da postoji chan koji ce mi slati poziciju korisnika da bi ja mogao da generisem izvore a potom potoke
//ovo treba da bude kao go update_voda() jer kao tako je okej

func update_voda(ch <-chan rl.Camera3D) { //pozicija pozicija, some rl.camera shit
	for poz := range ch {
		pos := Cord{x: int(math.Floor(float64(poz.Position.X))), y: int(math.Floor(float64(poz.Position.Y))), z: int(math.Floor(float64(poz.Position.Z)))}
		//pozicija igraca iz cha
		generisanjeVode(&pos)
	}
}
