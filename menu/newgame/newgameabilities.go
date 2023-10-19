package newgame

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/dzwiedz90/go-public-game-demo/constants"
	"github.com/dzwiedz90/go-public-game-demo/dices"
	"github.com/dzwiedz90/go-public-game-demo/engine/core/states"
	"github.com/dzwiedz90/go-public-game-demo/models/skills"
)

type item struct {
	Text       skills.Ability
	Value      int
	PositionX  int32
	PositionY  int32
	IsSelected bool
}

var (
	retake        = 10
	throwsCounter = 0
	itemCounter   = 0

	stageFinished = false

	items = []item{
		{Text: skills.Strength, Value: 0, PositionX: 100, PositionY: 150, IsSelected: true},
		{Text: skills.Dexteriety, Value: 0, PositionX: 100, PositionY: 200, IsSelected: false},
		{Text: skills.Constitution, Value: 0, PositionX: 100, PositionY: 250, IsSelected: false},
		{Text: skills.Ingelligence, Value: 0, PositionX: 100, PositionY: 300, IsSelected: false},
		{Text: skills.Wisdom, Value: 0, PositionX: 100, PositionY: 350, IsSelected: false},
		{Text: skills.Charisma, Value: 0, PositionX: 100, PositionY: 400, IsSelected: false},
	}

	throws []int

	NewGameAbilitiesState = states.GameState{
		StateName: "AbilitiesState",
		OnEnter: func() {
			drawAbilities()
		},
		OnUpdate: func() {
			rl.ClearBackground(rl.NewColor(23, 20, 22, 1))

			rl.UpdateMusicStream(*musicStream)

			input()

			draw()
		},
		OnExit: func() {

		},
	}
)

func input() {
	if rl.IsKeyPressed(rl.KeyDown) {
		if throwsCounter == 5 {
			throwsCounter = 0
		} else {
			throwsCounter += 1
		}
	} else if rl.IsKeyPressed(rl.KeyUp) {
		if throwsCounter == 0 {
			throwsCounter = 5
		} else {
			throwsCounter -= 1
		}
	} else if rl.IsKeyPressed(rl.KeyP) {
		if throws[throwsCounter] != 0 && retake > 0 {
			throws[throwsCounter] = drawAbility()
			retake -= 1
		}
	} else if rl.IsKeyPressed(rl.KeyEnter) {
		if !stageFinished {
			if throws[throwsCounter] != 0 {
				items[itemCounter].Value = throws[throwsCounter]
				items[itemCounter].IsSelected = false
				throws[throwsCounter] = 0
				if itemCounter >= 5 {
					stageFinished = true
				} else {
					itemCounter += 1
					items[itemCounter].IsSelected = true
				}
			}
		} else {
			// nowy stan
		}
	} else if rl.IsKeyPressed(rl.KeyEscape) {
		if stageFinished {
			retake = 10
			throwsCounter = 0
			itemCounter = 0
			stageFinished = false

			throws = throws[:0]

			for i := range items {
				items[i].Value = 0
				items[i].IsSelected = false
			}

			items[0].IsSelected = true

			drawAbilities()
		}
	}
}

func draw() {
	rl.DrawText("Tworzenie postaci", 25, 25, constants.FontSize, rl.DarkGray)
	var selectedItem string

	for _, item := range items {
		if item.IsSelected {
			rl.DrawText(string(item.Text), item.PositionX, item.PositionY, constants.FontSize, rl.White)
			rl.DrawText(fmt.Sprint(item.Value), item.PositionX+300, item.PositionY, constants.FontSize, rl.White)
			selectedItem = string(item.Text)
		} else {
			rl.DrawText(string(item.Text), item.PositionX, item.PositionY, constants.FontSize, rl.DarkGray)
			rl.DrawText(fmt.Sprint(item.Value), item.PositionX+300, item.PositionY, constants.FontSize, rl.DarkGray)
		}
	}

	for i, val := range throws {
		if i == throwsCounter {
			rl.DrawText(fmt.Sprint(val), 500, items[i].PositionY, constants.FontSize, rl.White)
		} else {
			rl.DrawText(fmt.Sprint(val), 500, items[i].PositionY, constants.FontSize, rl.DarkGray)
		}
	}

	rl.DrawText("Ustaw wartosc dla "+selectedItem+" [ENTER]", 500, 25, constants.FontSize, rl.DarkGray)
	rl.DrawText("Przerzuc "+fmt.Sprint(retake)+" wartosci [P]", 500, 75, constants.FontSize, rl.DarkGray)

	if stageFinished {
		rl.DrawText("Cechy rozdane", 100, 500, constants.FontSize, rl.DarkGray)
		rl.DrawText("Kontynuuj [ENTER], losuj jeszcze raz [ESC]", 100, 550, constants.FontSize, rl.DarkGray)
	}
}

func drawAbilities() {
	for i := 0; i <= 5; i++ {
		throws = append(throws, drawAbility())
	}
}

func drawAbility() int {
	tempThrows := make([]int, 4)
	tempThrows[0] = dices.RollD6()
	tempThrows[1] = dices.RollD6()
	tempThrows[2] = dices.RollD6()
	tempThrows[3] = dices.RollD6()

	index := findIndexOfMinValue(tempThrows)
	tempThrows[index] = 0

	return sumIntSlice(tempThrows)
}

func findIndexOfMinValue(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	minIndex := 0
	minValue := arr[0]

	for i, v := range arr {
		if v < minValue {
			minIndex = i
			minValue = v
		}
	}

	return minIndex
}

func sumIntSlice(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
