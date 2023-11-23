package newgame

import (
	"fmt"

	"github.com/dzwiedz90/go-public-game-demo/constants"
	"github.com/dzwiedz90/go-public-game-demo/engine/core/states"
	"github.com/dzwiedz90/go-public-game-demo/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	NewGameAncestryState = states.GameState{
		StateName: "AbilitiesState",
		OnEnter: func() {
			mkAvailableSkillsSelections()
			skillsItems = mkSkillItems()
		},
		OnUpdate: func() {
			rl.ClearBackground(rl.NewColor(23, 20, 22, 1))

			rl.UpdateMusicStream(*musicStream)

			inputAncestry()

			drawAncestryOutput()
		},
		OnExit: func() {

		},
	}
)

func inputAncestry() {
	if rl.IsKeyPressed(rl.KeyDown) && !stageDone {
		if skillsCounter == len(skillsItems)-1 {
			skillsCounter = 0
		} else {
			skillsCounter += 1
		}
	} else if rl.IsKeyPressed(rl.KeyUp) && !stageDone {
		if skillsCounter == 0 {
			skillsCounter = len(skillsItems) - 1
		} else {
			skillsCounter -= 1
		}
	} else if rl.IsKeyPressed(rl.KeyEnter) {
		if stageDone {
			saveSkills()
			// states.CurrentGameState = NewGameSkillsState
			// NewGameSkillsState.OnEnter()
		} else {
			if skillsItems[skillsCounter].isSelected {
				skillsItems[skillsCounter].isSelected = false
				availableSkillsSelections += 1
				skillsSelected = utils.RemoveElement(skillsSelected, skillsItems[skillsCounter].text)
			} else if !skillsItems[skillsCounter].isSelected && availableSkillsSelections > 0 {
				skillsItems[skillsCounter].isSelected = true
				skillsSelected = append(skillsSelected, skillsItems[skillsCounter].text)
				availableSkillsSelections -= 1
				skillsSelected = append(skillsSelected, skillsItems[skillsCounter].text)

				if availableSkillsSelections == 0 {
					stageDone = true
				}
			}
		}
	} else if rl.IsKeyPressed(rl.KeyEscape) {
		skillsItems = mkSkillItems()
		skillsSelected = []string{}
		skillsCounter = 0
		mkAvailableSkillsSelections()
		stageDone = false
	}
}

func drawAncestryOutput() {
	rl.DrawText("Tworzenie postaci", 25, 25, constants.FontSize, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("Umiejetnosci, wybierz %v sposrod nastepujacych,", availableSkillsSelections), 500, 25, constants.FontSize, rl.DarkGray)
	rl.DrawText("w ktorych bedziesz miec bieglosc", 500, 75, constants.FontSize, rl.DarkGray)

	var y int32 = 200

	for i, skill := range skillsItems {
		if i <= 9 {
			if i == skillsCounter {
				rl.DrawText(skill.text, 25, y, constants.FontSize, rl.White)

				if skillsItems[i].isSelected {
					rl.DrawText("[X]", 500, y, constants.FontSize, rl.White)
				} else {
					rl.DrawText("[ ]", 500, y, constants.FontSize, rl.White)
				}
			} else {
				rl.DrawText(skill.text, 25, y, constants.FontSize, rl.DarkGray)

				if skillsItems[i].isSelected {
					rl.DrawText("[X]", 500, y, constants.FontSize, rl.DarkGray)
				} else {
					rl.DrawText("[ ]", 500, y, constants.FontSize, rl.DarkGray)
				}
			}

			y += 50
		} else {
			if i == 10 {
				y = 200
			}

			if i == skillsCounter {
				rl.DrawText(skill.text, 600, y, constants.FontSize, rl.White)

				if skillsItems[i].isSelected {
					rl.DrawText("[X]", 1075, y, constants.FontSize, rl.White)
				} else {
					rl.DrawText("[ ]", 1075, y, constants.FontSize, rl.White)
				}
			} else {
				rl.DrawText(skill.text, 600, y, constants.FontSize, rl.DarkGray)

				if skillsItems[i].isSelected {
					rl.DrawText("[X]", 1075, y, constants.FontSize, rl.DarkGray)
				} else {
					rl.DrawText("[ ]", 1075, y, constants.FontSize, rl.DarkGray)
				}
			}

			y += 50
		}
	}

	if stageDone {
		rl.DrawText("Umiejetnosci wybrane, nacisnij enter, aby kontynuowac lub ESC aby wybrac ponownie...", 25, 125, constants.FontSize, rl.White)
	}
}
