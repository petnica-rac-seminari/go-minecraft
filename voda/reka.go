

type Reka struct {
	deo []Deo
}

type Deo struct {
	brzinaFloat float64
	boja        string
	x1, y1      float64
	x2, y2      float64
	imavode     [][]bool
}

func (d *Deo) NewDeo(x1, y1, x2, y2 int) {
	d.imavode = [x2 - x1][y2 - y1]bool
	d.x1 = x1
	d.x2 = x2
	d.y1 = y1
	d.y2 = y2
}

func (d *Deo) Postavi(x int, y int) {
	d.imavode[x-x1][y-y1] = voda
}

func (r *Reka) AddDeo(d *Deo) {
	r.deo = append(r.deo, d)
}