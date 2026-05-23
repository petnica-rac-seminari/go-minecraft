package voda

import (
	"main/blocks"
)

// Imaginarnaa funkcija koja menja blok
func setBlock(x, y, z int, block blocks.Block) {

}

// Moze da se prelije
// ako je to problem, treba popraviti (dodati zasigurne ograde svuda)
func potokMask(x, y, z int) {
	setBlock(x, y-1, z, blocks.BlockWater)
	setBlock(x-1, y-1, z, blocks.BlockWater)
	b := BlockTypeAt(x-2, y-2, z)
	setBlock(x-2, y-1, z, b)
	setBlock(x, y-1, z-1, blocks.BlockWater)
	b := BlockTypeAt(x, y-2, z-2)
	setBlock(x, y-1, z-2, b)
	setBlock(x, y, z, blocks.BlockAir)
	setBlock(x-1, y, z, blocks.BlockAir)
	setBlock(x, y, z-1, blocks.BlockAir)
}

func recicaMask(x, y, z int) {

	xDno := x
	zDno := z
	yDno := y - 10

	xd1 := xDno - 16
	xd2 := xDno + 16
	zd1 := zDno - 16
	zd2 := zDno + 16
	yd1 := yDno - 16
	for i := xd1; i <= xd2; i++ {
		for j := zd1; j <= zd2; j++ {
			for k := y; k >= yd1; k-- {
				if (i-xDno)*(i-xDno)+(j-zDno)*(j-zDno)+(k-yDno)*(k-yDno) <= 1700 {
					setBlock(i, k, j, blocks.BlockDirt)
				}
			}
		}
	}

	x1 := x - 5
	x2 := x + 5
	z1 := z - 5
	z2 := z + 5
	y1 := y - 5
	for i := x1; i <= x2; i++ {
		for j := z1; j <= z2; j++ {
			for k := y1; k <= y; k++ {
				if (i-x)*(i-x)+(j-z)*(j-z)+(k-y)*(k-y) <= 70 {
					setBlock(i, k, j, blocks.BlockWater)
				}
			}
		}
	}
	xVrh := x
	zVrh := z
	yVrh := y + 10

	xv1 := xVrh - 16
	xv2 := xVrh + 16
	zv1 := zVrh - 16
	zv2 := zVrh + 16
	for i := xv1; i <= xv2; i++ {
		for j := zv1; j <= zv2; j++ {
			for k := y; k <= yVrh; k++ {
				if (i-xVrh)*(i-xVrh)+(j-zVrh)*(j-zVrh)+(k-yVrh)*(k-yVrh) <= 1350 {
					setBlock(i, k, j, blocks.BlockAir)
				}
			}
		}
	}

}

func rekaMask(x, y, z int) {

	xDno := x
	zDno := z
	yDno := y - 15

	xd1 := xDno - 24
	xd2 := xDno + 24
	zd1 := zDno - 24
	zd2 := zDno + 24
	yd1 := yDno - 24
	for i := xd1; i <= xd2; i++ {
		for j := zd1; j <= zd2; j++ {
			for k := y; k >= yd1; k-- {
				if (i-xDno)*(i-xDno)+(j-zDno)*(j-zDno)+(k-yDno)*(k-yDno) <= 4100 {
					setBlock(i, k, j, blocks.BlockDirt)
				}
			}
		}
	}

	x1 := x - 8
	x2 := x + 8
	z1 := z - 8
	z2 := z + 8
	y1 := y - 8
	for i := x1; i <= x2; i++ {
		for j := z1; j <= z2; j++ {
			for k := y1; k <= y; k++ {
				if (i-x)*(i-x)+(j-z)*(j-z)+(k-y)*(k-y) <= 230 {
					setBlock(i, k, j, blocks.BlockWater)
				}
			}
		}
	}
	xVrh := x
	zVrh := z
	yVrh := y + 15

	xv1 := xVrh - 24
	xv2 := xVrh + 24
	zv1 := zVrh - 24
	zv2 := zVrh + 24
	for i := xv1; i <= xv2; i++ {
		for j := zv1; j <= zv2; j++ {
			for k := y; k <= yVrh; k++ {
				if (i-xVrh)*(i-xVrh)+(j-zVrh)*(j-zVrh)+(k-yVrh)*(k-yVrh) <= 4100 {
					setBlock(i, k, j, blocks.BlockAir)
				}
			}
		}
	}

}
