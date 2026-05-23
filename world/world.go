package world

import (
	"main/blocks"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Chunk struct {
	GlobalX          int
	GlobalZ          int
	Blocks           [][][]blocks.Block
	Rendered         bool
	CachedTransforms map[blocks.Block][]rl.Matrix
	IsDirty          bool
}

type ChunkPos struct {
	X, Z int
}

var LoadedChunks = make(map[ChunkPos]*Chunk)
var LoadedStructures = make(map[ChunkPos]*Chunk)

var BlockModelRegistryInstance *BlockRegistry

type BlockRegistry struct {
	BlockModels      map[blocks.Block]rl.Model
	InstancingShader rl.Shader
}

func (br *BlockRegistry) RegisterNewModel(block blocks.Block, cubeMesh rl.Mesh, path string) {
	new_model := rl.LoadModelFromMesh(cubeMesh)

	// CRUCIAL: Load texture using raylib defaults
	// Bind our verified instancing shader program safely to the mesh material
	new_model.Materials.Shader = br.InstancingShader

	// Initialize default material diffuse color tint so it isn't transparent black (0,0,0,0)
	new_model.Materials.GetMap(rl.MapDiffuse).Color = rl.White
	texture := rl.LoadTexture(path)
	rl.SetMaterialTexture(new_model.Materials, rl.MapDiffuse, texture)

	br.BlockModels[block] = new_model
}

func BlockModelRegistry() *BlockRegistry {
	if BlockModelRegistryInstance == nil {
		shader := rl.LoadShader("instancing.vs", "instancing.fs")

		// 2. Map standard raylib camera uniform location maps
		shader.UpdateLocation(rl.ShaderLocMatrixMvp, rl.GetShaderLocation(shader, "mvp"))
		shader.UpdateLocation(rl.ShaderLocMapDiffuse, rl.GetShaderLocation(shader, "texture0"))
		shader.UpdateLocation(rl.ShaderLocColorDiffuse, rl.GetShaderLocation(shader, "colDiffuse"))

		// 3. Map the instancing location attribute channel
		// This tells raylib's internal C loop exactly where to dump your matrix data array
		shader.UpdateLocation(rl.ShaderLocMatrixModel, rl.GetShaderLocationAttrib(shader, "instanceTransform"))

		cubeMesh := rl.GenMeshCube(1.0, 1.0, 1.0)
		BlockModelRegistryInstance = &BlockRegistry{
			BlockModels:      make(map[blocks.Block]rl.Model),
			InstancingShader: shader,
		}

		BlockModelRegistryInstance.RegisterNewModel(blocks.Grass, cubeMesh, "assets/grass.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Dirt, cubeMesh, "assets/dirt.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Water, cubeMesh, "assets/water.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Stone, cubeMesh, "assets/stone.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Bedrock, cubeMesh, "assets/bedrock.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Log, cubeMesh, "assets/log.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Leaves, cubeMesh, "assets/leaves.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Netherrack, cubeMesh, "assets/netherrack.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Sand, cubeMesh, "assets/sand.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Snow, cubeMesh, "assets/snow.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Sudomil_bot, cubeMesh, "assets/sudomil_bottom.png")
		BlockModelRegistryInstance.RegisterNewModel(blocks.Sudomil_top, cubeMesh, "assets/sudomil_top.png")
	}
	return BlockModelRegistryInstance
}

func GetGlobalBlock(worldX, worldY, worldZ int) blocks.Block {
	if worldY < 0 || worldY >= 64 {
		return blocks.Air
	}
	cx := int(math.Floor(float64(worldX) / 16.0))
	cz := int(math.Floor(float64(worldZ) / 16.0))

	if chunk, exists := LoadedChunks[ChunkPos{cx, cz}]; exists {
		lx := worldX - cx*16
		lz := worldZ - cz*16
		if lx >= 0 && lx < 16 && lz >= 0 && lz < 16 {
			return chunk.Blocks[lx][worldY][lz]
		}
	}
	return blocks.Air
}

func SetGlobalBlock(worldX, worldY, worldZ int, b blocks.Block) bool {
	if worldY < 0 || worldY >= 64 {
		return false
	}
	cx := int(math.Floor(float64(worldX) / 16.0))
	cz := int(math.Floor(float64(worldZ) / 16.0))

	if chunk, exists := LoadedChunks[ChunkPos{cx, cz}]; exists {
		lx := worldX - cx*16
		lz := worldZ - cz*16
		if lx >= 0 && lx < 16 && lz >= 0 && lz < 16 {
			chunk.Blocks[lx][worldY][lz] = b
			chunk.IsDirty = true
			return true
		}
	}
	return false
}

func (c *Chunk) RebuildMeshCache() {
	c.CachedTransforms = make(map[blocks.Block][]rl.Matrix)

	xLen := len(c.Blocks)
	if xLen == 0 {
		return
	}
	yLen := len(c.Blocks[0])

	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			for z := 0; z < 16; z++ {
				blockType := c.Blocks[x][y][z]
				if blockType == blocks.Air {
					continue
				}

				// FAST LOCAL ARRAY READS: Avoid running global map lookups in walls
				var left, right, down, up, back, front blocks.Block

				if x > 0 {
					left = c.Blocks[x-1][y][z]
				} else {
					left = GetGlobalBlock(x+c.GlobalX*16-1, y, z+c.GlobalZ*16)
				}
				if x < xLen-1 {
					right = c.Blocks[x+1][y][z]
				} else {
					right = GetGlobalBlock(x+c.GlobalX*16+1, y, z+c.GlobalZ*16)
				}
				if y > 0 {
					down = c.Blocks[x][y-1][z]
				} else {
					down = blocks.Air
				}
				if y < yLen-1 {
					up = c.Blocks[x][y+1][z]
				} else {
					up = blocks.Air
				}
				if z > 0 {
					back = c.Blocks[x][y][z-1]
				} else {
					back = GetGlobalBlock(x+c.GlobalX*16, y, z+c.GlobalZ*16-1)
				}
				if z < 15 {
					front = c.Blocks[x][y][z+1]
				} else {
					front = GetGlobalBlock(x+c.GlobalX*16, y, z+c.GlobalZ*16+1)
				}

				// Face Culling optimization
				if left == blocks.Air || right == blocks.Air || down == blocks.Air || up == blocks.Air || back == blocks.Air || front == blocks.Air {
					wx := x + c.GlobalX*16
					wz := z + c.GlobalZ*16

					translation := rl.MatrixTranslate(float32(wx), float32(y), float32(wz))
					c.CachedTransforms[blockType] = append(c.CachedTransforms[blockType], translation)
				}
			}
		}
	}
	c.IsDirty = false
}

func RenderChunk(c *Chunk) {
	if c.IsDirty || c.CachedTransforms == nil {
		c.RebuildMeshCache()
	}

	RegistryInstance := BlockModelRegistry()

	for blockType, transforms := range c.CachedTransforms {
		count := len(transforms)
		if count == 0 {
			continue
		}

		val, ok := RegistryInstance.BlockModels[blockType]
		if !ok {
			continue
		}

		mesh := *val.Meshes
		material := *val.Materials

		// Feed the low-level rendering loop a clean address to index zero
		rl.DrawMeshInstanced(mesh, material, transforms, count)
	}
}

func WorldToLocal(chunk Chunk, wx, wy, wz int) (lx, ly, lz int) {
	lx = wx - chunk.GlobalX*16
	ly = wy
	lz = wz - chunk.GlobalZ*16
	return lx, ly, lz
}

func LocalToWorld(chunk Chunk, lx, ly, lz int) (wx, wy, wz int) { //mozda nije dobro, prekopiro sam dimijevo otp
	wx = lx + chunk.GlobalX*16
	wy = ly
	wz = lz - chunk.GlobalZ*16
	return lx, ly, lz
}

// func Highlight() {
// 	return
// }
