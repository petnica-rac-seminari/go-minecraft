package voda

type Cord struct {
	x int
	y int
	z int
}

type DeoPointer struct {
}

type Reke struct {
	velike []VelikaReka
	male   []MalaReka
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
	start   Cord
	end     Cord
	pointer DeoPointer
}

func NewDeo(x1 int, y1 int, z1 int, x2 int, y2 int, z2 int) Deo {
	d := Deo{
		start: Cord{x: x1, y: y1, z: z1},
		end:   Cord{x: x1, y: y1, z: z1},
	}
	return d
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

func (r *Reke) AddMale(mr *MalaReka) {
	mr.rank = len(r.male)
	r.male = append(r.male, *mr)
}
func (r *Reke) AddVelike(vr *VelikaReka) {
	vr.rank = len(r.velike)
	r.velike = append(r.velike, *vr)
}
