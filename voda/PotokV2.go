package voda

type cods struct {
	x     int
	z     int
	check bool
	deep  int
}

func CheckVisina(c *cods, y int) bool {
	if hight(c.x, c.z) == y {

	} else if hight(c.x, c.z) > y {
		c.check = true
	} else {
		return true
	}
	return false
}

func FindNext(co *Cord) Cord {

	queue := make(map[cods][]cods, 0)
	var c cods = cods{x: co.x, z: co.z, check: false, deep: 0}
	var y int = co.y
	queue[c] = nil
	i := 0
	jezero := false
	for !jezero {
		jezero = true
		for cords, hisory := range queue {
			if cords.check || cords.deep != i {
				continue
			} else {
				jezero = false
				c = cords
				c.x++
				if queue[c] == nil {
					queue[c] = append(hisory, c)
					if CheckVisina(&c, y) {
						geto := Cord{x: c.x, z: c.z, y: y}
						return geto
					}
				}
				c = cords
				c.x++
				c.z++
				if queue[c] == nil {
					queue[c] = append(hisory, c)
					if CheckVisina(&c, y) {
						geto := Cord{x: c.x, z: c.z, y: y}
						return geto
					}
				}
				c = cords
				c.z++
				if queue[c] == nil {
					queue[c] = append(hisory, c)
					if CheckVisina(&c, y) {
						geto := Cord{x: c.x, z: c.z, y: y}
						return geto
					}
				}
				c = cords
				c.z++
				c.x--
				if queue[c] == nil {
					queue[c] = append(hisory, c)
					if CheckVisina(&c, y) {
						geto := Cord{x: c.x, z: c.z, y: y}
						return geto
					}
				}
				c = cords
				c.x--
				if queue[c] == nil {
					queue[c] = append(hisory, c)
					if CheckVisina(&c, y) {
						geto := Cord{x: c.x, z: c.z, y: y}
						return geto
					}
				}
				c = cords
				c.x--
				c.z--
				if queue[c] == nil {
					queue[c] = append(hisory, c)
					if CheckVisina(&c, y) {
						geto := Cord{x: c.x, z: c.z, y: y}
						return geto
					}
				}
				c = cords
				c.z--
				if queue[c] == nil {
					queue[c] = append(hisory, c)
					if CheckVisina(&c, y) {
						geto := Cord{x: c.x, z: c.z, y: y}
						return geto
					}
				}
				c = cords
				c.z--
				c.x++
				if queue[c] == nil {
					queue[c] = append(hisory, c)
					if CheckVisina(&c, y) {
						geto := Cord{x: c.x, z: c.z, y: y}
						return geto
					}
				}
				c.check = true
			}
		}

		i++
	}
	return Cord{x: 0, y: 1200, z: 0}
}

func NewPotok(c Cord) {

}
