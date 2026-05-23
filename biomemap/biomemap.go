/*
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/KEINOS/go-noise"
)

const seed = 100
const width, height = 512, 512
const smoothness = 40

func Octaves(n noise.Generator, v float64, x int, y int, br_o int) float64 {
	var (
		value             = v
		amplitude float64 = 1.0
		frequency float64 = 1.0
		maxAmp    float64 = 1.0
	)

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

func assignValues(n noise.Generator, br_oct int) []float64 {
	values := make([]float64, width*height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			v := n.Eval64(float64(x)/smoothness, float64(y)/smoothness)
			v = Octaves(n, v, x, y, br_oct)
			values[x*height+y] = v
		}
	}
	return values
}

func newNoise(seedMod int, br_oct int) (noise.Generator, []float64) {
	n, _ := noise.New(noise.OpenSimplex, int64(seed+seedMod))
	values := assignValues(n, br_oct)

	return n, values
}

func main() {

	n, values := newNoise(0, 1)
	n1, values1 := newNoise(1, 7)

	_, _ = n, n1

	min, max := values[0], values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	min1, max1 := values1[0], values1[0]
	for _, v := range values1 {
		if v < min1 {
			min1 = v
		}
		if v > max1 {
			max1 = v
		}
	}

	// Second pass: normalize using actual range
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			v := values[x*height+y]
			normalized := (v - min) / (max - min)

			v = values1[x*height+y]
			normalized1 := (v - min1) / (max1 - min1)

			norm := (normalized*255)/2 + (normalized1*255)/2

			if norm < 145 {
				img.SetRGBA(x, y, color.RGBA{31, 22, 186, 255})
			} else if norm < 148 {
				img.SetRGBA(x, y, color.RGBA{196, 186, 88, 255})
			} else if norm < 200 {
				img.SetRGBA(x, y, color.RGBA{19, 145, 27, 255})
			} else {
				img.SetRGBA(x, y, color.RGBA{75, 76, 79, 255})
			}
		}
	}

	f, err := os.Create("noise.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, img)
}
*/

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/KEINOS/go-noise"
)

const seed = 100

// 1 pixel = 1 CHUNK
const width, height = 512, 512

// chunk size used only for scaling noise
const chunkSize = 16

// controls continent scale (bigger = larger continents)
const smoothness = 640

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

// CHUNK VISUAL MAP (1 pixel = 1 chunk)
func assignValues(n noise.Generator, br_oct int) []float64 {
	values := make([]float64, width*height)

	for cx := 0; cx < width; cx++ {
		for cy := 0; cy < height; cy++ {

			// convert chunk coordinate → world coordinate (center of chunk)
			worldX := cx * chunkSize
			worldY := cy * chunkSize

			v := n.Eval64(float64(worldX)/smoothness, float64(worldY)/smoothness)
			v = Octaves(n, v, worldX, worldY, br_oct)

			values[cy*width+cx] = v
		}
	}

	return values
}

func newNoise(seedMod int, br_oct int) (noise.Generator, []float64) {
	n, _ := noise.New(noise.OpenSimplex, int64(seed+seedMod))
	values := assignValues(n, br_oct)
	return n, values
}

func main() {

	_, values := newNoise(0, 1)
	_, values1 := newNoise(1, 7)

	// normalize first noise
	min, max := values[0], values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// normalize second noise
	min1, max1 := values1[0], values1[0]
	for _, v := range values1 {
		if v < min1 {
			min1 = v
		}
		if v > max1 {
			max1 = v
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			v := values[y*width+x]
			norm1 := (v - min) / (max - min)

			v = values1[y*width+x]
			norm2 := (v - min1) / (max1 - min1)

			norm := (norm1*255)/2 + (norm2*255)/2

			if norm < 145 {
				img.SetRGBA(x, y, color.RGBA{31, 22, 186, 255})
			} else if norm < 148 {
				img.SetRGBA(x, y, color.RGBA{196, 186, 88, 255})
			} else if norm < 200 {
				img.SetRGBA(x, y, color.RGBA{19, 145, 27, 255})
			} else {
				img.SetRGBA(x, y, color.RGBA{75, 76, 79, 255})
			}
		}
	}

	f, err := os.Create("noise.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	png.Encode(f, img)
}
