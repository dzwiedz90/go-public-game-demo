package newgame

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/dzwiedz90/go-public-game-demo/constants"
	"github.com/dzwiedz90/go-public-game-demo/dices"
	"github.com/dzwiedz90/go-public-game-demo/engine/core"
	"github.com/dzwiedz90/go-public-game-demo/engine/core/states"
	"github.com/dzwiedz90/go-public-game-demo/models/class"
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

	abilitiesItems = []item{
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

			inputAbilities()

			drawAbilitiesOutput()
		},
		OnExit: func() {

		},
	}
)

func inputAbilities() {
	if rl.IsKeyPressed(rl.KeyDown) && !stageFinished {
		if throwsCounter == 5 {
			throwsCounter = 0
		} else {
			throwsCounter += 1
		}
	} else if rl.IsKeyPressed(rl.KeyUp) && !stageFinished {
		if throwsCounter == 0 {
			throwsCounter = 5
		} else {
			throwsCounter -= 1
		}
	} else if rl.IsKeyPressed(rl.KeyP) && !stageFinished {
		if throws[throwsCounter] != 0 && retake > 0 {
			throws[throwsCounter] = drawAbility()
			retake -= 1
		}
	} else if rl.IsKeyPressed(rl.KeyEnter) {
		if !stageFinished {
			if throws[throwsCounter] != 0 {
				abilitiesItems[itemCounter].Value = throws[throwsCounter]
				abilitiesItems[itemCounter].IsSelected = false
				throws[throwsCounter] = 0
				if itemCounter >= 5 {
					stageFinished = true
				} else {
					itemCounter += 1
					abilitiesItems[itemCounter].IsSelected = true
				}
			}
		} else {
			core.Character.Abilities.Strength = abilitiesItems[0].Value
			core.Character.Abilities.Dexteriety = abilitiesItems[1].Value
			core.Character.Abilities.Constitution = abilitiesItems[2].Value
			core.Character.Abilities.Ingelligence = abilitiesItems[3].Value
			core.Character.Abilities.Wisdom = abilitiesItems[4].Value
			core.Character.Abilities.Charisma = abilitiesItems[5].Value

			core.Character.PassiveWisdomPerception = 10 + constants.ModifierFromAbility[core.Character.Abilities.Wisdom]

			core.Character.MaxHitPoints = class.ClassHitPointsDie[core.Character.Class] + constants.ModifierFromAbility[core.Character.Abilities.Constitution]
			core.Character.CurrentHitPoint = core.Character.MaxHitPoints
			core.Character.TemporaryHitPoints = 0

			states.CurrentGameState = NewGameSkillsState
			NewGameSkillsState.OnEnter()
		}
	} else if rl.IsKeyPressed(rl.KeyEscape) {
		if stageFinished {
			retake = 10
			throwsCounter = 0
			itemCounter = 0
			stageFinished = false

			throws = throws[:0]

			for i := range abilitiesItems {
				abilitiesItems[i].Value = 0
				abilitiesItems[i].IsSelected = false
			}

			abilitiesItems[0].IsSelected = true

			drawAbilities()
		}
	}
}

func drawAbilitiesOutput() {
	rl.DrawText("Tworzenie postaci", 25, 25, constants.FontSize, rl.DarkGray)
	var selectedItem string

	for _, item := range abilitiesItems {
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
			rl.DrawText(fmt.Sprint(val), 500, abilitiesItems[i].PositionY, constants.FontSize, rl.White)
		} else {
			rl.DrawText(fmt.Sprint(val), 500, abilitiesItems[i].PositionY, constants.FontSize, rl.DarkGray)
		}
	}

	rl.DrawText("Cechy", 500, 25, constants.FontSize, rl.DarkGray)
	rl.DrawText("Ustaw wartosc dla "+selectedItem+" [ENTER]", 700, 25, constants.FontSize, rl.DarkGray)
	rl.DrawText("Przerzuc "+fmt.Sprint(retake)+" wartosci [P]", 700, 75, constants.FontSize, rl.DarkGray)

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
