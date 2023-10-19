package menu

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/dzwiedz90/go-public-game-demo/constants"
	"github.com/dzwiedz90/go-public-game-demo/engine/core"
	"github.com/dzwiedz90/go-public-game-demo/engine/core/states"
	"github.com/dzwiedz90/go-public-game-demo/menu/newgame"
	"github.com/dzwiedz90/go-public-game-demo/models/menumodels"
)

var (
	menuItems = []menumodels.MenuItem{
		{Text: "Nowa gra", Value: 0, PositionX: 100, PositionY: 300, IsSelected: true},
		{Text: "Wyjscie", Value: 0, PositionX: 100, PositionY: 350, IsSelected: false},
	}

	selectedItem = 0
	musicStream  rl.Music

	MenuState = states.GameState{
		StateName: "MenuState",
		OnEnter: func() {
			musicStream = core.LoadMusic("resources/music/db.mp3")
		},
		OnUpdate: func() {
			rl.ClearBackground(rl.NewColor(23, 20, 22, 1))

			if rl.IsKeyPressed(rl.KeyDown) {
				menuItems[selectedItem].IsSelected = false
				selectedItem = (selectedItem + 1) % len(menuItems)
				menuItems[selectedItem].IsSelected = true
			} else if rl.IsKeyPressed(rl.KeyUp) {
				menuItems[selectedItem].IsSelected = false
				selectedItem = (selectedItem - 1 + len(menuItems)) % len(menuItems)
				menuItems[selectedItem].IsSelected = true
			} else if rl.IsKeyPressed(rl.KeyEnter) {
				selectedOption := menuItems[selectedItem].Text
				switch selectedOption {
				case "Nowa gra":
					newgame.SetNewGameMusicStream(&musicStream)
					states.CurrentGameState = newgame.NewGameState
					newgame.NewGameState.OnEnter()
				case "Wyjscie":
					rl.UnloadMusicStream(musicStream)
					rl.CloseAudioDevice()
					rl.CloseWindow()
				}
			}

			rl.UpdateMusicStream(musicStream)

			drawMenu(menuItems)

		},
		OnExit: func() {

		},
	}
)

func drawMenu(menuItems []menumodels.MenuItem) {
	for _, item := range menuItems {
		if item.IsSelected {
			rl.DrawText(item.Text, item.PositionX, item.PositionY, constants.FontSize, rl.White)
		} else {
			rl.DrawText(item.Text, item.PositionX, item.PositionY, constants.FontSize, rl.DarkGray)
		}
	}
}
