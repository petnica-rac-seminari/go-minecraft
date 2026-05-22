package main

import (
	"fmt"
	"image/color"
	"math"
	"time"
)

func skyColor(currentTick int) color.RGBA {
	Cycle := (math.Cos((math.Pi * float64(currentTick%480) / 120.0)) + 1) * 255 / 2
	c := color.RGBA{
		R: uint8(Cycle * 120 / 255),
		G: uint8(Cycle * 223 / 255),
		B: uint8((Cycle + 25) * 255 / 280),
		A: 255,
	}

	return c
}

func main() {

	tick := 0
	for {

		fmt.Println(skyColor(tick))

		tick++
		time.Sleep(time.Second)
	}
}
