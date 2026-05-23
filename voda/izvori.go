package voda

import "math"

type OnlyXandZCord struct {
	x int
	z int
}

func ValidneCordinateZaIzvor(mapaZaIzvore map[int][]OnlyXandZCord) []OnlyXandZCord {
	for i := minVisinaZaGenerisanjeIzvor; i <= 64-visinaOdVrhaSvetaZaIzvor; i++ {
		for _, j := range mapaZaIzvore[i] {
			postoji := true
			for l, _ := range listaIzvora { // funkcija math.Abc je jebeno sranje
				if math.Abs(float64(j.x-l.x)) <= float64(minUdaljenostIzmedjuIzvora) || math.Abs(float64(j.z-l.z)) <= float64(minUdaljenostIzmedjuIzvora) {
					postoji = false
				}
			}
			if postoji {
				NewPotok(Cord{x: j.x, z: j.z, y: i})
			}
		}
	}
}

func PostaviIzvore(PozicijaIgraca *Cord) {
	mapaZaIzvore := make(map[int][]OnlyXandZCord)
	for i := -radiusPretrageZaIzvore; i <= radiusPretrageZaIzvore; i++ {
		for j := -radiusPretrageZaIzvore; j <= radiusPretrageZaIzvore; j++ {
			h := hight(PozicijaIgraca.x+i, PozicijaIgraca.z+j)
			if h >= minVisinaZaGenerisanjeIzvor {
				mapaZaIzvore[h] = append(mapaZaIzvore[h], OnlyXandZCord{x: i, z: j})
			}
		}
	}
}
