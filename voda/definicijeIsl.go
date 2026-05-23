package voda

const (
	baseHeight                  int = 20 // Visina na kojoj potok postaje reka // file Potok.go
	radiusPretrageZaIzvore      int = 512
	brojIzvoraURadiosu          int = 10
	minUdaljenostIzmedjuIzvora  int = 45
	visinaOdVrhaSvetaZaIzvor    int = 5
	minVisinaZaGenerisanjeIzvor int = 44
	velikaHight                 int = 12
)

var (
	WorldReke   Reke
	listaIzvora map[Cord]bool
)

type Cord struct {
	x int
	y int
	z int
}

type DeoPointer struct {
	mali   *MalaReka
	veliki *VelikaReka
	potok  *Potok
}

type Reke struct {
	velike []VelikaReka
	male   []MalaReka
	potoci []Potok
}

type MalaReka struct {
	rank    int
	deo     []Deo
	pointer *Reke
}

type VelikaReka struct {
	rank    int
	deo     []Deo
	pointer *Reke
}

type Potok struct {
	deo     []Deo
	pointer *Reke
}

type Deo struct {
	start   Cord
	end     Cord
	pointer DeoPointer
	sirina  int
}

func NewDeo(x1 int, y1 int, z1 int, x2 int, y2 int, z2 int) Deo {
	d := Deo{
		start: Cord{x: x1, y: y1, z: z1},
		end:   Cord{x: x1, y: y1, z: z1},
	}
	return d
}

func (m *MalaReka) Add(d *Deo) {
	d.pointer.mali = *&m
	d.pointer.potok = nil
	d.pointer.veliki = nil
	m.deo = append(m.deo, *d)
}
func (v *VelikaReka) Add(d *Deo) {
	d.pointer.mali = nil
	d.pointer.veliki = *&v
	d.pointer.potok = nil
	v.deo = append(v.deo, *d)
}
func (p *Potok) Add(d *Deo) {
	d.pointer.mali = nil
	d.pointer.veliki = nil
	d.pointer.potok = &*p
	p.deo = append(p.deo, *d)
}

func (r *Reke) AddMale(mr *MalaReka) {
	mr.rank = len(r.male)
	r.male = append(r.male, *mr)
}
func (r *Reke) AddVelike(vr *VelikaReka) {
	vr.pointer = *&r
	vr.rank = len(r.velike)
	r.velike = append(r.velike, *vr)
}
func (r *Reke) AddPotoci(p *Potok) {
	p.pointer = *&r
	r.potoci = append(r.potoci, *p)
}
