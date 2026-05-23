package voda

func VelikaErekcina(rank int, start Cord, postojece *[]VelikaReka) *VelikaReka {
	reka := &VelikaReka{rank: rank}
	trenutna := start
	for i := 0; i < 200; i++ {
		sledeca := Cord{x: trenutna.x, y: trenutna.y + 1, z: trenutna.z}
		deo := Deo{start: trenutna, end: sledeca}
		for _, druga := range *postojece {
			if druga.rank < reka.rank {
				if PipkaReku(sledeca, druga) {
					reka.Add(&deo)
					return reka
				}
			}
		}
		reka.Add(&deo)
		trenutna = sledeca
	}
	return reka
}
func PipkaReku(c Cord, reka VelikaReka) bool {
	for _, deo := range reka.deo {
		if deo.start == c || deo.end == c {
			return true
		}
	}
	return false
}
