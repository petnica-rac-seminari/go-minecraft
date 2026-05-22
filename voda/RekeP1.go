package RekeP1

func GenerateBlueprintP1(startX, startY int) []float64 {

}
func BlockTypeAt(x, y, z int) int {

}

func findNearestAir(startX, startY, startZ int) (int, int, int) {
	// BFS

	var queue [][]int
	queue = append(queue, []int{startX, startZ})
	for {

		element := queue[0]
		queue = queue[1:]
		x := element[0]
		z := element[1]
		if x == startX {
			if z > startZ {
				queue = append(queue, []int{x, z + 1})
				queue = append(queue, []int{x + 1, z})
				queue = append(queue, []int{x - 1, z})
			} else {
				queue = append(queue, []int{x, z - 1})
				queue = append(queue, []int{x + 1, z})
				queue = append(queue, []int{x - 1, z})
			}
		}
		if z == startZ {
			if x > startX {
				queue = append(queue, []int{x + 1, z})
			} else {
				queue = append(queue, []int{x - 1, z})
			}
		}
		if z > startZ && x > startX {
			queue = append(queue, []int{x + 1, z})
		}
		if z > startZ && x < startX {
			queue = append(queue, []int{x - 1, z})
		}
		if z < startZ && x > startX {
			queue = append(queue, []int{x + 1, z})
		}
		if z < startZ && x < startX {
			queue = append(queue, []int{x - 1, z})
		}

		if BlockTypeAt(x, startY, z) == block.BlockAir {
			return x, startY, z
		}
	}
}

func GenerateRiverP1(startX, startY, startZ int) []float64 {

	var baseHeight int = 20 // Visina na kojoj potok postaje reka

	var x1, y1, z1 int
	x2, y2, z2 := startX, startY, startZ

	var blueprint []definicijeisl.Deo

	for y := startY; y > baseHeight; y-- {
		x1, y1, z1 = x2, y2, z2
		x2, y2, z2 = findNearestAir(x1, y, z1)
		d := NewDeo(x1, y1, z1, x2, y2, z2)
		blueprint = append(blueprint, d)
	}

	return blueprint
}
