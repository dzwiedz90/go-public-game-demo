package newgame

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/dzwiedz90/go-public-game-demo/constants"
	"github.com/dzwiedz90/go-public-game-demo/engine/core/states"
)

var (
	NewGameSpellsState = states.GameState{
		StateName: "SpellsState",
		OnEnter: func() {
		},
		OnUpdate: func() {
			rl.ClearBackground(rl.NewColor(23, 20, 22, 1))

			rl.UpdateMusicStream(*musicStream)

			inputSpells()

			drawSpellsOutput()
		},
		OnExit: func() {

		},
	}
)

func inputSpells() {
	rl.DrawText("Press enter to choose", 500, 100, constants.FontSize, rl.DarkGray)

	if rl.IsKeyPressed(rl.KeyDown) {

	} else if rl.IsKeyPressed(rl.KeyUp) {

	} else if rl.IsKeyPressed(rl.KeyEnter) {

	} else if rl.IsKeyPressed(rl.KeyEscape) {

	}
}

func drawSpellsOutput() {
	rl.DrawText("Tworzenie postaci", 25, 25, constants.FontSize, rl.DarkGray)

	rl.DrawText("Pochodzenie", 50, 100, constants.FontSize+5, rl.White)
}

func saveSpells() {

}
