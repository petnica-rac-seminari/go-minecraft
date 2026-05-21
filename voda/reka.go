
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
	d.imavode = [c2.x - c1.x][c2.y - c1.y]bool
	d.cords1 = c1
	d.cords2 = c2
}

func (d *Deo) Postavi(c Cords) {
	d.imavode[c.x-d.cords1.x][c.y-d.cords1.y] = voda
}

func (r *Reka) AddDeo(d *Deo) {
	r.deo = append(r.deo, d)
}
