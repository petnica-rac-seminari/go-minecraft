package voda
	import (
		"main/blocks")
var (
	granicaY float32 = 10.0
)

func postaviVodu(world *World) {
	for int i := 0; i < len(world.Chunks); i++ {
		chunk := & world.Chunks[i]
		
				for y := 0; y < 64; y++ {
					
					if float32(y) < granicaY {
						if chunk.Blocks[x][y][z] == blocks.Air {
							chunk.Blocks[x][y][z] = blocks.Water

						}
					}
					
				}
			}
		}