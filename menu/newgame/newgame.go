package newgame

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/dzwiedz90/go-public-game-demo/engine/core"
	"github.com/dzwiedz90/go-public-game-demo/engine/core/states"
)

type letter struct {
	letter     string
	isSelected bool
}

var (
	// TODO
	nameStage      = true
	raceStage      = false
	classStage     = false
	characterStage = false

	musicStream *rl.Music

	selectedRow    = 0
	selectedColumn = 0

	name string

	alphabet = [][]letter{
		{
			{letter: "a", isSelected: true},
			{letter: "b", isSelected: false},
			{letter: "c", isSelected: false},
			{letter: "d", isSelected: false},
			{letter: "e", isSelected: false},
		},
		{
			{letter: "f", isSelected: false},
			{letter: "g", isSelected: false},
			{letter: "h", isSelected: false},
			{letter: "i", isSelected: false},
			{letter: "j", isSelected: false},
		},
		{
			{letter: "k", isSelected: false},
			{letter: "l", isSelected: false},
			{letter: "m", isSelected: false},
			{letter: "n", isSelected: false},
			{letter: "o", isSelected: false},
		},
		{
			{letter: "p", isSelected: false},
			{letter: "q", isSelected: false},
			{letter: "r", isSelected: false},
			{letter: "s", isSelected: false},
			{letter: "t", isSelected: false},
		},
		{
			{letter: "u", isSelected: false},
			{letter: "v", isSelected: false},
			{letter: "w", isSelected: false},
			{letter: "x", isSelected: false},
			{letter: "y", isSelected: false},
		},
		{
			{letter: "z", isSelected: false},
			{letter: "DEL", isSelected: false},
			{letter: "SPACE", isSelected: false},
			{letter: "OK", isSelected: false},
		},
	}

	NewGameState = states.GameState{
		StateName: "MenuState",
		OnEnter: func() {
			core.Character.Level = 1
			core.Character.ExperiencePoints = 0
		},
		OnUpdate: func() {
			rl.ClearBackground(rl.NewColor(23, 20, 22, 1))

			// TODO dodac ze jak jest na maksa z prawej lub lewej to idzie do nastepnej lini zerowa pozycja

			if rl.IsKeyPressed(rl.KeyDown) {
				if selectedRow == 4 {
					if selectedColumn == 3 {
						alphabet[selectedRow][selectedColumn].isSelected = false
						selectedColumn = 2
					}
					if selectedColumn == 4 {
						alphabet[selectedRow][selectedColumn].isSelected = false
						selectedColumn = 3
					}
				}
				if selectedRow < 5 {
					alphabet[selectedRow][selectedColumn].isSelected = false
					selectedRow += 1
					alphabet[selectedRow][selectedColumn].isSelected = true
				}
			} else if rl.IsKeyPressed(rl.KeyUp) {
				if selectedRow > 0 {
					if selectedRow == 5 && selectedColumn == 3 {
						alphabet[selectedRow][selectedColumn].isSelected = false
						selectedRow = 4
						selectedColumn = 4
						alphabet[selectedRow][selectedColumn].isSelected = true
					} else {
						alphabet[selectedRow][selectedColumn].isSelected = false
						selectedRow -= 1
						alphabet[selectedRow][selectedColumn].isSelected = true
					}
				}
			} else if rl.IsKeyPressed(rl.KeyRight) {
				if selectedRow == 5 && selectedColumn < 3 {
					alphabet[selectedRow][selectedColumn].isSelected = false
					selectedColumn += 1
					alphabet[selectedRow][selectedColumn].isSelected = true
				}
				if selectedRow < 5 && selectedColumn < 4 {
					alphabet[selectedRow][selectedColumn].isSelected = false
					selectedColumn += 1
					alphabet[selectedRow][selectedColumn].isSelected = true
				}
			} else if rl.IsKeyPressed(rl.KeyLeft) {
				if selectedColumn > 0 {
					alphabet[selectedRow][selectedColumn].isSelected = false
					selectedColumn -= 1
					alphabet[selectedRow][selectedColumn].isSelected = true
				}
			} else if rl.IsKeyPressed(rl.KeyEnter) {
				l := alphabet[selectedRow][selectedColumn].letter
				if l == "DEL" {
					name = name[:len(name)-1]
				} else if l == "SPACE" {
					name += " "
				} else if l == "OK" {
					nameStage = false
					core.Character.Name = name
				} else {
					if len(name) == 0 {
						name += strings.ToUpper(l)
					} else if string(name[len(name)-1]) == " " {
						name += strings.ToUpper(l)
					} else {
						name += l
					}
				}
			}

			rl.UpdateMusicStream(*musicStream)

			drawMenu()
		},
		OnExit: func() {

		},
	}
)

func drawMenu() {
	rl.DrawText("Tworzenie postaci", 25, 25, 32, rl.DarkGray)

	if nameStage {
		rl.DrawText("Imie: "+name, 50, 100, 32, rl.White)
		drawAlphabet()
	} else {
		rl.DrawText("Imie: "+name, 50, 100, 32, rl.DarkGray)
	}
}

func drawAlphabet() {
	var x int32
	var y int32
	x = 700
	y = 100
	for _, row := range alphabet {
		for _, letter := range row {
			if letter.letter == "OK" {
				x += 75
			}
			if letter.isSelected {
				rl.DrawText(letter.letter, x, y, 32, rl.White)
			} else {
				rl.DrawText(letter.letter, x, y, 32, rl.DarkGray)
			}
			x += 75
		}
		x = 700
		y += 75
	}
}

func SetNewGameMusicStream(musicStreamIn *rl.Music) {
	musicStream = musicStreamIn
}
