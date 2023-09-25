package engine

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/dzwiedz90/go-public-game-demo/constants"
)

type Wall struct {
	X, Y int
}

type Engine struct {
	Running  bool
	BkgColor color.RGBA
	CloseMap bool

	GrassSprite  rl.Texture2D
	PlayerSprite rl.Texture2D
	HillSprite   rl.Texture2D
	FenceSprite  rl.Texture2D
	HouseSprite  rl.Texture2D
	WaterSprite  rl.Texture2D
	TilledSprite rl.Texture2D
	Tex          rl.Texture2D

	PlayerSrc                                     rl.Rectangle
	PlayerDest                                    rl.Rectangle
	PlayerMoving                                  bool
	PlayerDir                                     int
	PlayerUp, PlayerDown, PlayerRight, PlayerLeft bool
	PlayerFrame                                   int

	PlayerSpeed float32

	FrameCount int

	TileDest   rl.Rectangle
	TileSrc    rl.Rectangle
	TileMap    []int
	SrcMap     []string
	MapW, MapH int

	MusicPaused bool
	Music       rl.Music

	Cam rl.Camera2D

	Walls []Wall
}

func NewEngine() *Engine {
	return &Engine{
		Running:     true,
		BkgColor:    rl.NewColor(191, 252, 219, 1),
		PlayerSpeed: 3,
	}
}

func (e *Engine) DrawScene() {
	for i := 0; i < len(e.TileMap); i++ {
		if e.TileMap[i] != 0 {
			e.TileDest.X = e.TileDest.Width * float32(i%e.MapW)
			e.TileDest.Y = e.TileDest.Height * float32(i/e.MapW)

			if e.SrcMap[i] == "g" {
				e.Tex = e.GrassSprite
			}
			if e.SrcMap[i] == "l" {
				e.Tex = e.HillSprite
			}
			if e.SrcMap[i] == "f" {
				e.Tex = e.FenceSprite
			}
			if e.SrcMap[i] == "h" {
				e.Tex = e.HouseSprite
			}
			if e.SrcMap[i] == "w" {
				e.Tex = e.WaterSprite
			}
			if e.SrcMap[i] == "t" {
				e.Tex = e.TilledSprite
			}

			if e.SrcMap[i] == "h" || e.SrcMap[i] == "f" {
				e.TileSrc.X = 0
				e.TileSrc.Y = 0
				rl.DrawTexturePro(e.GrassSprite, e.TileSrc, e.TileDest, rl.NewVector2(e.TileDest.Width, e.TileDest.Height), 0, rl.White)
			}

			e.TileSrc.X = e.TileSrc.Width * float32((e.TileMap[i]-1)%int(e.Tex.Width/int32(e.TileSrc.Width)))
			e.TileSrc.Y = e.TileSrc.Height * float32((e.TileMap[i]-1)/int(e.Tex.Width/int32(e.TileSrc.Width)))
			rl.DrawTexturePro(e.Tex, e.TileSrc, e.TileDest, rl.NewVector2(e.TileDest.Width, e.TileDest.Height), 0, rl.White)
		}
	}

	rl.DrawTexturePro(e.PlayerSprite, e.PlayerSrc, e.PlayerDest, rl.NewVector2(e.PlayerDest.Width, e.PlayerDest.Height), 0, rl.White)
}

func (e *Engine) Input() {
	if rl.IsKeyDown(rl.KeyUp) {
		e.PlayerMoving = true
		e.PlayerDir = 1
		e.PlayerUp = true
	}
	if rl.IsKeyDown(rl.KeyDown) {
		e.PlayerMoving = true
		e.PlayerDir = 0
		e.PlayerDown = true
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		e.PlayerMoving = true
		e.PlayerDir = 2
		e.PlayerLeft = true
	}
	if rl.IsKeyDown(rl.KeyRight) {
		e.PlayerMoving = true
		e.PlayerDir = 3
		e.PlayerRight = true
	}
	if rl.IsKeyPressed(rl.KeyP) {
		e.MusicPaused = !e.MusicPaused
	}
	if rl.IsKeyPressed(rl.KeyO) {
		e.CloseMap = true
	}
	// if rl.IsKeyPressed(rl.KeyI) {
	// 	a, _, _ := inventory.Inventory(e.LoadMap, e.TileMap, e.SrcMap)
	// 	e.PlayerDest = a
	// 	fmt.Println(e.PlayerDest)
	// 	fmt.Println(e.SrcMap)
	// }
}
func (e *Engine) Update() {
	e.Running = !rl.WindowShouldClose()

	e.PlayerSrc.X = e.PlayerSrc.Width * float32(e.PlayerFrame)

	if e.PlayerMoving {
		if e.PlayerUp {
			newY := e.PlayerDest.Y - e.PlayerSpeed
			if !e.isCollision(e.PlayerDest.X, newY) {
				e.PlayerDest.Y = newY
			}
		}
		if e.PlayerDown {
			newY := e.PlayerDest.Y + e.PlayerSpeed
			if !e.isCollision(e.PlayerDest.X, newY) {
				e.PlayerDest.Y = newY
			}
		}
		if e.PlayerLeft {
			newX := e.PlayerDest.X - e.PlayerSpeed
			if !e.isCollision(newX, e.PlayerDest.Y) {
				e.PlayerDest.X = newX
			}
		}
		if e.PlayerRight {
			newX := e.PlayerDest.X + e.PlayerSpeed
			if !e.isCollision(newX, e.PlayerDest.Y) {
				e.PlayerDest.X = newX
			}
		}
		if e.FrameCount%8 == 1 {
			e.PlayerFrame++
		}
	} else if e.FrameCount%45 == 1 {
		e.PlayerFrame++
	}

	e.FrameCount++

	if e.PlayerFrame > 3 {
		e.PlayerFrame = 0
	}
	if !e.PlayerMoving && e.PlayerFrame > 1 {
		e.PlayerFrame = 0
	}

	e.PlayerSrc.X = e.PlayerSrc.Width * float32(e.PlayerFrame)
	e.PlayerSrc.Y = e.PlayerSrc.Height * float32(e.PlayerDir)

	rl.UpdateMusicStream(e.Music)
	if e.MusicPaused {
		rl.PauseMusicStream(e.Music)
	} else {
		rl.ResumeMusicStream(e.Music)
	}

	e.Cam.Target = rl.NewVector2(float32(e.PlayerDest.X-(e.PlayerDest.Width/2)), float32(e.PlayerDest.Y-(e.PlayerDest.Height/2)))

	e.PlayerMoving = false
	e.PlayerUp, e.PlayerDown, e.PlayerRight, e.PlayerLeft = false, false, false, false
}

func (e *Engine) isCollision(playerX, playerY float32) bool {
	pX := int(playerX) / 16
	pY := int(playerY) / 16

	fmt.Println("pX: ", pX)
	fmt.Println("pY: ", pY)
	fmt.Println(" ")

	for _, wall := range e.Walls {
		if pX == wall.X && pY == wall.Y {
			fmt.Println("wallX: ", wall.X)
			fmt.Println("wallY: ", wall.Y)
			fmt.Println(" ")
			return true
		}
	}
	return false
}

func (e *Engine) Render() bool {
	rl.BeginDrawing()
	rl.ClearBackground(e.BkgColor)
	rl.BeginMode2D(e.Cam)
	e.DrawScene()
	rl.EndMode2D()
	rl.EndDrawing()

	return e.CloseMap
}

func (e *Engine) LoadMap(mapFile string) {
	file, err := ioutil.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	remNewLines := strings.Replace(string(file), "\n", " ", -1)
	sliced := strings.Split(remNewLines, " ")

	e.MapW = -1
	e.MapH = -1

	for i := 0; i < len(sliced); i++ {
		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		m := int(s)
		// Initially MapW and MapH are -1, two first digits in the map are describing maps width and height and first if and else if loads those
		if e.MapW == -1 {
			e.MapW = m
		} else if e.MapH == -1 {
			e.MapH = m
		} else if i < e.MapW*e.MapH+2 {
			// loads tile number from image from which it loads starting from top left corner of image
			e.TileMap = append(e.TileMap, m)
		} else {
			// load type of sprite to be used
			l := sliced[i]
			e.SrcMap = append(e.SrcMap, l)
			if l == "w" || l == "f" || l == "h" {
				x := (i - e.MapW*e.MapH - 2) % e.MapW
				y := (i - e.MapW*e.MapH - 2) / e.MapW
				e.Walls = append(e.Walls, Wall{X: x, Y: y})
			}
		}
	}

	if len(e.TileMap) > e.MapW*e.MapH {
		e.TileMap = e.TileMap[:len(e.TileMap)-1]
	}
}

func (e *Engine) Init() {
	e.GrassSprite = rl.LoadTexture("resources/Tilesets/groundtiles/oldtiles/Grass.png")
	e.HillSprite = rl.LoadTexture("resources/Tilesets/groundtiles/oldtiles/Hills.png")
	e.FenceSprite = rl.LoadTexture("resources/Tilesets/Fences.png")
	e.HouseSprite = rl.LoadTexture("resources/Tilesets/WoodenHouse.png")
	e.WaterSprite = rl.LoadTexture("resources/Tilesets/Water.png")
	e.TilledSprite = rl.LoadTexture("resources/Tilesets/groundtiles/oldtiles/Tilled Dirt.png")

	e.TileDest = rl.NewRectangle(0, 0, 16, 16)
	e.TileSrc = rl.NewRectangle(0, 0, 16, 16)

	e.PlayerSprite = rl.LoadTexture("resources/Characters/BasicCharakterSpritesheet.png")
}

func (e *Engine) SetPlayerPosition(x, y float32) {
	e.PlayerSrc = rl.NewRectangle(0, 0, 48, 48)
	// last 2 describes player size, first 2 player's starting position on map
	e.PlayerDest = rl.NewRectangle(x, y, 48, 48)

	e.Cam = rl.NewCamera2D(rl.NewVector2(float32(constants.ScreenWidth/2), float32(constants.ScreenHeight/2)), rl.NewVector2(float32(e.PlayerDest.X-(e.PlayerDest.Width/2)), float32(e.PlayerDest.Y-(e.PlayerDest.Height/2))), 0.0, 1.5)
	e.Cam.Zoom = 2.5
}

func (e *Engine) LoadMusic(stream string) {
	rl.InitAudioDevice()
	e.Music = rl.LoadMusicStream(stream)
	e.MusicPaused = false
	rl.PlayMusicStream(e.Music)
}

func (e *Engine) UnloadMusic() {
	rl.UnloadMusicStream(e.Music)
	rl.CloseAudioDevice()
}

func (e *Engine) UnloadMap() {
	e.TileMap = []int{}
	e.SrcMap = []string{}
	e.Walls = []Wall{}
	e.Tex = rl.Texture2D{}
}

func (e *Engine) Quit() {
	rl.UnloadTexture(e.GrassSprite)
	rl.UnloadTexture(e.PlayerSprite)
	rl.UnloadMusicStream(e.Music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
