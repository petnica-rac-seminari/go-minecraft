package voda

type Cord struct {
	x int
	y int
	z int
}

type MalaReka struct {
	rank int
	deo  []Deo
}

type VelikaReka struct {
	rank int
	deo  []Deo
}

type Potok struct {
	deo []Deo
}

type Deo struct {
	start Cord
	end   Cord
}
