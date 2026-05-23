package main

import (
	"math"
	"math/rand"

	reljef "main/Reljef"
	nebo "main/dayNightCycle"
	"main/menu"
	"main/navigation"
	"main/oblaci"
	"main/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const render_dist = 5

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1600, 900, "Raylib Go - 3D Kocka i Skakanje")
	defer rl.CloseWindow()

	// Muzika
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	music := rl.LoadMusicStream("muzika\\Minecraft.mp3")
	defer rl.UnloadMusicStream(music)

	rl.PlayMusicStream(music)

	menu.UcitajMenuSliku()
	defer menu.UnloadujMenuSliku()

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 160.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 90.0
	camera.Projection = rl.CameraPerspective
	rl.EnableCursor()
	rl.SetTargetFPS(60)
	rng := rand.New(rand.NewSource(int64(menu.Seed)))
	_ = rng

	var currentTick float32 = 0
	var clouds []oblaci.CLOUDS = oblaci.GenerateClouds()
	sunce_model := nebo.RenderSun()
	var dim bool = true

	for !rl.WindowShouldClose() {
		currentTick += rl.GetFrameTime()
		rl.UpdateMusicStream(music)
		if rl.IsMusicStreamPlaying(music) == false {
			rl.PlayMusicStream(music)
		}

		if !menu.IsMenu {
			navigation.HandleBlockManipulation(&camera)
			navigation.HandleMovement(&camera)
			playerCX := int(math.Floor(float64(camera.Position.X) / 16.0))
			playerCZ := int(math.Floor(float64(camera.Position.Z) / 16.0))

			halfDist := render_dist / 2
			for z := -halfDist; z <= halfDist; z++ {
				for x := -halfDist; x <= halfDist; x++ {
					pos := world.ChunkPos{X: playerCX + x, Z: playerCZ + z}

					if _, exists := world.LoadedChunks[pos]; !exists {
						if dim {
							c := reljef.GenerateOW(pos.X*16, pos.Z*16, menu.Seed)
							world.LoadedChunks[pos] = &c
						} else {
							c := reljef.GenerateNether(pos.X*16, pos.Z*16, menu.Seed)
							world.LoadedChunks[pos] = &c
						}
					}
				}
			}

			// SKLADIŠTIMO TASTER U PROMENLJIVU DA GA NE BISMO PROGUTALI!
			trenutniTaster := rl.GetKeyPressed()

			if trenutniTaster == rl.KeyG {
				camera.Position.Y += 5
				world.LoadedChunks = make(map[world.ChunkPos]*world.Chunk)
				dim = !dim
				for z := -halfDist; z <= halfDist; z++ {
					for x := -halfDist; x <= halfDist; x++ {
						pos := world.ChunkPos{X: playerCX + x, Z: playerCZ + z}

						if _, exists := world.LoadedChunks[pos]; !exists {
							if dim {
								c := reljef.GenerateOW(pos.X*16, pos.Z*16, menu.Seed)
								world.LoadedChunks[pos] = &c
							} else {
								c := reljef.GenerateNether(pos.X*16, pos.Z*16, menu.Seed)
								world.LoadedChunks[pos] = &c
							}
						}
					}
				}

				pos := world.ChunkPos{X: playerCX, Z: playerCZ}
				playerX := (int(math.Abs(float64(camera.Position.X))) % 16)
				playerZ := (int(math.Abs(float64(camera.Position.Z))) % 16)
				playerY := int(camera.Position.Y)

				if playerX > 13 {
					if playerZ > 13 {
						reljef.GeneratePortal(playerX-4, playerY, playerZ-4, world.LoadedChunks[pos].Blocks, dim)
					} else if playerZ < 3 {
						reljef.GeneratePortal(playerX-4, playerY, playerZ+4, world.LoadedChunks[pos].Blocks, dim)
					} else {
						reljef.GeneratePortal(playerX-4, playerY, playerZ, world.LoadedChunks[pos].Blocks, dim)
					}
				} else if playerX < 3 {
					if playerZ > 13 {
						reljef.GeneratePortal(playerX+4, playerY, playerZ-4, world.LoadedChunks[pos].Blocks, dim)
					} else if playerZ < 3 {
						reljef.GeneratePortal(playerX+4, playerY, playerZ+4, world.LoadedChunks[pos].Blocks, dim)
					} else {
						reljef.GeneratePortal(playerX+4, playerY, playerZ, world.LoadedChunks[pos].Blocks, dim)
					}
				} else {
					reljef.GeneratePortal(playerX, playerY, playerZ, world.LoadedChunks[pos].Blocks, dim)
				}

				camera.Position.Y += 5
			}

			oblaci.MoveClouds(clouds, camera.Position.Y)

		}

		rl.BeginDrawing()

		if !menu.IsMenu {
			rl.DisableCursor()
			rl.ClearBackground(nebo.SkyColor(int(currentTick), dim))
			rl.BeginMode3D(camera)

			playerCX := int(math.Floor(float64(camera.Position.X) / 16.0))
			playerCZ := int(math.Floor(float64(camera.Position.Z) / 16.0))

			halfDist := render_dist / 2
			for z := -halfDist; z <= halfDist; z++ {
				for x := -halfDist; x <= halfDist; x++ {
					pos := world.ChunkPos{X: playerCX + x, Z: playerCZ + z}
					if chunk, exists := world.LoadedChunks[pos]; exists {
						world.RenderChunk(chunk)
					}
					if structure, exists := world.LoadedStructures[pos]; exists {
						world.RenderChunk(structure)
					}
				}
			}

			oblaci.DrawClouds(clouds, rl.Vector3{camera.Position.X, camera.Position.Y, camera.Position.Z})

			rl.DrawModel(*sunce_model, nebo.MoveSun(float64(nebo.SkyBodyAngle(currentTick)), camera), 1.0, rl.White)

			rl.EndMode3D()

			rl.DrawFPS(10, 10)
			rl.DrawText("WASD - Kretanje | Mis - Okretanje | Space - Skok", 10, 40, 20, rl.DarkGray)
			rl.DrawText("LMB - Unisti | RMB - Postavi blok", 10, 70, 20, rl.DarkGray)

			// Hotbar crtanje
			menu.CrtajHotbar(navigation.SelectedBlock)

		} else {
			rl.ClearBackground(rl.DarkGray)
			menu.Crtaj()
		}

		rl.EndDrawing()
	}
}
