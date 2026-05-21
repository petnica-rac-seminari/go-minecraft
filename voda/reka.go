package voda

type Cords struct {
	x float64
	y float64
}

type Reka struct {
	deo []Deo
}

type Deo struct {
	brzinaFloat     float64
	boja            string
	cords1          Cords
	cords2          Cords
	imavode         [][]bool
	centralnaLinija []Cords
}

func (d *Deo) NewDeo(c1 Cords, c2 Cords) {
	w := int(c2.x - c1.x)
	h := int(c2.y - c1.y)
	d.imavode = make([][]bool, w)
	for i, _ := range d.imavode {
		d.imavode[i] = make([]bool, h)
	}
	d.cords1 = c1
	d.cords2 = c2
}

func (d *Deo) Postavi(c Cords, v bool) {
	d.imavode[int(c.x-d.cords1.x)][int(c.y-d.cords1.y)] = v
}

func (r *Reka) AddDeo(d Deo) {
	r.deo = append(r.deo, d)
}

func (d *Deo) CentarDela() {

}
