package main

import (
	"fmt"

	"github.com/KEINOS/go-noise"
)

const seed = 100

const chunkSize = 16
const viewChunks = 5

const width = chunkSize * viewChunks
const height = chunkSize * viewChunks

const smoothness = 15

// noise
func Octaves(n noise.Generator, v float64, x int, y int, br_o int) float64 {
	value := v
	amplitude := 1.0
	frequency := 1.0
	maxAmp := 1.0

	for i := 0; i < br_o; i++ {
		nx := (float64(x) / smoothness) * frequency
		ny := (float64(y) / smoothness) * frequency

		value += n.Eval64(nx, ny) * amplitude

		maxAmp += amplitude

		amplitude *= 0.5
		frequency *= 2.0
	}

	return value / maxAmp
}

func assignValues(n noise.Generator, br_oct int, startX, startY int) []float64 {

	values := make([]float64, width*height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			worldX := startX + x
			worldY := startY + y

			v := n.Eval64(float64(worldX)/smoothness, float64(worldY)/smoothness)
			v = Octaves(n, v, worldX, worldY, br_oct)

			values[y*width+x] = v
		}
	}

	return values
}

func newNoise(seedMod int, br_oct int, startX, startY int) []float64 {
	n, _ := noise.New(noise.OpenSimplex, int64(seed+seedMod))
	return assignValues(n, br_oct, startX, startY)
}

func HeightMap(playerX, playerZ, seed int) [][]int {
	coords := make([][]int, viewChunks*16)
	for i := range coords {
		coords[i] = make([]int, viewChunks*16)
	}

	startX := playerX*chunkSize - (width / 2)
	startZ := playerZ*chunkSize - (height / 2)

	values := newNoise(0, 1, startX, startZ)
	values1 := newNoise(1, 7, startX, startZ)

	for x := 0; x < width; x++ {
		for z := 0; z < height; z++ {

			v := values[z*width+x]
			norm := (v + 1.0) / 2.0 * 16

			v = values1[z*width+x]
			norm1 := (v + 1.0) / 2.0 * 16

			normf := int(norm+norm1) / 2

			coords[x][z] = normf
		}
	}

	return coords
}

func main() {
	coords := HeightMap(0, 0, 100)

	for x := 0; x < width; x++ {
		for z := 0; z < height; z++ {
			fmt.Printf("%v ", coords[x][z])
		}
		fmt.Println()
	}
}
