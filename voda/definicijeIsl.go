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

func (m *MalaReka) Add(d *Deo) {
	m.deo = append(m.deo, *d)
}
func (v *VelikaReka) Add(d *Deo) {
	v.deo = append(v.deo, *d)
}
func (p *Potok) Add(d *Deo) {
	p.deo = append(p.deo, *d)
}
