package voda

func MalaErekcina(potok Potok, velike []VelikaReka) *MalaReka {
	reka := &MalaReka{}
	if len(potok.deo) == 0 {
		return reka
	}
	trenutna := potok.deo[len(potok.deo)-1].end
	for i := 0; i < 1000; i++ {
		sledeca := Cord{x: trenutna.x + 1, y: trenutna.y, z: trenutna.z}
		deo := Deo{start: trenutna, end: sledeca}
		reka.Add(&deo)
		for _, velika := range velike {
			if PipkaReku(sledeca, velika) {
				return reka
			}
		}
		trenutna = sledeca
	}
	return reka
}
