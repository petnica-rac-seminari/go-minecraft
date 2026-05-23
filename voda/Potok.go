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
		if BlockTypeAt(x, c.y, z) == block.BlockAir {
			c = &Cord{x: x, y: c.y - 1, z: z}
			break
		}
	}
}

func GeneratePotok(startX, startY, startZ int) {

	cordsForPotok := Cord{x: startX, y: startY, z: startZ}

	var privremeniPotok Potok
	var dio Deo

	for y := startY; y > baseHeight; y-- {
		dio.start = cordsForPotok
		findNearestAir(&cordsForPotok)
		dio.end = cordsForPotok
		privremeniPotok.Add(&dio)
	}
	worldReke.AddPotoci(&privremeniPotok)
}
