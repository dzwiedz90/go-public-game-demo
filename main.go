package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/dzwiedz90/go-public-game-demo/constants"
	"github.com/dzwiedz90/go-public-game-demo/engine/core/states"
	"github.com/dzwiedz90/go-public-game-demo/menu"
)

func main() {
	rl.InitWindow(constants.ScreenWidth, constants.ScreenHeight, "Demo Game")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	states.CurrentGameState = menu.MenuState
	states.CurrentGameState.OnEnter()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		states.CurrentGameState.OnUpdate()

		rl.EndDrawing()
	}
}

// func loadFirstMap(e *engine.Engine) {
// 	e.LoadMusic("resources/music/main_theme.mp3")
// 	e.SetPlayerPosition(550, 325)
// 	e.LoadMap("test_house.map")

// 	e.CloseMap = false

// 	for e.Running {
// 		e.Input()
// 		e.Update()
// 		r := e.Render()
// 		if r {
// 			break
// 		}
// 	}

// 	e.UnloadMap()
// }

// func loadSecondMap(e *engine.Engine) {
// 	e.UnloadMusic()
// 	e.LoadMusic("resources/music/db.mp3")
// 	e.SetPlayerPosition(100, 100)
// 	e.LoadMap("inventory.map")

// 	e.CloseMap = false

// 	for e.Running {
// 		e.Input()
// 		e.Update()
// 		r := e.Render()
// 		if r {
// 			break
// 		}
// 	}

// 	e.UnloadMap()
// }
