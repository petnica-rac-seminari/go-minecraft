package voda

func findNearestAir(c *Cord) {
	// BFS

	var queue [][]int
	queue = append(queue, []int{c.x, c.z})
	for {

		element := queue[0]
		queue = queue[1:]
		x := element[0]
		z := element[1]
		if x == c.x {
			if z > c.z {
				queue = append(queue, []int{x, z + 1})
				queue = append(queue, []int{x + 1, z})
				queue = append(queue, []int{x - 1, z})
			} else {
				queue = append(queue, []int{x, z - 1})
				queue = append(queue, []int{x + 1, z})
				queue = append(queue, []int{x - 1, z})
			}
		}
		if z == c.z {
			if x > c.x {
				queue = append(queue, []int{x + 1, z})
			} else {
				queue = append(queue, []int{x - 1, z})
			}
		}
		if z > c.z && x > c.x {
			queue = append(queue, []int{x + 1, z})
		}
		if z > c.z && x < c.x {
			queue = append(queue, []int{x - 1, z})
		}
		if z < c.z && x > c.x {
			queue = append(queue, []int{x + 1, z})
		}
		if z < c.z && x < c.x {
			queue = append(queue, []int{x - 1, z})
		}

		// Umesto funkcije which block koristicemo funkciju za trazenja visine na {x,y}
		if hight < c.y {
			c = &Cord{x: x, y: c.y - 1, z: z}
			break
		}
	}
}

func GeneratePotok(cordsForPotok Cord) {

	var privremeniPotok Potok
	var dio Deo

	for y := cordsForPotok.y; y > baseHeight; y-- {
		dio.start = cordsForPotok
		findNearestAir(&cordsForPotok)
		dio.end = cordsForPotok
		if dio.start.y > 80 {
			dio.sirina = 1
		} else if dio.start.y > 60 {
			dio.sirina = 2
		} else {
			dio.sirina = 3
		}
		privremeniPotok.Add(&dio)
	}
	worldReke.AddPotoci(&privremeniPotok)
}
