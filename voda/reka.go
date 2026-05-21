

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
}

func (d *Deo) Postavi(x int, y int) {
	d.imavode[x-x1][y-y1] = voda
}